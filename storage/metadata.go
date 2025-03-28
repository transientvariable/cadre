package storage

import (
	"github.com/transientvariable/validation"
	"github.com/transientvariable/validation/constraint"
	"strings"

	json "github.com/json-iterator/go"
	"github.com/transientvariable/schema"
	"github.com/transientvariable/support"
)

const (
	IndexPrefixMetadataStorage = "metadata-storage-"
	MetadataLabelKeyNamespace  = "namespace"
	NamespaceFragmentUpload    = ".upload"
)

// Metadata ...
type Metadata struct {
	schema.Base
	File *schema.File `json:"file,omitempty"`

	id        string
	namespace string
}

// NewMetadata ...
func NewMetadata(namespace string, file *schema.File) (*Metadata, error) {
	metadata := &Metadata{
		File:      file,
		id:        fileID(file),
		namespace: strings.TrimSpace(namespace),
	}

	if result := metadata.validate(); !result.IsValid() {
		return nil, result
	}

	metadata.Timestamp = file.Ctime
	metadata.Labels = map[string]any{MetadataLabelKeyNamespace: namespace}
	return metadata, nil
}

// ID ...
func (m *Metadata) ID() string {
	return m.id
}

// Namespace ...
func (m *Metadata) Namespace() string {
	return m.namespace
}

// ToMap converts the Metadata fields and their values to a map.
func (m *Metadata) ToMap() (map[string]any, error) {
	var mm map[string]any
	err := json.Unmarshal(support.ToJSON(m), &mm)
	if err != nil {
		return nil, err
	}
	return mm, nil
}

// String returns a string representation of the Metadata.
func (m *Metadata) String() string {
	mm := make(map[string]any)
	mm["id"] = m.id
	mm["metadata"] = m
	mm["namespace"] = m.namespace
	return string(support.ToJSONFormatted(mm))
}

// validate performs validation of a storage Metadata.
func (m *Metadata) validate() *validation.Result {
	var validators []validation.Validator
	validators = append(validators, constraint.NotBlank{
		Name:    "namespace",
		Field:   m.namespace,
		Message: "metadata: namespace is required",
	})

	validators = append(validators, validation.ValidatorFunc(func(result *validation.Result) {
		if m.File == nil {
			result.Add("file", "metadata: file is required")
		}
	}))

	validators = append(validators, constraint.NotBlank{
		Name:    "metadataID",
		Field:   m.id,
		Message: "metadata: ID is required",
	})
	return validation.Validate(validators...)
}
