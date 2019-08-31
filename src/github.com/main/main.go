// Package will get the environment variables for http requests to the API and route paths for the Project Device dashboard web app.
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

var (
	rootURL     = os.Getenv("ROOTURL")
	userToken   = os.Getenv("AUTHTOKEN")
	projID      = os.Getenv("PROJECTUUID")
	client      = &http.Client{Timeout: time.Duration(10 * time.Second)}
	validPath   = regexp.MustCompile("^/(Devices)/([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})*")
	validAction = regexp.MustCompile("^/(Device)/([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})/(power_on|power_off)")
	templates   = template.Must(template.ParseFiles("html/index.html", "html/deviceInfo.html"))
)

/**
 * Verification of the environment variables
 */
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

/**
 * Performs a graceful shut down of the server
 */
func gracefulShutDown(sigs chan os.Signal, srv chan *http.Server, done chan bool) {

	<-sigs
	s := <-srv
	fmt.Println()
	fmt.Println(color.HiRedString("Graceful Shutdown Started..."))

	// We received an interrupt signal, shut down.
	if err := s.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}

	// fmt.Println(color.RedString("Device inactive"))
	fmt.Println(color.HiRedString("Graceful Shutdown Complete."))
	done <- true
}

/**
 * indexHandler is the handler for the route page "/Devices/" that presents all of the project devices.
 */
func indexHandler(w http.ResponseWriter, r *http.Request, title string) {

	// Retrieve all of the devices from the project.
	devs, err := device.RetrieveDevices(client, userToken, rootURL, projID)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the template with the device list as the data pipeline.
	err = templates.ExecuteTemplate(w, "index.html", &devs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/**
 * deviceInfoHandler is the handler for the route page "/Devices/{ID}" which presents a more in depth
 * overview of a specific device.
 */
func deviceInfoHandler(w http.ResponseWriter, r *http.Request, id string) {

	// Get a specific device by it's UUID
	dev, err := device.Retrieve(client, userToken, rootURL, projID, id)
	if err != nil {
		log.Fatalln(err)
	}

	// Use the device if it was retrieved as the pipeline to the route.
	err = templates.ExecuteTemplate(w, "deviceInfo.html", dev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/**
 * changeState is the handler for the route page "/Device/{ID}/action"
 * The action can either be "power_on" or "power_off".
 * After the action is performed, it is then redirected to the "/Devices/" page.
 */
func changeState(w http.ResponseWriter, r *http.Request, id, action string) {
	device.PerformAction(client, userToken, rootURL, id, action)
	http.Redirect(w, r, "/Devices/", http.StatusSeeOther)
}

/**
 * makeHandler is a closure for the handlers indexHandler and deviceInfoHandler.
 * It checks if the route path is a valid path and if so, creates the specified handler
 * with the id.
 */
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

/**
 * actionHandler is a closure for the handler changeState.
 * It checks if the route path is a valid path and if so, creates the handler with the id and action.
 */
func actionHandler(fn func(http.ResponseWriter, *http.Request, string, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validAction.FindStringSubmatch(r.URL.Path)
		if m == nil {

			io.WriteString(w, "Can't find: "+r.URL.Path)
			return
		}

		fn(w, r, m[2], m[3])
	}
}

func main() {

	// Saying hello!
	fmt.Println(color.YellowString("Hello Packet!!!"))

	// Create channels for graceful shutdown
	sigs := make(chan os.Signal, 1)
	srv := make(chan *http.Server, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	/**
	* Create a new router and a new HTTP Server.
	 */
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router}
	srv <- server

	// Graceful shutdown of the server.
	go gracefulShutDown(sigs, srv, done)

	// Paths
	router.HandleFunc("/Devices/", makeHandler(indexHandler))
	router.HandleFunc("/Devices/{ID}", makeHandler(deviceInfoHandler))
	router.HandleFunc("/Device/{ID}/power_on", actionHandler(changeState))
	router.HandleFunc("/Device/{ID}/power_off", actionHandler(changeState))

	// Stylesheets
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

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
