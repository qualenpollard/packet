package device

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

// CreateDevice ...
// func CreateDevice(c *http.Client, token, url, host, facility, plan, packetos string, keys []string) DataBase {
// 	// Create a new Device
// 	dev := &Device{
// 		Hostname:        host,
// 		Facility:        facility,
// 		Plan:            plan,
// 		OperatingSystem: packetos,
// 		ProjectSSHkeys:  keys}

// 	// Encode the device data.
// 	reqBody, encodeErr := json.Marshal(dev)
// 	if encodeErr != nil {
// 		log.Fatalln(encodeErr)
// 	}

// 	// Create the request to post.
// 	req, reqErr := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
// 	if reqErr != nil {
// 		log.Fatalln(reqErr)
// 	}

// 	// Set the authentication token for the API
// 	req.Header.Set("X-Auth-Token", token)
// 	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
// 	resp, respErr := c.Do(req)
// 	if respErr != nil {
// 		log.Fatalln(respErr)
// 	}

// 	if resp.StatusCode < 300 {
// 		fmt.Println("HTTP Response Status for", color.BlueString("Device"), "POST:", resp.StatusCode, color.GreenString(http.StatusText(resp.StatusCode)))
// 	} else {
// 		fmt.Println("HTTP Response Status for", color.BlueString("Device"), "POST:", resp.StatusCode, color.RedString(http.StatusText(resp.StatusCode)))
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var db DataBase
// 	err = json.Unmarshal(body, &db)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return db
// }

// RetrieveDevices ...
func RetrieveDevices(c *http.Client, token, url string) DataBase {
	// Create a new request to get the Device data.
	req, reqErr := http.NewRequest("GET", url, nil)
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
		fmt.Println("HTTP Response Status for", color.BlueString("Device"), "GET:", resp.StatusCode, color.GreenString(http.StatusText(resp.StatusCode)))
	} else {
		fmt.Println("HTTP Response Status for", color.BlueString("Device"), "GET:", resp.StatusCode, color.RedString(http.StatusText(resp.StatusCode)))
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

// UpdateDevice ...
func UpdateDevice(c *http.Client, token, url, facility, plan, packetos string) []byte {
	bit := []byte{}
	return bit
}

// DeleteDevice ...
func DeleteDevice(c *http.Client, token, url, facility, plan, packetos string) []byte {
	bit := []byte{}
	return bit
}
