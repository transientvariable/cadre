package ecs

type NAT struct {
	// IP is the IP of the source based NAT sessions, typically connections traversing load balancers, firewalls, or
	// routers.
	IP string `json:"ip,omitempty"`

	// Port is the port of source based NAT sessions, typically connections traversing load balancers, firewalls, or
	// routers.
	Port int `json:"port,omitempty"`
}
