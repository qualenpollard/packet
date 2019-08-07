package plan

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetPlans ...
func GetPlans(c *http.Client, token, url string) DataBase {
	// Create a new request to get the Plan data.
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
