package main

import (
	"log"
	"net/http"
	"os"
)

type App struct {
	srv    *http.Server
	logger *log.Logger
}

// profileManager, TodoManager, AuthManager.

var (
	l = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {

	app := &App{
		logger: l,
	}

	initServer(app)

	app.logger.Println("app is running on ", app.srv.Addr)
	app.logger.Fatalln(app.srv.ListenAndServe())

}
