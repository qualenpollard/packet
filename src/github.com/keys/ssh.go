// Package keys contains the structs and methods for the keys of a project
package keys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

// Key is the information that describes the key for a project
type Key struct {
	ID          string      `json:"id"`
	Label       string      `json:"label"`
	Key         string      `json:"key"`
	Fingerprint string      `json:"fingerprint"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
	Owner       interface{} `json:"owner"`
	URL         string      `json:"href"`
}

// DataBase is a list of Keys
type DataBase struct {
	Keys []Key `json:"ssh_keys"`
}

/**
 * GetProjectSSHKeys makes a request to get the keys of a specific project
 * Returns Database - a list of Key
 */
func GetProjectSSHKeys(c *http.Client, token, url string) DataBase {
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
		fmt.Println("HTTP Response Status for", color.BlueString("SSH Keys:"), "GET:", resp.StatusCode, color.GreenString(http.StatusText(resp.StatusCode)))
	} else {
		fmt.Println("HTTP Response Status for", color.BlueString("SSH Keys:"), "GET:", resp.StatusCode, color.RedString(http.StatusText(resp.StatusCode)))
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

/**
 * ToString converts a list of Key to a list of String
 * Returns a slice of strings
 */
func ToString(keys []Key) []string {
	strKeys := []string{}

	for i := 0; i < len(keys); i++ {
		strKeys = append(strKeys, keys[i].Key)
	}

	return strKeys
}
