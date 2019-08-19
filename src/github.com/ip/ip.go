package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/facility"
	"github.com/fatih/color"
)

// DataBase ...
type DataBase struct {
	Addresses []Body `json:"ip_addresses"`
}

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

// RetrieveDeviceIPAddresses ...
func RetrieveDeviceIPAddresses(c *http.Client, token, url, devID string) DataBase {
	// Create a new request to get the Device data.
	req, reqErr := http.NewRequest("GET", url+"/devices/"+devID+"/ips/", nil)
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	// Seth the authentication token for the API
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-type", "application/json")
	resp, respErr := c.Do(req)
	if respErr != nil {
		log.Fatalln(respErr)
	}

	if resp.StatusCode < 300 {
		fmt.Println("HTTP Response Status for", color.BlueString("IPAddress"), "GET:", resp.StatusCode, color.GreenString(http.StatusText(resp.StatusCode)))
	} else {
		fmt.Println("HTTP Response Status for", color.BlueString("IPAddress"), "GET:", resp.StatusCode, color.RedString(http.StatusText(resp.StatusCode)))
	}

	defer resp.Body.Close()

	// Get the response from the request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var db DataBase
	err = json.Unmarshal(body, &db)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
