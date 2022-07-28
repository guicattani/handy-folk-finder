package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSFR protection to all POST requests
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   ac.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
