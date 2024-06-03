package session

import "github.com/alexedwards/scs/v2"

type SessionManager struct {
	session *scs.SessionManager
}

func CreateNewSessionManager(session *scs.SessionManager) *SessionManager {
	return &SessionManager{session: session}
}

func (sm *SessionManager) GetSession() *scs.SessionManager {
	return sm.session
}
