package handlers

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/renderer"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplates(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplates(w, "about.html")
}
