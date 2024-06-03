package handlers

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/model"
	"github.com/Emrul-Hasan-Emon/application/renderer"
	"github.com/Emrul-Hasan-Emon/application/session"
)

type Repository struct {
	rndr           *renderer.RenderTemplate
	sessionManager *session.SessionManager
}

func CreateNewRepository(rndr *renderer.RenderTemplate, sessionManager *session.SessionManager) *Repository {
	return &Repository{rndr, sessionManager}
}

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	rp.sessionManager.SetSessionCookie(r)
	rp.rndr.RenderTemplates(w, "home.page.html", model.TemplateData{})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = rp.sessionManager.GetSessionCookie(r)
	rp.rndr.RenderTemplates(w, "about.page.html", model.TemplateData{
		StringMap: stringMap,
	})
}
