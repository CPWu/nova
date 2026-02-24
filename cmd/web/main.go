package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cpwu/nova/pkg/config"
	"github.com/cpwu/nova/pkg/handlers"
	"github.com/cpwu/nova/pkg/render"
)

// Define a constant portNumber to specify the port on which the HTTP server will listen for incoming requests.
const portNumber = ":8080"

// Create an instance of the AppConfig struct to hold application configuration settings and pass it to the handlers and render packages.
var app config.AppConfig

// Declare a global variable session of type *scs.SessionManager to manage user sessions across the application.
var session *scs.SessionManager

// main function that sets up the HTTP server and defines the routes for the Home and About pages. It also starts the server on port 8080 and prints a message to the console indicating that the server is running.
func main() {

	// change this to true when in production
	app.InProduction = false

	// Set up session management using the scs package and configure the session cookie settings.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // Set to true in production with HTTPS
	app.Session = session

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
