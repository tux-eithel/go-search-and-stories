package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	flagPort             int
)

func init() {
	preferSources = loadSettings(settingsFileName)

	flag.BoolVar(&flagUseTemplateFiles, "t", false, "use html files instead of compiled templates (default false)")
	flag.IntVar(&flagPort, "p", 8080, "listen server port")
}

func main() {

	flag.Parse()

	if flagPort < 1 {
		fmt.Println("port number must be > 0")
		os.Exit(1)
	}

	flagPortStr := strconv.Itoa(flagPort)

	templateFolder := "templates" + string(os.PathSeparator)
	routeTemplates := fillTemplate(templateFolder, flagUseTemplateFiles)

	for _, tmpl := range routeTemplates {
		http.Handle("/"+tmpl.url, tmpl)
	}

	fmt.Println("open with browser: http://localhost:" + flagPortStr)

	http.ListenAndServe(":"+flagPortStr, nil)

}
