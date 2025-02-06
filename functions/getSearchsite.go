package functions

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

type Website struct {
	ErrorMsg        interface{} `json:"errorMsg"`
	ErrorType       string      `json:"errorType"`
	RegexCheck      string      `json:"regexCheck,omitempty"`
	URL             string      `json:"url"`
	URLMain         string      `json:"urlMain"`
	UsernameClaimed string      `json:"username_claimed"`
}

var assetsDir = "./assets/"

// GetWebsites loads and parses the websites configuration from a JSON file
func GetWebsites() (map[string]Website, error) {
	file, err := os.ReadFile(assetsDir + "websites.json")
	if err != nil {
		return nil, err
	}

	var websites map[string]Website
	err = json.Unmarshal(file, &websites)
	if err != nil {
		return nil, err
	}
	return websites, nil
}

// SendRequest sends an HTTP GET request with the provided username and user agent
func SendRequest(url, username, userAgent string) (*http.Response, error) {
	client := &http.Client{}
	// Correctly format the URL by manually replacing `{}` with username
	url = strings.Replace(url, "{}", username, -1)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CheckIfUsernameExists checks if a given username exists on a specific website
func CheckIfUsernameExists(website Website, username string) (bool, error) {
	userAgent, err := GetRandomUserAgent()
	if err != nil {
		return false, err
	}

	resp, err := SendRequest(website.URL, username, userAgent)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	switch website.ErrorType {
	case "status_code":
		if resp.StatusCode == 404 {
			return false, nil
		}
		return true, nil

	case "message":
		errorMsg := website.ErrorMsg.(string)
		if strings.Contains(string(body), errorMsg) {
			return false, nil
		}
		return true, nil

	default:
		return false, nil
	}
}

// SearchUserAcrossWebsites searches for a username across multiple websites and reports if it's found
func SearchUserAcrossWebsites(username string) error {
	websites, err := GetWebsites()
	if err != nil {
		return err
	}

	for _, website := range websites {
		// Correctly format the URL by replacing {} with the username
		url := strings.Replace(website.URL, "{}", username, -1)

		resp, err := SendRequest(url, username, "") // Passing empty string for the user agent (it'll be handled in the function)
		if err != nil {
			return fmt.Errorf("error sending request to %s: %v", website.URLMain, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			if website.ErrorType == "status_code" && resp.StatusCode == 404 {
				fmt.Printf("Profile for username '%s' not found on %s\n", username, website.URLMain)
			} else {
				fmt.Printf("Error accessing %s: %s\n", website.URLMain, resp.Status)
			}
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return fmt.Errorf("error parsing response for %s: %v", website.URLMain, err)
		}

		if strings.Contains(doc.Text(), website.UsernameClaimed) {
			fmt.Printf("Profile for username '%s' found on %s\n", username, website.URLMain)
		} else {
			fmt.Printf("Profile for username '%s' not found on %s\n", username, website.URLMain)
		}
	}
	return nil
}
