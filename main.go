package main

import (
	"fmt"
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/config"
	"github.com/Emrul-Hasan-Emon/application/handlers"
	"github.com/Emrul-Hasan-Emon/application/renderer"
)

const portNumber = ":8080"

func main() {
	appConfig := config.CreateNewConfigInstance()
	renderer.SetTemplateCache(appConfig.GetTemplateCache())

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Started application on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
