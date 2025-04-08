package ecs

// Client defines the properties for an initiator of a network connection event.
//
// For TCP events, the Client is the initiator of the TCP connection that sends the SYN packet(s). For other protocols,
// the Client is generally the initiator or requestor in the network transaction. Some systems use the term "originator"
// to refer the Client in TCP connections.
//
// The Client fields describe details about the system acting as the Client in the network event. Client fields are
// usually populated in conjunction with server fields and generally not populated for packet-level events. Client /
// server representations can add semantic context to an exchange, which is helpful to visualize the data in certain
// situations.
type Client struct {
	// Address is the IP, domain or a unix socket for a Client network connection event.
	//
	// The Address should always be set to the raw address and duplicated in IP or Domain field, depending on which one
	// it is.
	Address string `json:"address,omitempty"`

	// Bytes sent from the Client to the server.
	Bytes int64 `json:"bytes,omitempty"`

	// Domain of the Client.
	Domain string `json:"domain,omitempty"`

	// IP address of the Client (IPv4 or IPv6).
	IP string `json:"ip,omitempty"`

	// MAC address of the Client.
	//
	// The notation format from RFC 7042 is suggested: Each octet (that is, 8-bit byte) is represented by two
	// [uppercase] hexadecimal digits giving the value of the octet as an unsigned integer. Successive octets are
	// separated by a hyphen.
	MAC string `json:"mac,omitempty"`

	NAT *NAT `json:"nat,omitempty"`

	// Packets sent from the Client to the server.
	Packets int64 `json:"packets,omitempty"`

	// Port of the Client.
	Port int64 `json:"port,omitempty"`

	// RegisteredDomain is the highest registered Client domain, stripped of the subdomain.
	//
	// For example, the registered domain for "foo.example.com" is "example.com".
	//
	// This value can be determined precisely with a list like the public suffix list (http://publicsuffix.org). Trying
	// to approximate this by simply taking the last two labels will not work well for TLDs such as "co.uk".
	RegisteredDomain string `json:"registered_domain,omitempty"`

	// Subdomain is the subdomain portion of a fully qualified domain name which includes all the names except the host
	// name under the RegisteredDomain.
	//
	// In a partially qualified domain, or if the qualification level of the full name cannot be determined, subdomain
	// contains all the names below the registered domain.
	//
	// For example the subdomain portion of "www.east.mydomain.co.uk" is "east". If the domain has multiple levels of
	// subdomain, such as "sub2.sub1.example.com", the subdomain field should contain "sub2.sub1", with no trailing
	// period.
	Subdomain string `json:"subdomain,omitempty"`

	// TopLevelDomain is the effective top level domain (eTLD), also known as the domain suffix, is the last part of the
	// domain name.
	//
	// For example, the top level domain for example.com is "com". This value can be determined precisely with a list
	// like the public suffix list (http://publicsuffix.org). Trying to approximate this by simply taking the last label
	// will not work well for effective TLDs such as "co.uk".
	TopLevelDomain string `json:"top_level_domain,omitempty"`
}
