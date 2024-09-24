package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func getReviewItem(id string) (*ReviewItResponse, error) {
	api := os.Getenv("APP_API")

	if api == "" {
		fmt.Println("No APP Url found")
		return nil, fmt.Errorf("no APP_URL found")
	}

	// Create the request body
	requestBody := ReviewRequest{ID: id}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Make the POST request
	resp, err := http.Post(api, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %s", resp.Status)
	}

	// Decode the response body
	var reviewData ReviewItResponse
	if err := json.NewDecoder(resp.Body).Decode(&reviewData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return &reviewData, nil
}
