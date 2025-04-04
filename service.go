package cadre

// Service fields for describing the service for or from which the data was collected.
type Service struct {
	Address     string `json:"address,omitempty"`
	Environment string `json:"environment,omitempty"`
	EphemeralID string `json:"ephemeral_id,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	NodeName    string `json:"node.name,omitempty"`
	State       string `json:"state,omitempty"`
	Type        string `json:"type,omitempty"`
	Version     string `json:"version,omitempty"`
}
