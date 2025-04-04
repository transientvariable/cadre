package cadre

import "time"

// Enumeration of event kind values.
const (
	EventKindAlert         = "alert"
	EventKindEnrichment    = "enrichment"
	EventKindEvent         = "event"
	EventKindMetric        = "metric"
	EventKindState         = "state"
	EventKindPipelineError = "pipeline_error"
	EventKindSignal        = "signal"
)

// Enumeration of event category values.
const (
	EventCategoryAuthentication = "authentication"
	EventCategoryConfiguration  = "configuration"
	EventCategoryDatabase       = "database"
	EventCategoryDriver         = "driver"
	EventCategoryEmail          = "email"
	EventCategoryFile           = "file"
	EventCategoryHost           = "host"
	EventCategoryIAM            = "iam"
	EventCategoryNetwork        = "network"
	EventCategoryPackage        = "package"
	EventCategoryProcess        = "process"
	EventCategoryRegistry       = "registry"
	EventCategorySession        = "session"
	EventCategoryWeb            = "web"
)

// Enumeration of event type values.
const (
	EventTypeAccess     = "access"
	EventTypeAdmin      = "admin"
	EventTypeAllowed    = "allowed"
	EventTypeChange     = "change"
	EventTypeConnection = "connection"
	EventTypeCreation   = "creation"
	EventTypeDeletion   = "deletion"
	EventTypeDenied     = "denied"
	EventTypeEnd        = "end"
	EventTypeError      = "error"
	EventTypeGroup      = "group"
	EventTypeInfo       = "info"
	EventTypeProtocol   = "protocol"
	EventTypeStart      = "start"
	EventTypeUser       = "user"
)

// Enumeration of event outcome values.
const (
	EventOutcomeFailure = "failure"
	EventOutcomeSuccess = "success"
	EventOutcomeUnknown = "unknown"
)

// Enumeration of event action values.
const (
	EventActionFileCreated = "file-created"
	EventActionFileMoved   = "file-changed"
	EventActionFileRemoved = "file-removed"
)

// Event defines the attributes for context information about an event.
type Event struct {
	Action   string        `json:"action,omitempty"`
	Category []string      `json:"category"`
	Code     string        `json:"code,omitempty"`
	Created  *time.Time    `json:"created,omitempty"`
	Dataset  string        `json:"dataset,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
	End      *time.Time    `json:"end,omitempty"`
	Hash     string        `json:"hash,omitempty"`
	ID       string        `json:"id,omitempty"`
	Ingested *time.Time    `json:"ingested,omitempty"`
	Kind     string        `json:"kind"`
	Module   string        `json:"module,omitempty"`
	Outcome  string        `json:"outcome,omitempty"`
	Provider string        `json:"provider,omitempty"`
	Reason   string        `json:"reason,omitempty"`
	Sequence int64         `json:"sequence,omitempty"`
	Severity int64         `json:"severity,omitempty"`
	Start    *time.Time    `json:"start,omitempty"`
	Type     []string      `json:"type"`
}
