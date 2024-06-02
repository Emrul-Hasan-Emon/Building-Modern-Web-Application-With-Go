package router

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	Routes http.Handler
}

func SetRoutes(repo *handlers.Repository) *Router {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	return &Router{
		Routes: mux,
	}
}
