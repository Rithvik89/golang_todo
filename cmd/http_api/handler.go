package main

import (
	"github.com/go-chi/chi"
)

func initHandler(app *App, r *chi.Mux) {
	r.Post("/signup", app.HandleRegisterUser)
}
