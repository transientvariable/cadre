package ecs

import (
	"time"
)

// Base represents the common properties shared between composite data types. For example, for a composite type named
// `FooEvent` that represents some arbitrary event, Base would be used as follows:
//
//	   import "github.com/transientvariable/cadre/ecs"
//
//		  type FooEvent struct {
//	       ecs.Base
//	       DataStream  ecs.DataStream `json:"data_stream"`
//	       Event       ecs.Event      `json:"event"`
//	       FooFieldOne string         `json:"foo_field_one"`
//	   }
type Base struct {
	// Timestamp is the date/time when an event originated. This is the date/time extracted from the event, typically
	// representing when the event was generated by the source. If the timestamp cannot be derived from the original
	// event, the date/time the event was encountered by the pipeline should be used.
	Timestamp *time.Time `json:"@timestamp,omitempty"`

	// Tags is an optional list of keywords used to tag a schema type.
	Tags string `json:"tags,omitempty"`

	// Labels is an optional collection of key/value pairs for adding metadata to a schema type.
	Labels map[string]any `json:"labels,omitempty"`

	// Message is the message from the source event, if any.
	Message string `json:"message,omitempty"`
}
