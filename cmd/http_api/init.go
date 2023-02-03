package main

import (
	"golang_start/entities"
	"golang_start/pkg/app_config"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initServer(app *App) {
	r := chi.NewRouter()

	srv := http.Server{
		Addr:    app_config.Data.MustString("SERVER_ADDR"),
		Handler: r,
	}

	app.srv = &srv

}

func initDB(app *App) {
	dsn := app_config.Data.MustString("DB_ADDR")
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app.db = conn

	// schema migration ..
	(app.db).AutoMigrate(&entities.Users{})

}
