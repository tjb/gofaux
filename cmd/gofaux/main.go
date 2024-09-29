package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gofaux/api/server"
	"gofaux/internal/writer"
	"io"
	"log"
	"net/http"
)

func main() {
	// Parse command line arguments
	url := flag.String("url", "", "URL to fetch JSON from")
	name := flag.String("name", "", "Name to save the file as")
	serverPtr := flag.Bool("server", false, "Run as a server")
	flag.Parse()

	// Fetch URL and save to file
	if *url != "" {
		// Fetch the URL
		resp, err := http.Get(*url)
		if err != nil {
			log.Fatalf("Failed to fetch URL: %v", err)
		}
		defer resp.Body.Close()

		// Handle non-200 responses
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Error: Received status code %d", resp.StatusCode)
		}

		// Read the response body
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}

		// Check if response body is empty
		if len(bodyBytes) == 0 {
			log.Fatal("Received an empty response body")
		}

		// Unmarshal the response into a generic map
		var jsonData map[string]interface{}
		err = json.Unmarshal(bodyBytes, &jsonData)
		if err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
		}

		// Print or save the JSON data
		fmt.Println("Successfully fetched and parsed JSON")

		// Save to file
		writer.WriteToFile(*name, jsonData)
	}

	// Start server
	if *serverPtr {
		server.CreateAndStartServer()
	}
}
