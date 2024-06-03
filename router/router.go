package router

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	routes      http.Handler
	middlewares *MiddleWares
}

func CreateNewRouterInstance(middlewares *MiddleWares, repo *handlers.Repository) *Router {
	return &Router{
		routes:      setRoutes(repo, middlewares),
		middlewares: middlewares,
	}
}

func setRoutes(repo *handlers.Repository, middlewares *MiddleWares) http.Handler {
	mux := chi.NewRouter()

	setMiddlewares(mux, middlewares)
	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	return mux
}

func setMiddlewares(mux *chi.Mux, middlewares *MiddleWares) {
	mux.Use(middlewares.noSurf)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(middlewares.sessionLoad)
}

func (r *Router) GetRoutes() http.Handler {
	return r.routes
}
