package main

import "net/http"

func indexHandler(rt *routeTemplate, w http.ResponseWriter, r *http.Request) {

	preparedURL := filterSources(mainFeedURL, preferSources)

	var list []*news
	var err error
	if preparedURL != nil {
		list, err = getNews(preparedURL.String())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if rt.tmpl != nil {
		rt.tmpl.Execute(w, map[string]interface{}{
			"news":  list,
			"error": err,
		})
	}

}

func indexSettings(rt *routeTemplate, w http.ResponseWriter, r *http.Request) {

	var err error
	var list []*feed
	var listByCat map[string][]*feed

	if r.Method == "POST" {
		err = r.ParseForm()
		if err == nil {
			preferSources = r.Form["sources"]
		}
	}

	if err == nil {
		list, err = getSources(sourcesURL)
	}

	if err == nil {
		listByCat = feedByCategory(list)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if rt.tmpl != nil {
		rt.tmpl.Execute(w, map[string]interface{}{
			"sources":   listByCat,
			"error":     err,
			"mysources": preferSources,
		})
	}
}
