// package ch6methods
package main

import (
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	// Create a new request using http.NewRequest and use the provided fullURL.
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	// Set the X-API-Key header, with apiKey as its value
	req.Header.Set("X-API-Key", apiKey)

	// Make the request using the http.Client's Do method. You can create a new *http.Client or use the http.DefaultClient
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the response status code. If the status code indicates a non-successful response, return an error. Otherwise return nil.
	if res.StatusCode > 299 {
		return fmt.Errorf("request to delete location unsuccessful")
	}
	return nil

}
