package main

import (
	"golang_start/pkg/app_config"
	"net/http"

	"github.com/go-chi/chi"
)

func initServer(app *App) {
	r := chi.NewRouter()

	srv := http.Server{
		Addr:    app_config.Data.MustString("SERVER_ADDR"),
		Handler: r,
	}

	app.srv = &srv

}
