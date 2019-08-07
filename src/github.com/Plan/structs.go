package plan

// DataBase ...
type DataBase struct {
	Plans []Plan `json:"plans"`
}

// Plan ...
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

// Spec ...
type Spec struct {
	CPUs     []CPU   `json:"cpus"`
	Memory   Storage `json:"memory"`
	Drives   []Drive `json:"drives"`
	GPUs     []GPU   `json:"gpu"`
	Nics     []Nic   `json:"nics"`
	Features Feature `json:"features"`
}

// CPU ...
type CPU struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

// GPU ...
type GPU struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

// Storage ...
type Storage struct {
	Total string `json:"total"`
}

// Drive ...
type Drive struct {
	Count int    `json:"count"`
	Size  string `json:"size"`
	Type  string `json:"type"`
}

// Nic ...
type Nic struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

// Feature ...
type Feature struct {
	Raid bool `json:"raid"`
	Txt  bool `json:"txt"`
}

// Availability ...
type Availability struct {
	Href string `json:"href"`
}

// Prices ...
type Prices struct {
	Hour float32 `json:"hour"`
}
