package validation

import (
	"strings"
	"sync"

	"github.com/transientvariable/support-go"
)

// Result is a container for tracking constraint violations during the validation process.
type Result struct {
	mutex      sync.RWMutex
	violations map[string][]string
}

// Validator must be implemented in order to pass the validator object into the Validate function.
type Validator interface {
	Validate(result *Result)
}

// NewResult creates a new validation Result.
func NewResult() *Result {
	return &Result{violations: make(map[string][]string)}
}

// Error implements the error interface.
func (r *Result) Error() string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var errs []string
	for _, v := range r.violations {
		errs = append(errs, v...)
	}
	return strings.Join(errs, "\n")
}

// Size returns the number of violations.
func (r *Result) Size() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.violations)
}

// IsValid returns true if the Result contains no constraint violations, false otherwise.
func (r *Result) IsValid() bool {
	if r == nil {
		return true
	}
	return r.Size() < 1
}

// Append adds the provided Result to this Result.
func (r *Result) Append(other *Result) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for key, value := range other.violations {
		for _, msg := range value {
			r.Add(key, msg)
		}
	}
}

// Add will add a new message to the list of violations using the given key. If the key already exists the message will
// be appended to the array of the existing messages.
func (r *Result) Add(key string, message any) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	switch message.(type) {
	case error:
		r.violations[key] = append(r.violations[key], message.(error).Error())
		return
	case string:
		r.violations[key] = append(r.violations[key], message.(string))
		return
	default:
		r.violations[key] = append(r.violations[key], "unknown message type")
	}
}

// Violations returns the names of fields which have constraint violations.
func (r *Result) Violations(key string) []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.violations[key]
}

// Keys returns the name of all fields which have constraint violations.
func (r *Result) Keys() []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var keys []string
	for k := range r.violations {
		keys = append(keys, k)
	}
	return keys
}

func (r *Result) String() string {
	return string(support.ToJSONFormatted(r.violations))
}

// Validate applies the provided validation constraints and returns a Result representing the results.
func Validate(validators ...Validator) *Result {
	result := NewResult()
	wg := &sync.WaitGroup{}
	for i := range validators {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			validator := validators[i]
			validator.Validate(result)
		}(wg, i)
	}
	wg.Wait()
	return result
}

type vfWrapper struct {
	vf func(errors *Result)
}

func (v vfWrapper) Validate(errors *Result) {
	v.vf(errors)
}

// ValidatorFunc wraps any function in a "Validator" to make it easy to write custom ones.
func ValidatorFunc(fn func(errors *Result)) Validator {
	return vfWrapper{fn}
}
