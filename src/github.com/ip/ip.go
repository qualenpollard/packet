package ip

import "github.com/facility"

// Body ...
type Body struct {
	ID            string            `json:"id"`
	AddressFamily float32           `json:"address_family"`
	NetMask       string            `json:"netmask"`
	CreatedAt     string            `json:"created_at"`
	Details       interface{}       `json:"details"`
	Tags          interface{}       `json:"tags"`
	Public        bool              `json:"public"`
	CIDR          float32           `json:"cidr"`
	Management    bool              `json:"management"`
	Manageable    bool              `json:"manageable"`
	Enabled       bool              `json:"enabled"`
	GlobalIP      interface{}       `json:"global_ip"`
	CustomData    interface{}       `json:"customdata"`
	Project       interface{}       `json:"project"`
	ProjectTitle  string            `json:"projcct_title"`
	Facility      facility.Facility `json:"facility"`
	AssignedTo    interface{}       `json:"assigned_to"`
	Interface     interface{}       `json:"interface"`
	Network       string            `json:"network"`
	Address       string            `json:"address"`
	Gateway       string            `json:"gateway"`
	Href          string            `json:"href"`
	Type          interface{}       `json:"type"`
}
