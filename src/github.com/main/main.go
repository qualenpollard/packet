// Package will get the environment variables for http requests to the API and start the device that will host the server.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/device"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/ip"
)

// Variables used for API requests
var (
	rootURL   = os.Getenv("ROOTURL")
	userToken = os.Getenv("AUTHTOKEN")
	projID    = os.Getenv("PROJECTUUID")
)

// Verify if the environment variables were set
func init() {
	// Get the Root URL.
	fmt.Print("Checking for Root URL... ")
	if rootURL == "" {
		log.Fatalln(color.RedString("No Root URL present"))
	}
	fmt.Println(color.GreenString("Initialized"))

	// Get the authorization token.
	fmt.Print("Checking for Auth Token... ")
	if userToken == "" {
		log.Fatalln(color.RedString("No Auth token present"))
	}
	fmt.Println(color.GreenString("Initialized"))

	// Get the project uuid.
	fmt.Print("Checking for Project ID... ")
	if projID == "" {
		log.Fatalln(color.RedString("No Project ID present"))
	}
	fmt.Println(color.GreenString("Initialized"))
}

func gracefulShutDown(sigs chan os.Signal, srv chan *http.Server, done chan bool, cli chan *http.Client, devID chan string) {

	<-sigs
	c := <-cli
	dID := <-devID
	s := <-srv
	fmt.Println()
	fmt.Println("Graceful Shutdown Started...")

	// We received an interrupt signal, shut down.
	if err := s.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}

	d := device.Retrieve(c, userToken, rootURL, dID)

	// If the device is active, power it off, else, do nothing.
	if d.State == "active" {

		log.Println(color.BlueString("Powering device off..."))
		device.ChangeState(c, userToken, rootURL, dID, "TurnOff")

		// Check for the device to be inactive every 5 seconds.
		duration := time.Duration(5) * time.Second
		for d.State != "inactive" {
			time.Sleep(duration)
			d = device.Retrieve(c, userToken, rootURL, dID)
		}
	}

	fmt.Println(color.RedString("Device inactive"))
	done <- true
}

func getDevice(c *http.Client, id chan string) device.Device {

	// Get all of the Devices in the project.
	log.Println("Getting all devices in project.")
	deviceList := device.RetrieveDevices(c, userToken, rootURL, projID)
	if len(deviceList.Devices) == 0 {
		log.Fatalln(color.RedString("There are no devices available."))
	}

	// Seclect device to power on, channels device.ID to gracefulShutDown()
	log.Println("Getting device ID")
	devID := deviceList.Devices[0].ID
	id <- devID
	fmt.Println("Device ID: ", devID)

	/**
	 * Get the device that'll be used as the server and turn it on
	 * if it's not already active.
	 */
	d := device.Retrieve(c, userToken, rootURL, devID)
	active := device.CheckState(d)
	if !active {
		fmt.Println(color.BlueString("Powering on device..."))
		stateChanged := device.ChangeState(c, userToken, rootURL, devID, "TurnOn")
		if !stateChanged {
			log.Fatalln(color.RedString("Turn on device incomplete."))
		}

		// Check for the device to be active every 5 seconds.
		duration := time.Duration(5) * time.Second
		for d.State != "active" {
			time.Sleep(duration)
			d = device.Retrieve(c, userToken, rootURL, devID)
		}

		log.Println(color.GreenString("Device Active"))
	}

	return *d
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is test")
}

func main() {

	// Saying hello!
	fmt.Println(color.YellowString("Hello Packet!!!"))

	// Create channels for graceful shutdown
	sigs := make(chan os.Signal, 1)
	srv := make(chan *http.Server, 1)
	done := make(chan bool, 1)
	c := make(chan *http.Client, 1)
	id := make(chan string, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	/**
	* Create a new router and a new HTTP client/Server.
	* Channels client into gracefulShutDown()
	 */
	r := mux.NewRouter()
	client := &http.Client{}
	server := &http.Server{
		Addr:    ":80",
		Handler: r}
	c <- client
	srv <- server

	go gracefulShutDown(sigs, srv, done, c, id)

	//Index page
	r.HandleFunc("/QualenPollard", indexHandler)

	// Choose a device if there is one available and power it on.
	// Get the ip addresses for the device and set the port number for the server.
	d := getDevice(client, id)
	ipAddrList := ip.RetrieveDeviceIPAddresses(client, userToken, rootURL, d.ID)
	ipAddr := ipAddrList.Addresses[0].Network
	fmt.Println("IP Address: ", ipAddr)
	server.Addr = strings.Replace(server.Addr, ":", ipAddr+":", -1)

	log.Println(color.YellowString("Trying to listen..."))
	_, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Fatalln(status)

	log.Println(color.YellowString("Listing to port: " + server.Addr))
	// if err := server.ListenAndServe(); err != nil {
	// 	// Error starting or closing listener:
	// 	log.Printf("HTTP server ListenAndServe: %v", err)
	// }
	<-done
	fmt.Println("Exiting...")
}
