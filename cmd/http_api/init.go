package main

import (
	"golang_start/entities"
	authmanager "golang_start/internal/auth_manager"
	"golang_start/pkg/app_config"
	kvstore "golang_start/pkg/kv_store"
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
	initHandler(app, r)
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

func initManagers(app *App) {
	app.auth_manager = authmanager.New(kvstore.New(), app.db)
}
