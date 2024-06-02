package renderer

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type RenderTemplate struct {
	templateCache map[string]*template.Template
}

func CreateNewRenderTemplateInstance(cache map[string]*template.Template) *RenderTemplate {
	return &RenderTemplate{cache}
}

func (rndr *RenderTemplate) RenderTemplates(w http.ResponseWriter, tmpl string) {
	myTemplate, ok := rndr.templateCache[tmpl]
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
