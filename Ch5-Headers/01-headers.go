// package ch5headers
package main

import (
	"net/http"
)

func getContentType(res *http.Response) string {
	// ?
	header := res.Header.Get("Content-Type")
	//fmt.Println(header)
	return header
}
