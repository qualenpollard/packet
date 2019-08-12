package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/device"
	"github.com/facility"
	"github.com/packetos"
	"github.com/plan"
)

//ROOT URL and Auth token
var rootURL = "https://api.packet.net"
var userToken = ""

func main() {
	// Saying hello!
	fmt.Println("Hello Packet!")

	// Get the authorization token.
	fmt.Println("Checking for Auth Token...")
	authToken, tokenInit := os.LookupEnv("AUTHTOKEN")
	if !tokenInit {
		log.Fatalln("No Auth token present")
	}
	userToken = authToken
	fmt.Println("Received Auth Token.")

	// Create new HTTP client
	client := &http.Client{}

	/**
	 * Use the HTTP Client and authentication token
	 * to get the available operating systems.
	 */
	osData := packetos.GetOSes(client, userToken, rootURL+"/operating-systems")

	// /**
	// * Use the HTTP Client and authentication token
	// * to get the available facilities.
	//  */
	facilityData := facility.GetFacilities(client, userToken, rootURL+"/facilities")

	// /**
	// * Use the HTTP Client and authentication token
	// * to get the available plans.
	//  */
	planData := plan.GetPlans(client, userToken, rootURL+"/plans")

	// Post a Device
	deviceData := device.CreateDevice(client, userToken, rootURL+"/projects/c69d830d-d2d4-402b-bdb8-403cd9196f6a/devices",
		facilityData.Facilities[0].Name,
		planData.Plans[0].Name,
		osData.OSes[0].Name)
	fmt.Println(deviceData)
}
