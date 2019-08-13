package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/device"
	"github.com/fatih/color"
)

/**
 * Environment variables
 * rootURL - api.packet.net
 * userToken - API auth token
 * projID - the id of the project
 */
var rootURL = ""
var userToken = ""
var projID = ""

/**
 * Will get the environment variables for the project.
 * API Authentication Token
 * Project UUID
 */
func initEnvVariables() {
	// Get the Root URL.
	fmt.Println("Checking for Root URL...")
	url, urlInit := os.LookupEnv("ROOTURL")
	if !urlInit {
		log.Fatalln(color.RedString("No Root URL present"))
	}
	rootURL = url
	fmt.Println(color.GreenString("Received Root URL"))

	// Get the authorization token.
	fmt.Println("Checking for Auth Token...")
	authToken, tokenInit := os.LookupEnv("AUTHTOKEN")
	if !tokenInit {
		log.Fatalln(color.RedString("No Auth token present"))
	}
	userToken = authToken
	fmt.Println(color.GreenString("Received Auth Token"))

	// Get the project uuid.
	fmt.Println("Checking for Project ID...")
	id, idInit := os.LookupEnv("PROJECTUUID")
	if !idInit {
		log.Fatalln(color.RedString("No Project ID present"))
	}
	projID = id
	fmt.Println(color.GreenString("Received Project UUID"))
}

func main() {
	// Saying hello!
	fmt.Println(color.YellowString("Hello Packet!!!"))

	//Initialize Environment variables.
	initEnvVariables()

	// Create new HTTP client
	client := &http.Client{}

	// Get all of the Devices in the project.
	deviceList := device.RetrieveDevices(client, userToken, rootURL+"/projects/"+projID+"/devices")
	ipv4Addr := net.ParseIP(deviceList.Devices[0].IPAddresses[0].Address)
	fmt.Println(ipv4Addr)

	/**
	 * Use the HTTP Client and authentication token
	 * to get the available operating systems.
	 */
	// osData := packetos.GetOSes(client, userToken, rootURL+"/operating-systems")
	// fmt.Println(osData.OSes[0].Name)

	/**
	 * Use the HTTP Client and authentication token
	 * to get the available facilities.
	 */
	// facilityData := facility.GetFacilities(client, userToken, rootURL+"/facilities")
	// fmt.Println(facilityData.Facilities[0].Name)

	/**
	 * Use the HTTP Client and authentication token
	 * to get the available plans.
	 */
	// planData := plan.GetPlans(client, userToken, rootURL+"/plans")
	// fmt.Println(planData.Plans[0].Name)

	/**
	 * Use the HTTP Client and authentication token
	 * to get the available plans.
	 *
	 * This method was made when I created my own SSH Key for the device provisioning.
	 */
	// keyData := keys.GetProjectSSHKeys(client, userToken, rootURL+"/projects/"+projID+"/ssh-keys")
	// fmt.Println(keyData.Keys[0].Key)

	// Post a Device
	// deviceData := device.CreateDevice(client, userToken, rootURL+"/projects/"+projID+"/devices",
	// 	"www.QualenPollard.com",
	// 	facilityData.Facilities[0].Code,
	// 	planData.Plans[0].Slug,
	// 	osData.OSes[0].Slug,
	// 	keys.ToString(keyData.Keys))
	// fmt.Println(deviceData)
}
