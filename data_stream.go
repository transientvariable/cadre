package cadre

import "fmt"

// DataStream defines the attributes for uniquely identifying data streams. Attributes values are combined into the
// following canonical form: `{DataStream.Type}-{DataStream.Dataset}-{DataStream.Namespace}`
type DataStream struct {
	Type      string `json:"type"`
	Dataset   string `json:"dataset"`
	Namespace string `json:"namespace"`
}

// String returns a string representing the canonical form of the DataStream.
func (d DataStream) String() string {
	return fmt.Sprintf("%s-%s-%s", d.Type, d.Dataset, d.Namespace)
}
