package router

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/handlers"
	"github.com/gorilla/pat"
)

type Router struct {
	Routes http.Handler
}

func SetRoutes(repo *handlers.Repository) *Router {
	mux := pat.New()
	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	return &Router{
		Routes: mux,
	}
}
