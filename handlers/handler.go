package handlers

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/model"
	"github.com/Emrul-Hasan-Emon/application/renderer"
)

type Repository struct {
	rndr *renderer.RenderTemplate
}

func CreateNewRepository(rndr *renderer.RenderTemplate) *Repository {
	return &Repository{rndr}
}

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Data is coming from backend"
	rp.rndr.RenderTemplates(w, "home.page.html", model.TemplateData{
		StringMap: stringMap,
	})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	rp.rndr.RenderTemplates(w, "about.page.html", model.TemplateData{})
}
