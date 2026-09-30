// package ch6methods
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	// Encode the user data as JSON using json.Marshal
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}
	// Create a new request using http.NewRequest
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData)) //Set the body as a bytes.Buffer containing the encoded JSON data using bytes.NewBuffer
	if err != nil {
		return User{}, err
	}

	// Modify the request headers
	// Set the Content-Type header, with application/json as its value
	req.Header.Set("Content-Type", "application/json")
	// Set the X-API-Key header, with apiKey as its value
	req.Header.Set("X-API-Key", apiKey)

	// Make the request using the http.Client's Do method. You can create a new *http.Client or use the http.DefaultClient
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	// Decode and return the response's JSON body (which is also a User)
	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil

}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	// Create a new request using http.NewRequest
	// Set the method as GET
	// Set the url to fullURL
	req, err := http.NewRequest("GET", fullURL, nil) // Set the body as nil
	if err != nil {
		return User{}, err
	}

	// Set the X-API-Key header, with apiKey as its value
	req.Header.Set("X-API-Key", apiKey)

	// Make the request using the http.Client's Do method. You can create a new *http.Client or use the http.DefaultClient
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	// Decode and return the response's JSON body (which is also a User)
	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
