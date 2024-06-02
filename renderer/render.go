package renderer

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache map[string]*template.Template

func SetTemplateCache(cache map[string]*template.Template) {
	templateCache = cache
}

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	myTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(errors.New("template is missing"))
	}

	buffer := new(bytes.Buffer)
	_ = myTemplate.Execute(buffer, nil)

	_, err := buffer.WriteTo(w)

	if err != nil {
		fmt.Println("An error occured while writing template to response. The error: ", err)
	}
}
