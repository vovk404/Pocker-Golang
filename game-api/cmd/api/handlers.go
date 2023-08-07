package main

import (
	"fmt"
	"game-api/model"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    []string
}

func (app *AppConfig) NewGame(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse {
		Error:   false,
		Message: fmt.Sprintf("Success game start"),
		Data:    []string{},
	}

	game := model.CreateGame(3)
	payload.Data = game.CurrentDeck
	log.Println("game created")
	app.writeJson(w, http.StatusAccepted, payload)
}
