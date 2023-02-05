package main

import (
	authmanager "golang_start/internal/auth_manager"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type App struct {
	srv          *http.Server
	logger       *log.Logger
	db           *gorm.DB
	auth_manager *authmanager.AuthManager
}

// profileManager, TodoManager, AuthManager.

var (
	l = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {

	app := &App{
		logger: l,
	}

	initDB(app)
	initServer(app)
	initManagers(app)
	app.logger.Println("app is running on ", app.srv.Addr)
	app.logger.Fatalln(app.srv.ListenAndServe())

}
