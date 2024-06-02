package renderer

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	myTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(errors.New("template is missing"))
	}

	buffer := new(bytes.Buffer)
	_ = myTemplate.Execute(buffer, nil)

	_, err = buffer.WriteTo(w)

	if err != nil {
		fmt.Println("An error occured while writing template to response. The error: ", err)
	}
}

// This function creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	layouts, err := filepath.Glob("templates/*.layout.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		if len(layouts) != 0 {
			templateSet, err = templateSet.ParseGlob("templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet
	}
	return myCache, nil
}
