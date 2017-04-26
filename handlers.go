package main

import "net/http"

func indexHandler(rt *routeTemplate, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if rt.tmpl != nil {
		rt.tmpl.Execute(w, nil)
	}

}

func indexSettings(rt *routeTemplate, w http.ResponseWriter, r *http.Request) {

	list, err := getSources(sourcesURL)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if rt.tmpl != nil {
		rt.tmpl.Execute(w, map[string]interface{}{
			"sources": list,
			"error":   err,
		})
	}
}
