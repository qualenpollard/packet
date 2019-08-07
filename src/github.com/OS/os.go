package packetos

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// OperatingSystem ...
type OperatingSystem struct {
	ID             string      `json:"id"`
	Slug           string      `json:"slug"`
	Name           string      `json:"name"`
	Distro         string      `json:"distro"`
	Version        string      `json:"version"`
	Provisionable  []string    `json:"provisionable_on"`
	Preinstallable bool        `json:"preinstallable"`
	Pricing        interface{} `json:"pricing"`
	Licensed       bool        `json:"licensed"`
}

// DataBase ...
type DataBase struct {
	OSes []OperatingSystem `json:"operating_systems"`
}

// GetOSes ...
func GetOSes(c *http.Client, token, url string) DataBase {
	// Create a new request to get the available Operating Systems.
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	// Seth the authentication token for the API
	req.Header.Set("X-Auth-Token", token)
	resp, respErr := c.Do(req)
	if respErr != nil {
		log.Fatalln(respErr)
	}

	defer resp.Body.Close()

	// // Get the response from the request
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
