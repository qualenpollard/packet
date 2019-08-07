package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	facility "github.com/Facility"
	"github.com/OS"
	plan "github.com/Plan"
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
	fmt.Println(osData.OSes[0])

	/**
	* Use the HTTP Client and authentication token
	* to get the available facilities.
	 */
	facilityData := facility.GetFacilities(client, userToken, rootURL+"/facilities")
	fmt.Println(facilityData.Facilities[0])

	/**
	* Use the HTTP Client and authentication token
	* to get the available plans.
	 */
	planData := plan.GetPlans(client, userToken, rootURL+"/plans")
	fmt.Println(planData.Plans[0])

	// Create a Device

}
