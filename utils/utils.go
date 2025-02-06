package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var assetsDir = "./assets/"

// Function to read websites from JSON
func getWebsites() ([]string, error) {
	var data struct {
		Websites []string `json:"websites"`
	}
	file, err := ioutil.ReadFile(assetsDir + "websites.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return data.Websites, nil
}

// Function to read User-Agent strings from the file
func getUserAgents() ([]string, error) {
	file, err := ioutil.ReadFile(assetsDir + "useragents.txt")
	if err != nil {
		return nil, err
	}

	// Split file by newline and return as a slice of strings
	userAgents := string(file)
	return splitLines(userAgents), nil
}

// Function to split lines from a string
func splitLines(text string) []string {
	var lines []string
	for _, line := range strings.Split(text, "\n") {
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	return lines
}

// Function to scan websites for the username
func ScanWebsites(username string) {
	websites, err := getWebsites()
	if err != nil {
		log.Fatalf("Error reading websites from JSON: %v", err)
	}

	userAgents, err := getUserAgents()
	if err != nil {
		log.Fatalf("Error reading user agents from file: %v", err)
	}

	// Set a random seed
	rand.Seed(time.Now().UnixNano())

	for _, site := range websites {
		fmt.Printf("Scanning site: %s for username: %s\n", site, username)
		// Pick a random User-Agent from the list
		randomUserAgent := userAgents[rand.Intn(len(userAgents))]
		fmt.Printf("Using User-Agent: %s\n", randomUserAgent)

		// Call the search site function with the random User-Agent
		searchSite(site, username, randomUserAgent)
	}
}

// Function to perform a simple search on the site with a random User-Agent
func searchSite(site, username, userAgent string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/search?q=%s", site, username), nil)
	if err != nil {
		fmt.Printf("Error creating request for site %s: %v\n", site, err)
		return
	}

	// Set random User-Agent header
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request to site %s: %v\n", site, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Found user '%s' on site: %s\n", username, site)
	} else {
		fmt.Printf("No results for user '%s' on site: %s\n", username, site)
	}
}
