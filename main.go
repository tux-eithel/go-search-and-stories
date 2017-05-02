package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	mainFeedURL = "https://watrcoolr.duckduckgo.com/watrcoolr.js?o=json"
	sourcesURL  = "https://watrcoolr.duckduckgo.com/watrcoolr.js?o=json&type_info=1"

	settingsFileName = ".ddg_settings"
)

type (
	preference struct {
		Sources []string `json:"sources"`
	}
)

var (
	preferSources = &preference{}

	flagUseTemplateFiles bool
)

func init() {
	preferSources = loadSettings(settingsFileName)

	flag.BoolVar(&flagUseTemplateFiles, "t", false, "use html filse instead of compiled templates")
}

func main() {

	flag.Parse()

	templateFolder := "templates" + string(os.PathSeparator)
	routeTemplates := fillTemplate(templateFolder, flagUseTemplateFiles)

	for _, tmpl := range routeTemplates {
		http.Handle("/"+tmpl.url, tmpl)
	}

	fmt.Println("open with browser: http://localhost:8080")

	http.ListenAndServe(":8080", nil)

}
