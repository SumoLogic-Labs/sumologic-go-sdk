package types

type AllowlistingStatus struct {
	// Whether service allowlisting is enabled for Content.
	ContentEnabled bool `json:"contentEnabled"`
	// Whether service allowlisting is enabled for Login.
	LoginEnabled bool `json:"loginEnabled"`
}

type Cidr struct {
	// The string representation of the CIDR notation or IP address.
	Cidr string `json:"cidr"`
	// Description of the CIDR notation or IP address.
	Description string `json:"description,omitempty"`
}

type CidrList struct {
	// An array of CIDR notations and/or IP addresses.
	Data []Cidr `json:"data"`
}
