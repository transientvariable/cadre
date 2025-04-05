package ecs

// Group defines the properties for a group relevant to an event.
type Group struct {
	// ID is the unique identifier for the group on the system/platform.
	ID string `json:"id,omitempty"`

	// Name is the name of the group.
	Name string `json:"name,omitempty"`

	// Domain is the name of the directory the group is a member of.
	//
	// For example, an LDAP or Active Directory domain name.
	Domain string `json:"domain,omitempty"`
}
