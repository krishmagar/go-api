package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/krishmagar/go-api/pkg/config"
	"github.com/krishmagar/go-api/pkg/handlers"
	"github.com/krishmagar/go-api/pkg/render"
)

const PORT string = ":8080"

func main() {
	var app config.AppConfig

	// Create template cache
	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tempCache // set TemplateCache in the AppConfig
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app) // Provide AppConfig access to the render package

	fmt.Println(fmt.Sprintf("Starting application on port%s", PORT))

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
