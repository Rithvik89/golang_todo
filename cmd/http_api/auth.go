package main

import "net/http"

func (app *App) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	// sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
}

func (app *App) HandleRegisterUser(rw http.ResponseWriter, r *http.Request) {
	mailId := "rithvikalkanti@gmail.com"
	pass := "mypassword"
	session, err := app.auth_manager.RegisterUser(mailId, pass)
	if err != nil {
		sendErrorResponse(rw, 500, nil, "Internal Server Error")
	}
	sendResponse(rw, 201, session, "User Created")

}

func (app *App) HandleLogout() {

}
