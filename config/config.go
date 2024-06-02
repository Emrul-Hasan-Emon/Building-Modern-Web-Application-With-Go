package config

import (
	"html/template"
	"log"
	"path/filepath"
)

// AppConfig holds the application configuration which will be accessible from anywhere of the application
type AppConfig struct {
	templateCache map[string]*template.Template
}

func (ac *AppConfig) GetTemplateCache() map[string]*template.Template {
	return ac.templateCache
}

var functions = template.FuncMap{}

func CreateNewConfigInstance() *AppConfig {
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	return &AppConfig{
		templateCache: templateCache,
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
