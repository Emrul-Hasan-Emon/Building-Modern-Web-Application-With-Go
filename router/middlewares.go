package router

import (
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/session"
	"github.com/justinas/nosurf"
)

type MiddleWares struct {
	sessionManager *session.SessionManager
}

func CreateNewMiddlewareInstance(sessionManager *session.SessionManager) *MiddleWares {
	return &MiddleWares{sessionManager: sessionManager}
}

// no surf adds CSRF protection to all POST requests.
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

// sessionLoad middleware handles session
func (m *MiddleWares) sessionLoad(next http.Handler) http.Handler {
	return m.sessionManager.GetSession().LoadAndSave(next)
}
