package schema

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/transientvariable/support"
	"github.com/transientvariable/validation"
	"github.com/transientvariable/validation/constraint"

	"github.com/google/uuid"

	json "github.com/json-iterator/go"
)

// User defines the properties for a user.
type User struct {
	Active        bool           `json:"active" swaggerignore:"true"`
	Created       *time.Time     `json:"created"`
	DisplayName   sql.NullString `json:"display_name,omitempty"`
	Email         string         `json:"email"`
	EmailVerified bool           `json:"email_verified" swaggerignore:"true"`
	FullName      sql.NullString `json:"full_name,omitempty"`
	ID            uuid.UUID      `json:"id" swaggerignore:"true"`
	LastLogin     *time.Time     `json:"last_login,omitempty"`
	Updated       *time.Time     `json:"updated"`
}

// ToMap converts the User fields and their values to a map.
func (u *User) ToMap() (map[string]any, error) {
	var m map[string]any
	err := json.Unmarshal(support.ToJSON(u), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Validate performs validation of a User.
func (u *User) Validate(result *validation.Result) {
	var validators []validation.Validator
	validators = append(validators, constraint.Pattern{
		Name:    "email",
		Field:   u.Email,
		Expr:    support.EmailPattern.String(),
		Message: fmt.Sprintf("invalid format: %s", u.Email),
	})

	for _, v := range validators {
		v.Validate(result)
	}
}

// String returns a human-readable string representation of the User.
func (u *User) String() string {
	return string(support.ToJSONFormatted(u))
}
