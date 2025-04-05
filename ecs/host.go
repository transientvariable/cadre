package ecs

// Host defines the attributes for a general computing instance that details about the host on which
// an event happened, or from which a measurement was taken.
//
// A Host can be hardware, virtual machine, Docker container, and Kubernetes node.
type Host struct {
	// Architecture is the operating system architecture for the Host.
	Architecture string `json:"architecture,omitempty"`

	// CpuUsage is the percent CPU used which is normalized by the number of CPU cores within a range from 0 to 1 and
	// has scaling factor of 1000.
	//
	// For example: For a two core host, this value should be the average of the two cores, between 0 and 1.
	CpuUsage float64 `json:"cpu.usage,omitempty"`

	// DiskReadBytes is the total number of bytes (gauge) read successfully (aggregated from all disks) since the last
	// metric collection.
	DiskReadBytes int64 `json:"disk.read.bytes,omitempty"`

	// DiskWriteBytes is the total number of bytes (gauge) written successfully (aggregated from all disks) since the
	// last metric collection.
	DiskWriteBytes int64 `json:"disk.write.bytes,omitempty"`

	// Domain is the domain name of which the Host is a member.
	//
	// For example, on Windows this could be the host's Active Directory domain or NetBIOS domain name. For Linux this
	// could be the domain of the host's LDAP provider.
	Domain string `json:"domain,omitempty"`

	// Hostname of the host.
	Hostname string `json:"hostname,omitempty"`

	// ID is a unique identifier for the Host.
	ID string `json:"id,omitempty"`

	// IP address of the Host.
	IP string `json:"ip,omitempty"`

	// MAC address for the Host.
	//
	// The notation format from RFC 7042 is suggested: Each octet (that is, 8-bit byte) is represented by two
	// [uppercase] hexadecimal digits giving the value of the octet as an unsigned integer. Successive octets are
	// separated by a hyphen.
	MAC string `json:"mac,omitempty"`

	// Name of the host.
	Name string `json:"name,omitempty"`

	// NetworkIngressBytes is the number of bytes received (gauge) on all network interfaces by the host since the last
	// metric collection.
	NetworkIngressBytes int64 `json:"network.ingress.bytes,omitempty"`

	// NetworkIngressPackets is the number of packets (gauge) received on all network interfaces by the host since the
	// last metric collection.
	NetworkIngressPackets int64 `json:"network.ingress.packets,omitempty"`

	// NetworkEgressBytes is the number of bytes (gauge) sent out on all network interfaces by the host since the last
	// metric collection.
	NetworkEgressBytes int64 `json:"network.egress.bytes,omitempty"`

	// NetworkEgressPackets is the number of packets (gauge) sent out on all network interfaces by the host since the
	// last metric collection.
	NetworkEgressPackets int64 `json:"network.egress.packets,omitempty"`

	// Type the type of Host.
	//
	// For Cloud providers this can be the machine type (e.g. `t2.medium`). If vm, this could be the container, for
	// example, or other information meaningful in your environment.
	Type string `json:"type,omitempty"`

	// Uptime the time in seconds the Host has been up.
	Uptime int64 `json:"uptime,omitempty"`
}
