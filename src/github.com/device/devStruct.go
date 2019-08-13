package device

import (
	"github.com/facility"
	"github.com/ip"
	"github.com/packetos"
	"github.com/plan"
)

// DataBase ...
type DataBase struct {
	Devices []Device `json:"devices"`
}

// Device ...
type Device struct {
	ID                    string                   `json:"id"`
	Facility              facility.Facility        `json:"facility"`
	Plan                  plan.Plan                `json:"plan"`
	Hostname              string                   `json:"hostname"`
	Description           string                   `json:"description"`
	BillingCycle          string                   `json:"billing_cycle"`
	OperatingSystem       packetos.OperatingSystem `json:"operating_system"`
	AlwaysPXE             bool                     `json:"always_pxe"`
	IPXEScriptURL         string                   `json:"ipxe_script_url"`
	UserData              string                   `json:"userdata"`
	Locked                bool                     `json:"locked"`
	CustomData            interface{}              `json:"customdata"`
	HWReservationID       string                   `json:"hardware_reservation_id"`
	SpotInstance          bool                     `json:"spot_instance"`
	SpotPriceMax          float32                  `json:"spot_price_max"`
	TerminationTime       string                   `json:"termination_time"`
	Tags                  []string                 `json:"tags"`
	ProjectSSHkeys        []string                 `json:"project_ssh_keys"`
	UserSSHkeys           []string                 `json:"user_ssh_keys"`
	Features              []string                 `json:"features"`
	PublicIPV4SubnetSize  float32                  `json:"public_ipv4_subnet_size"`
	PrivateIPV4SubnetSize float32                  `json:"private_ipv4_subnet_size"`
	IPAddresses           []ip.Body                `json:"ip_addresses"`
}
