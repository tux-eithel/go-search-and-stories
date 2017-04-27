package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type (
	routeTemplate struct {
		routeName string
		url       string
		fileNames []string
		tmpl      *template.Template
		fnc       func(*routeTemplate, http.ResponseWriter, *http.Request)
	}
)

func (rt *routeTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.fnc(rt, w, r)
}

func fillTemplate() []*routeTemplate {

	rTemplates := []*routeTemplate{
		&routeTemplate{
			routeName: "index",
			url:       "",
			fileNames: []string{
				tplBase,
				tplIndex,
			},
			fnc: indexHandler,
		},

		&routeTemplate{
			routeName: "settings",
			url:       "settings",
			fileNames: []string{
				tplBase,
				tplSettings,
			},
			fnc: indexSettings,
		},
	}

	for index, tmpl := range rTemplates {
		var t *template.Template
		var err error
		for _, str := range tmpl.fileNames {
			if t == nil {
				t = template.New(tmpl.routeName)
				t.Funcs(template.FuncMap{
					"inarray": inArray,
					"validid": validID,
				})
			}
			t, err = t.Parse(str)
			if err != nil {
				log.Fatalln(err)
			}

		}
		rTemplates[index].tmpl = t
	}

	return rTemplates
}

func inArray(input []string, str string) bool {
	for _, current := range input {
		if current == str {
			return true
		}
	}
	return false
}

func validID(str string) string {
	localStr := strings.ToLower(str)
	return strings.Replace(localStr, " ", "-", -1)
}
