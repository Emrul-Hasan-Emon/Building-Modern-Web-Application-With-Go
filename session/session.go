package session

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type SessionManager struct {
	session *scs.SessionManager
}

func CreateNewSessionManager(session *scs.SessionManager) *SessionManager {
	return &SessionManager{session: session}
}

func (sm *SessionManager) GetSession() *scs.SessionManager {
	return sm.session
}

func (sm *SessionManager) SetSessionCookie(r *http.Request) {
	remoteIP := r.RemoteAddr
	sm.session.Put(r.Context(), "remote_ip", remoteIP)
}

func (sm *SessionManager) GetSessionCookie(r *http.Request) string {
	return sm.session.GetString(r.Context(), "remote_ip")
}
