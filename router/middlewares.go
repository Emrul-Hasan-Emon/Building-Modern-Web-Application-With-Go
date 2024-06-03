package router

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

type MiddleWares struct {
	session *scs.SessionManager
}

func CreateNewMiddlewareInstance(session *scs.SessionManager) *MiddleWares {
	return &MiddleWares{session: session}
}

func (m *MiddleWares) noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
