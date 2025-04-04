package constraint

import (
	"fmt"
	"strings"

	"github.com/transientvariable/cadre/validation"
)

// NotBlank defines the properties required for asserting that a character sequence is not blank.
type NotBlank struct {
	Name    string
	Field   string
	Message string
}

// Validate performs the validation for the character sequence.
func (v NotBlank) Validate(result *validation.Result) {
	if strings.TrimSpace(v.Field) != "" {
		return
	}

	if len(v.Message) > 0 {
		result.Add(v.Name, v.Message)
		return
	}
	result.Add(v.Name, fmt.Sprintf("field cannot be blank: %s", v.Name))
}
