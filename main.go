package main

import (
	"fmt"
	"net/http"

	"github.com/Emrul-Hasan-Emon/application/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Started application on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
