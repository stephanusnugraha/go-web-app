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

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
