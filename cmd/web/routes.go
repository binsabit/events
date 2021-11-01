package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() *pat.PatternServeMux {

	mux := pat.New()
	mux.Get("/events", http.HandlerFunc(app.showEvents))
	mux.Get("/event/:id", http.HandlerFunc(app.showSpecificEvent))
	mux.Post("/create", http.HandlerFunc(app.createEvent))
	mux.Post("/subscribe", http.HandlerFunc(app.subscribeToEvent))

	mux.Post("/login", http.HandlerFunc(app.userLogin))
	mux.Post("/register", http.HandlerFunc(app.userRegister))

	return mux
}
