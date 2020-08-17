package govultr

type IPv4 struct {
	IP      string `json:"ip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Type    string `json:"type,omitempty"`
	Reverse string `json:"reverse,omitempty"`
}

type IPv6 struct {
	IP          string `json:"ip,omitempty"`
	Network     string `json:"network,omitempty"`
	NetworkSize int    `json:"network_size,omitempty"`
	Type        string `json:"type,omitempty"`
}

type ipv4sBase struct {
	IPv4S []IPv4 `json:"ipv4s"`
	Meta  *Meta  `json:"meta"`
}

type ipv6sBase struct {
	IPv6S []IPv6 `json:"ipv6s"`
	Meta  *Meta  `json:"meta"`
}
