// Package will get the environment variables for http requests to the API and start the device that will host the server.
package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/device"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

// Variables used for API requests
var (
	rootURL   = os.Getenv("ROOTURL")
	userToken = os.Getenv("AUTHTOKEN")
	projID    = os.Getenv("PROJECTUUID")
	client    = &http.Client{Timeout: time.Duration(10 * time.Second)}
	validPath = regexp.MustCompile("^/(Devices)/([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})*")
	templates = template.Must(template.ParseFiles("html/index.html", "html/deviceInfo.html"))
)

/**
 *
 *
 * TODO: Change the font in the buttons to match more like Packet
 * TODO: Finish formatting the info for the devices
 * TODO: Set up the action to activate and figure out how to check the value in javascript
 *
 *
 */

// Verification of the environment variables
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

func gracefulShutDown(sigs chan os.Signal, srv chan *http.Server, done chan bool /*devID chan string*/) {

	<-sigs
	s := <-srv
	// dID := <-devID
	fmt.Println()
	fmt.Println(color.HiRedString("Graceful Shutdown Started..."))

	// We received an interrupt signal, shut down.
	if err := s.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}

	/**
	 * Dead code from a previous idea that I was approaching.
	 */
	// d := device.Retrieve(c, userToken, rootURL, dID)

	// // If the device is active, power it off, else, do nothing.
	// if d.State == "active" {

	// 	log.Println(color.BlueString("Powering device off..."))
	// 	device.ChangeState(c, userToken, rootURL, dID, "TurnOff")

	// 	// Check for the device to be inactive every 5 seconds.
	// 	duration := time.Duration(5) * time.Second
	// 	for d.State != "inactive" {
	// 		time.Sleep(duration)
	// 		d = device.Retrieve(c, userToken, rootURL, dID)
	// 	}
	// }

	// fmt.Println(color.RedString("Device inactive"))
	fmt.Println(color.HiRedString("Graceful Shutdown Complete."))
	done <- true
}

func indexHandler(w http.ResponseWriter, r *http.Request, title string) {
	devs, err := device.RetrieveDevices(client, userToken, rootURL, projID)
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.ExecuteTemplate(w, "index.html", &devs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func deviceInfoHandler(w http.ResponseWriter, r *http.Request, id string) {

	fmt.Println(id)
	dev, err := device.Retrieve(client, userToken, rootURL, projID, id)
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.ExecuteTemplate(w, "deviceInfo.html", dev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func renderTemplate(w http.ResponseWriter, tmpl string, d *device.Device) {
// 	err := templates.ExecuteTemplate(w, tmpl+".html", d)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {

			io.WriteString(w, "Can't find: "+r.URL.Path)
			return
		} else if m[2] == "" {
			fn(w, r, m[1])
		} else {
			fn(w, r, m[2])
		}
	}
}

func main() {

	// Saying hello!
	fmt.Println(color.YellowString("Hello Packet!!!"))

	// Create channels for graceful shutdown
	sigs := make(chan os.Signal, 1)
	srv := make(chan *http.Server, 1)
	done := make(chan bool, 1)
	// id := make(chan string, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	/**
	* Create a new router and a new HTTP Server.
	 */
	r := mux.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r}
	srv <- server

	go gracefulShutDown(sigs, srv, done /*id*/)

	//Index page
	r.HandleFunc("/Devices/", makeHandler(indexHandler))
	r.HandleFunc("/Devices/{ID}", makeHandler(deviceInfoHandler))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	/**
	 * Dead code from a previous idea that I was approaching.
	 */
	// Choose a device if there is one available and power it on.
	// d := getDevice(client, id)
	// ipAddrList := ip.RetrieveDeviceIPAddresses(client, userToken, rootURL, d.ID)
	// ipAddr := ipAddrList.Addresses[0].Network
	// fmt.Println("IP Address: ", ipAddr)
	// server.Addr = strings.Replace(server.Addr, ":", ipAddr+":", -1)

	log.Println(color.YellowString("Listing to port: " + server.Addr))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
	<-done
	fmt.Println("Exiting...")
}
