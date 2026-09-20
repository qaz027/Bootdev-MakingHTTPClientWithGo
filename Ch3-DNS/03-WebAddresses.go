// package ch2
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://1.1.1.1/dns-query?name=%s&type=A", domain)
	//url := fmt.Sprintf("https://dns.google/resolve?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var DNSResp DNSResponse
	if err := json.Unmarshal(body, &DNSResp); err != nil {
		return "", err
	} else {
		if len(DNSResp.Answer) == 0 {
			return "", fmt.Errorf("no IP address found")
		} else {
			return DNSResp.Answer[0].Data, nil
		}
	}

	//return string(body), nil
}
