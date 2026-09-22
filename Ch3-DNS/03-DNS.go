// package ch3dns
package main

import (
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
	// ?
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		//fmt.Println("error parsing url:", err)
		return "", err
	}
	return parsedURL.Hostname(), nil
}
