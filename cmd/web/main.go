package main

import (
	"fmt"
	"github.com/stephanusnugraha/go-web-app/pkg/config"
	"github.com/stephanusnugraha/go-web-app/pkg/handlers"
	"github.com/stephanusnugraha/go-web-app/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8084"

// main app
func main() {
	var app config.AppConfig

	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = cache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Println(err)
	}
}
