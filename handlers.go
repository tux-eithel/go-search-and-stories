package main

import (
	"log"
	"net/http"
)

func indexHandler(rt *routeTemplate, w http.ResponseWriter, r *http.Request) {

	preparedURL := filterSources(mainFeedURL, preferSources.Sources)

	var list []*news
	var err error
	if preparedURL != nil {
		list, err = getNews(preparedURL.String())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if rt.tmpl != nil {
		err = rt.tmpl.Execute(w, map[string]interface{}{
			"news":  list,
			"error": err,
		})
		if err != nil {
			log.Println(err)
		}
	}

}

func indexSettings(rt *routeTemplate, w http.ResponseWriter, r *http.Request) {

	var err error
	var list []*feed
	var listByCat map[string][]*feed

	if r.Method == "POST" {
		err = r.ParseForm()
		if err == nil {
			preferSources.Sources = r.Form["sources"]
			saveSettings(settingsFileName, preferSources)
		}
	}

	if err == nil {
		list, err = getSources(sourcesURL)
	}

	if err == nil {
		listByCat = feedByCategory(list)
		listByCat = orderFeedByName(listByCat)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if rt.tmpl != nil {
		err = rt.tmpl.Execute(w, map[string]interface{}{
			"sources":   listByCat,
			"error":     err,
			"mysources": preferSources.Sources,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
