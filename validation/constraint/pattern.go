package constraint

import (
	"fmt"
	"regexp"

	"github.com/transientvariable/schema-go/validation"
)

// Pattern defines the properties required for performing regular expression matches.
type Pattern struct {
	Name    string
	Field   string
	Expr    string
	Message string
}

// Validate performs the validation based on the regular expression match.
func (v Pattern) Validate(result *validation.Result) {
	r := regexp.MustCompile(v.Expr)
	if r.Match([]byte(v.Field)) {
		return
	}

	if len(v.Message) > 0 {
		result.Add(v.Name, v.Message)
		return
	}
	result.Add(v.Name, fmt.Sprintf("field does not match the expected format: %s", v.Name))
}
