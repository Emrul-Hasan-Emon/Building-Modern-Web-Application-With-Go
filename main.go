package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/config"
	"github.com/Emrul-Hasan-Emon/application/handlers"
	"github.com/Emrul-Hasan-Emon/application/renderer"
	"github.com/Emrul-Hasan-Emon/application/router"
)

const portNumber = ":8080"

func main() {
	appConfig := config.CreateNewConfigInstance()
	rndr := renderer.CreateNewRenderTemplateInstance(appConfig.GetTemplateCache())
	repo := handlers.CreateNewRepository(rndr)
	router := router.SetRoutes(repo)

	server := http.Server{
		Addr:    portNumber,
		Handler: router.Routes,
	}

	fmt.Println("Started application on port: ", portNumber)

	err := server.ListenAndServe()
	log.Fatal(err)
}
