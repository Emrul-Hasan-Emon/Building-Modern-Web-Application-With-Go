package config

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application configuration which will be accessible from anywhere of the application
type AppConfig struct {
	templateCache map[string]*template.Template
	session       *scs.SessionManager
}

func (ac *AppConfig) GetTemplateCache() map[string]*template.Template {
	return ac.templateCache
}

func (ac *AppConfig) GetSessionManager() *scs.SessionManager {
	return ac.session
}

var functions = template.FuncMap{}

func CreateNewConfigInstance() *AppConfig {
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	return &AppConfig{
		templateCache: templateCache,
		session:       initiateSession(),
	}
}

// This function creates a template cache as a map
func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	layouts, err := filepath.Glob("templates/*.layout.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		if len(layouts) != 0 {
			templateSet, err = templateSet.ParseGlob("templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet
	}
	return myCache, nil
}

func initiateSession() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	return session
}
