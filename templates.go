package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type (
	routeTemplate struct {
		routeName     string
		url           string
		fileNames     []string
		fileTemplates []string
		tmpl          *template.Template
		fnc           func(*routeTemplate, http.ResponseWriter, *http.Request)
	}
)

func (rt *routeTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.fnc(rt, w, r)
}

func fillTemplate(baseString string, useTemplateFiles bool) []*routeTemplate {

	rTemplates := []*routeTemplate{
		{
			routeName: "index",
			url:       "",
			fileNames: []string{
				tplBase,
				tplIndex,
			},
			fileTemplates: []string{
				baseString + "base.html",
				baseString + "index.html",
			},
			fnc: indexHandler,
		},

		{
			routeName: "settings",
			url:       "settings",
			fileNames: []string{
				tplBase,
				tplSettings,
			},
			fileTemplates: []string{
				baseString + "base.html",
				baseString + "settings.html",
			},
			fnc: indexSettings,
		},
	}

	for index, tmpl := range rTemplates {
		var t *template.Template
		var err error

		if (useTemplateFiles && len(tmpl.fileTemplates) > 0) || (!useTemplateFiles && len(tmpl.fileNames) > 0) {
			t = template.New("base.html")
			t.Funcs(template.FuncMap{
				"inarray": inArray,
				"validid": validID,
			})
		}

		if useTemplateFiles && len(tmpl.fileTemplates) > 0 {
			_, err = t.ParseFiles(tmpl.fileTemplates...)
			if err != nil {
				log.Fatalln(err)
			}
		} else if !useTemplateFiles && len(tmpl.fileNames) > 0 {
			for _, str := range tmpl.fileNames {
				_, err = t.Parse(str)
				if err != nil {
					log.Fatalln(err)
				}
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
