package schema

import (
	"time"

	"github.com/transientvariable/support-go"

	"github.com/google/uuid"

	json "github.com/json-iterator/go"
)

// Team ...
type Team struct {
	ID      uuid.UUID  `json:"id"`
	Name    string     `json:"name"`
	Creator uuid.UUID  `json:"creator"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

// ToMap converts the Team fields and their values to a map.
func (t *Team) ToMap() (map[string]any, error) {
	var m map[string]any
	err := json.Unmarshal(support.ToJSON(t), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// String returns a human-readable string representation of the Team.
func (t *Team) String() string {
	return string(support.ToJSONFormatted(t))
}

// TeamUser defines the attributes for a single mapping of a User to a Team.
type TeamUser struct {
	ID      uuid.UUID  `json:"id"`
	UserID  uuid.UUID  `json:"user_id"`
	TeamID  uuid.UUID  `json:"team_id"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

// ToMap converts the TeamUser fields and their values to a map.
func (t *TeamUser) ToMap() (map[string]any, error) {
	var m map[string]any
	err := json.Unmarshal(support.ToJSON(t), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// String returns a human-readable string representation of the TeamUser mapping.
func (t *TeamUser) String() string {
	return string(support.ToJSONFormatted(t))
}
