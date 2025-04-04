package storage

import (
	"time"

	"github.com/transientvariable/anchor"
	"github.com/transientvariable/cadre"
	"github.com/transientvariable/cadre/validation"
	"github.com/transientvariable/cadre/validation/constraint"
)

const (
	IndexLogsEventStorage = "logs-event-storage"
)

// Event ...
type Event struct {
	cadre.Base
	Event *cadre.Event `json:"event,omitempty"`
	File  *cadre.File  `json:"file,omitempty"`

	id        string
	eventType string
	namespace string
}

// NewStorageEvent ...
func NewStorageEvent(eventType string, namespace string, file *cadre.File) (*Event, error) {
	created := time.Now().UTC()
	event := &Event{
		Event: &cadre.Event{
			ID:       fileID(file),
			Created:  &created,
			Kind:     cadre.EventKindEvent,
			Category: []string{cadre.EventCategoryFile},
			Type:     []string{eventType},
		},
		File:      file,
		eventType: eventType,
		namespace: namespace,
	}

	if result := event.validate(); !result.IsValid() {
		return nil, result
	}

	event.Timestamp = file.Ctime
	if eventType == cadre.EventTypeCreation {
		event.File.Created = file.Ctime
	}

	event.Labels = map[string]any{MetadataLabelKeyNamespace: namespace}
	return event, nil
}

// ID ...
func (e *Event) ID() string {
	return e.id
}

// Namespace ...
func (e *Event) Namespace() string {
	return e.namespace
}

// String returns a string representation of the Event.
func (e *Event) String() string {
	em := make(map[string]any)
	em["event"] = e
	em["id"] = e.id
	em["type"] = e.eventType
	em["namespace"] = e.namespace
	return string(anchor.ToJSONFormatted(em))
}

// validate performs validation of a storage Event.
func (e *Event) validate() *validation.Result {
	var validators []validation.Validator
	validators = append(validators, constraint.NotBlank{
		Name:    "eventType",
		Field:   e.eventType,
		Message: "storage_event: type is required",
	})

	validators = append(validators, constraint.NotBlank{
		Name:    "namespace",
		Field:   e.namespace,
		Message: "storage_event: namespace is required",
	})

	validators = append(validators, validation.ValidatorFunc(func(result *validation.Result) {
		if e.File == nil {
			result.Add("file", "storage_event: file is required")
		}
	}))

	if (e.eventType == cadre.EventTypeCreation && !e.File.IsDir()) || e.eventType == cadre.EventTypeChange {
		validators = append(validators, constraint.NotBlank{
			Name:    "eventID",
			Field:   e.Event.ID,
			Message: "storage_event: ID is required",
		})
	}
	return validation.Validate(validators...)
}
