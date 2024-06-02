package router

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/handlers"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Routes http.Handler
}

func SetRoutes(repo *handlers.Repository) *Router {
	mux := chi.NewRouter()
	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	return &Router{
		Routes: mux,
	}
}
