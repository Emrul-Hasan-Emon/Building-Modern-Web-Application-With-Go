package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.html")
}

func renderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
