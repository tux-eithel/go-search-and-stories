package main

import (
	"fmt"
	"net/http"
)

const (
	mainFeedURL = "https://watrcoolr.duckduckgo.com/watrcoolr.js?o=json"
	sourcesURL  = "https://watrcoolr.duckduckgo.com/watrcoolr.js?o=json&type_info=1"
)

var (
	preferSources []string
)

func main() {

	routeTemplates := fillTemplate()

	for _, tmpl := range routeTemplates {
		http.Handle("/"+tmpl.url, tmpl)
	}

	fmt.Println("open with browser: http://localhost:8080")

	http.ListenAndServe(":8080", nil)

}
