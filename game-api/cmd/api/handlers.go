package main

import (
	"fmt"
	"net/http"
	"game-api/model"
	"log"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func (app *AppConfig) StartGame(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse {
		Error:   false,
		Message: fmt.Sprintf(""),
		Data:    "",
	}

	game := model.CreateGame(3)
	log.Println(game)
	app.writeJson(w, http.StatusAccepted, payload)
}
