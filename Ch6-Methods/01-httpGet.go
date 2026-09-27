// package ch6methods
package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	// ?
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var user []User
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&user); err != nil {
		return nil, err

	}
	return user, nil
}
