package main

import (
	"fmt"
	"github.com/stephanusnugraha/go-web-app/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8084"

// main app
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Println(err)
	}
}
