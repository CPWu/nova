package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cpwu/nova/pkg/config"
	"github.com/cpwu/nova/pkg/handlers"
	"github.com/cpwu/nova/pkg/render"
)

const portNumber = ":8080"

// main function that sets up the HTTP server and defines the routes for the Home and About pages. It also starts the server on port 8080 and prints a message to the console indicating that the server is running.
func main() {
	var app config.AppConfig
	// Create a template cache and store it in the app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache:", err)
		return
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemmplates(&app)

	fmt.Printf("Starting server on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
