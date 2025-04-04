package cadre

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/transientvariable/anchor"
	"github.com/transientvariable/cadre/validation"
	"github.com/transientvariable/cadre/validation/constraint"

	"github.com/google/uuid"
)

// Organization defines the properties for an organization.
type Organization struct {
	Active       bool           `json:"active"`
	Created      *time.Time     `json:"created"`
	BillingEmail sql.NullString `json:"billing_email,omitempty"`
	ID           uuid.UUID      `json:"id" swaggerignore:"true"`
	Name         string         `json:"name,omitempty"`
	StoragePath  []string       `json:"storage_path,omitempty"`
	Updated      *time.Time     `json:"updated"`
}

// Validate performs validation of an Organization.
func (o *Organization) Validate(result *validation.Result) {
	var validators []validation.Validator
	for _, sp := range o.StoragePath {
		validators = append(validators, constraint.Pattern{
			Name:    "Storage filePath",
			Field:   sp,
			Expr:    anchor.StoragePathPattern.String(),
			Message: fmt.Sprintf("invalid format for storage path: %s", sp),
		})
	}

	for _, v := range validators {
		v.Validate(result)
	}
}

// String returns a string representation of the Organization.
func (o *Organization) String() string {
	return string(anchor.ToJSONFormatted(o))
}

// OrganizationUser ...
type OrganizationUser struct {
	OrgID  uuid.UUID `json:"org_id"`
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}

// String returns a string representation of the OrganizationUser.
func (o *OrganizationUser) String() string {
	return string(anchor.ToJSONFormatted(o))
}
