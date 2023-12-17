package main

import (
	"net/http"
	"front-end/cmd/web/controller"
)

func (app *AppConfig) routes() {
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/", controller.OpenMainPage)
	http.HandleFunc("/new_game", controller.OpenGamePage)
	return 
}