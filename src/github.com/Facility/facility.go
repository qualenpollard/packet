package facility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

// DataBase ...
type DataBase struct {
	Facilities []Facility `json:"facilities"`
}

// Facility ...
type Facility struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Code     string      `json:"code"`
	Features []string    `json:"features"`
	Address  interface{} `json:"address"`
	IPRanges []string    `json:"ip_ranges"`
}

// GetFacilities ...
func GetFacilities(c *http.Client, token, url string) DataBase {
	// Create a new request to get the Facility data.
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

	if resp.StatusCode < 300 {
		fmt.Println("HTTP Response Status for", color.BlueString("Facility"), "GET:", resp.StatusCode, color.GreenString(http.StatusText(resp.StatusCode)))
	} else {
		fmt.Println("HTTP Response Status for", color.BlueString("Facility"), "GET:", resp.StatusCode, color.RedString(http.StatusText(resp.StatusCode)))
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
