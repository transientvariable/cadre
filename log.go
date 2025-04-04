package cadre

// Log defines attributes for representing details about a logging mechanism or transport.
type Log struct {
	Level              string         `json:"level,omitempty"`
	FilePath           string         `json:"file.path,omitempty"`
	Logger             string         `json:"logger,omitempty"`
	OriginFileName     string         `json:"origin.file.name,omitempty"`
	OriginFileLine     int64          `json:"origin.file.line,omitempty"`
	OriginFunction     string         `json:"origin.function,omitempty"`
	Syslog             map[string]any `json:"syslog,omitempty"`
	SyslogSeverityCode int64          `json:"syslog.severity.code,omitempty"`
	SyslogSeverityName string         `json:"syslog.severity.name,omitempty"`
	SyslogFacilityCode int64          `json:"syslog.facility.code,omitempty"`
	SyslogFacilityName string         `json:"syslog.facility.name,omitempty"`
	SyslogPriority     int64          `json:"syslog.priority,omitempty"`
}
