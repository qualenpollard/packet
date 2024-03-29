// Package plan contains the structs and methods that each device uses.
package plan

// DataBase represents a list of all of the plans in a device.
type DataBase struct {
	Plans []Plan `json:"plans"`
}

// Plan contains the information of the Device plan for each Device
type Plan struct {
	ID              string         `json:"id"`
	Slug            string         `json:"slug"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Line            string         `json:"line"`
	Specs           Spec           `json:"specs"`
	Legacy          bool           `json:"legacy"`
	DeploymentTypes []string       `json:"deployment_types"`
	AvailableIn     []Availability `json:"available_int"`
	Class           string         `json:"class"`
	Pricing         Prices         `json:"pricing"`
}

// Spec is the struct that encompasses the information for the specification of the plan
type Spec struct {
	CPUs     []CPU   `json:"cpus"`
	Memory   Storage `json:"memory"`
	Drives   []Drive `json:"drives"`
	GPUs     []GPU   `json:"gpu"`
	Nics     []Nic   `json:"nics"`
	Features Feature `json:"features"`
}

// CPU - Central Processing Unit in the plan
type CPU struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

// GPU - Graphics Processing Unit in the plan
type GPU struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

// Storage is the amount of memory the plan contains
type Storage struct {
	Total string `json:"total"`
}

// Drive is the drives in the device
type Drive struct {
	Count int    `json:"count"`
	Size  string `json:"size"`
	Type  string `json:"type"`
}

// Nic - The Network Interface Card for the plan
type Nic struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

// Feature - the type of feature
type Feature struct {
	Raid bool `json:"raid"`
	Txt  bool `json:"txt"`
}

// Availability - Where the plan is available at
type Availability struct {
	Href string `json:"href"`
}

// Prices - The prices for the plan
type Prices struct {
	Hour float32 `json:"hour"`
}
