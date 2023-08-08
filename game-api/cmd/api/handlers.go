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
	Data    CardsResponse
}

type CardsResponse struct {
	CurrentDeck  []string
	PlayersCards map[string][]string
	Preflop      []string
}

func (app *AppConfig) NewGame(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse {
		Error:   false,
		Message: fmt.Sprintf("Success game start"),
		Data:    CardsResponse{},
	}
	log.Println(r.PostForm)
	log.Println(fmt.Sprintf("Number of players recieved: %s", r.PostFormValue("players")))
	

	game := model.CreateGame(int(4))
	payload.Data.CurrentDeck  = game.Cards.CurrentDeck
	payload.Data.PlayersCards = game.Cards.PlayersCards
	payload.Data.Preflop      = game.Cards.Preflop
	log.Println("game created")
	app.writeJson(w, http.StatusAccepted, payload)
}
