package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requests. This is a custom middleware function that wraps the nosurf library.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   app.InProduction, // Set to true in production with HTTPS
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request. This is a custom middleware function that wraps the session management logic.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
