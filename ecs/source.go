package ecs

// Source captures details about the sender of a network exchange/packet.
type Source struct {
	// Address represents the raw address.
	Address string `json:"address"`

	// Bytes sent from the source to the destination.
	Bytes int64 `json:"bytes"`

	// Domain is the Source domain.
	Domain string `json:"domain"`

	// IP address of the source (IPv4 or IPv6).
	IP string `json:"ip"`

	// MAC address of the source.
	//
	// The notation format from RFC 7042 is suggested: Each octet (that is, 8-bit byte) is represented by two
	// [uppercase] hexadecimal digits giving the value of the octet as an unsigned integer. Successive octets are
	// separated by a hyphen.
	MAC string `json:"mac"`

	// Packets sent from the source to the destination.
	Packets int64 `json:"packets"`

	// Port of the source.
	Port int64 `json:"port"`

	NAT *NAT `json:"nat,omitempty"`

	// RegisteredDomain is the highest registered Source domain, stripped of the subdomain.
	//
	// For example, the registered domain for "foo.example.com" is "example.com".
	//
	// This value can be determined precisely with a list like the public suffix list (http://publicsuffix.org). Trying
	// to approximate this by simply taking the last two labels will not work well for TLDs such as "co.uk".
	RegisteredDomain string `json:"registered_domain"`

	// TopLevelDomain is the effective top level domain (eTLD), also known as the domain suffix, is the last part of the
	// domain name.
	//
	// For example, the top level domain for example.com is "com". This value can be determined precisely with a list
	// like the public suffix list (http://publicsuffix.org). Trying to approximate this by simply taking the last label
	// will not work well for effective TLDs such as "co.uk".
	TopLevelDomain string `json:"top_level_domain"`

	// Subdomain is the subdomain portion of a fully qualified domain name which includes all the names except the host
	// name under the RegisteredDomain.
	//
	// In a partially qualified domain, or if the qualification level of the full name cannot be determined, subdomain
	// contains all the names below the registered domain.
	//
	// For example the subdomain portion of "www.east.mydomain.co.uk" is "east". If the domain has multiple levels of
	// subdomain, such as "sub2.sub1.example.com", the subdomain field should contain "sub2.sub1", with no trailing
	// period.
	Subdomain string `json:"subdomain"`
}
