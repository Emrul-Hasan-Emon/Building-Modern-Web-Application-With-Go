package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/config"
	"github.com/Emrul-Hasan-Emon/application/handlers"
	"github.com/Emrul-Hasan-Emon/application/renderer"
	"github.com/Emrul-Hasan-Emon/application/router"
	"github.com/Emrul-Hasan-Emon/application/session"
)

const portNumber = ":8080"

func main() {
	appConfig := config.CreateNewConfigInstance()
	rndr := renderer.CreateNewRenderTemplateInstance(appConfig.GetTemplateCache())
	repo := handlers.CreateNewRepository(rndr)

	sessionManager := session.CreateNewSessionManager(appConfig.GetSessionManager())
	middlwares := router.CreateNewMiddlewareInstance(sessionManager)
	router := router.CreateNewRouterInstance(middlwares, repo)

	server := http.Server{
		Addr:    portNumber,
		Handler: router.GetRoutes(),
	}

	fmt.Println("Started application on port: ", portNumber)

	err := server.ListenAndServe()
	log.Fatal(err)
}
