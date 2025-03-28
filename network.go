package schema

// Network represents metadata for a communication path over which a host or network event happens.
type Network struct {
	// Application is the name given to an application level protocol. This can be arbitrarily assigned for things like
	// microservices, but also apply to things like skype, icq, facebook, twitter. This would be used in situations
	// where the vendor or service can be decoded such as from the source/dest IP owners, ports, or wire format.
	Application string `json:"application"`

	// Bytes Total bytes transferred in both directions.
	Bytes int64 `json:"bytes"`

	// CommunityID is a hash of source and destination IPs and ports, as well as the protocol  used in a communication.
	// This is a tool-agnostic standard to identify flows. See: https://github.com/corelight/community-id-spec.
	CommunityID string `json:"community_id"`

	// Direction of the network traffic.
	//
	// Recommended values:
	//
	//   * ingress
	//   * egress
	//   * inbound
	//   * outbound
	//   * internal
	//   * external
	//   * unknown
	//
	// When mapping events from a host-based monitoring context, populate this field from the host's point of view,
	// using the values "ingress" or "egress".
	//
	// When mapping events from a network or perimeter-based monitoring context, populate this field from the point of
	// view of the network perimeter, using the values "inbound", "outbound", "internal" or "external".
	//
	// Note that "internal" is not crossing perimeter boundaries, and is meant to describe communication between two
	// hosts within the perimeter. Note also that "external" is meant to describe traffic between two hosts that
	// are external to the perimeter. This could for example be useful for ISPs  or VPN service providers.
	Direction string `json:"direction"`

	// ForwardedIP is the host IP address when the source IP address is the proxy.
	ForwardedIP string `json:"forwarded_ip"`

	// IANANumber is the IANA Protocol Number (https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	// This aligns well with NetFlow and sFlow  related logs which use the IANA protocol number.
	IANANumber string `json:"iana_number"`

	// Inner are the fields are added in addition to network.vlan fields to describe the innermost VLAN when q-in-q VLAN
	// tagging is present. Allowed fields include vlan.id and vlan.name. Inner vlan fields are typically used when
	// sending traffic with multiple 802.1q encapsulations to a network sensor (e.g. Zeek, Wireshark.)
	Inner map[string]any `json:"inner"`

	// Name given by operators to sections of their network.
	Name string `json:"name"`

	// Packets is the total packets transferred in both directions.
	Packets int64 `json:"packets"`

	// Protocol is the L7 network protocol name (e.g. support, lumberjack).
	Protocol string `json:"protocol"`

	// Transport is same as Network.IANANumber, but instead using the Keyword name of the transport layer
	// (udp, tcp, ipv6-icmp, etc.)
	Transport string `json:"transport"`

	// Type is the network layer of the OSI Model (ipv4, ipv6, ipsec, pim, etc.).
	Type string `json:"type"`
}
