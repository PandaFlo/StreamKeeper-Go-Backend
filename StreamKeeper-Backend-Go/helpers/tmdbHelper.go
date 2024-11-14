package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Load environment variables (assuming you have a .env file or set environment variables)
var tmdbApiKey = os.Getenv("TMDB_API_KEY")
var tmdbBaseUrl = func() string {
	baseUrl := os.Getenv("TMDB_BASE_URL")
	if len(baseUrl) > 0 && baseUrl[len(baseUrl)-1] == '/' {
		return baseUrl[:len(baseUrl)-1]
	}
	return baseUrl
}()

// FetchFromTmdb is a helper function to make TMDb API calls
func FetchFromTmdb(endpoint string, params map[string]string) ([]byte, error) {
	// Construct the URL
	fullUrl := fmt.Sprintf("%s%s", tmdbBaseUrl, endpoint)
	requestURL, err := url.Parse(fullUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TMDB URL: %v", err)
	}

	// Add API key and additional parameters
	query := requestURL.Query()
	query.Add("api_key", tmdbApiKey)
	for key, value := range params {
		query.Add(key, value)
	}
	requestURL.RawQuery = query.Encode()

	// Log the request details
	log.Printf("\n[Sending Request to TMDb]\n- URL: %s\n- Params: %s\n", requestURL.String(), query.Encode())

	// Make the request
	resp, err := http.Get(requestURL.String())
	if err != nil {
		return nil, fmt.Errorf("request to TMDB failed: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read TMDB response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB request failed with status code %d: %s", resp.StatusCode, body)
	}

	return body, nil
}
