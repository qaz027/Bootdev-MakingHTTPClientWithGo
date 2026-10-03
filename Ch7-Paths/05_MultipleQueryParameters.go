// package ch7paths
package main

func fetchTasks(baseURL, availability string) []Issue {
	// ?
	limit := "1"
	if availability == "Medium" {
		limit = "3"
	} else if availability == "High" {
		limit = "5"
	}

	fullURL := baseURL + "?sort=estimate&limit=" + limit

	return getIssues(fullURL)
}
