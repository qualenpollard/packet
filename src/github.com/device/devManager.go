package device

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

// ChangeState returns true if the action was complete, false otherwise
func ChangeState(client *http.Client, token, url, devID, action string) bool {
	if action == "TurnOn" {
		performAction(client, token, url, devID, "power_on")
		return true
	} else if action == "TurnOff" {
		performAction(client, token, url, devID, "power_off")
		return true
	}

	return false
}

// CheckState returns true if the device is active, false otherwise.
func CheckState(d *Device) bool {
	if d.State == "inactive" {
		return false
	}

	return true
}

// RetrieveDevices ...
func RetrieveDevices(c *http.Client, token, url, projectID string) (*DataBase, error) {
	// Create a new request to get the Device data.
	req, reqErr := http.NewRequest("GET", url+"/projects/"+projectID+"/devices", nil)
	if reqErr != nil {
		return nil, reqErr
	}

	// Seth the authentication token for the API
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-type", "application/json")
	resp, respErr := c.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode < 300 {
		fmt.Println("HTTP Response Status for", color.BlueString("Device \"Retrieve Devices"), "GET:", resp.StatusCode, color.GreenString(http.StatusText(resp.StatusCode)))
	} else {
		fmt.Println("HTTP Response Status for", color.BlueString("Device \"Retrieve Devices"), "GET:", resp.StatusCode, color.RedString(http.StatusText(resp.StatusCode)))
	}

	defer resp.Body.Close()

	// Get the response from the request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var db DataBase
	err = json.Unmarshal(body, &db)
	if err != nil {
		return nil, err
	}

	return &db, nil
}

// Retrieve ...
func Retrieve(c *http.Client, token, url, projID, deviceID string) (*Device, error) {

	// Create request to get the device.
	req, reqErr := http.NewRequest("GET", url+"/devices/"+deviceID, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	// Set the headers for authentication and the content-type
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-type", "application/json")
	resp, respErr := c.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode < 300 {
		fmt.Println("HTTP Response Status for", color.BlueString("Device \"Retrieve\""), "GET:", resp.StatusCode, color.GreenString(http.StatusText(resp.StatusCode)))
	} else {
		fmt.Println("HTTP Response Status for", color.BlueString("Device \"Retrieve\""), "GET:", resp.StatusCode, color.RedString(http.StatusText(resp.StatusCode)))
	}

	defer resp.Body.Close()

	// Get the body from the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Decode the json data into the device.
	d := &Device{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Fatalln(err)
	}

	return d, nil
}

// UpdateDevice ...
func UpdateDevice(c *http.Client, token, url string) []byte {
	bit := []byte{}
	return bit
}

// DeleteDevice ...
func DeleteDevice(c *http.Client, token, url string) []byte {
	bit := []byte{}
	return bit
}

// PerformAction ...
func performAction(c *http.Client, token, url, deviceID, action string) {

	// Encode the device data.
	reqBody, encodeErr := json.Marshal(map[string]string{"type": action})
	if encodeErr != nil {
		log.Fatalln(encodeErr)
	}

	// Create the request to post.
	req, reqErr := http.NewRequest("POST", url+"/devices/"+deviceID+"/actions", bytes.NewBuffer(reqBody))
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	// Set the authentication token for the API
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-type", "application/json")
	resp, respErr := c.Do(req)
	if respErr != nil {
		log.Fatalln(respErr)
	}

	if resp.StatusCode < 300 {
		fmt.Println("HTTP Response Status for", color.BlueString("Device \"Perform Action: "+action+"\""), "POST:", resp.StatusCode, color.GreenString(http.StatusText(resp.StatusCode)))
	} else {
		fmt.Println("HTTP Response Status for", color.BlueString("Device \"Perform Action: "+action+"\""), "POST:", resp.StatusCode, color.RedString(http.StatusText(resp.StatusCode)))
	}

	resp.Body.Close()
}

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
