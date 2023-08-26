package main

import (
	"fmt"
	"game-api/model"
	"log"
	"net/http"
	"encoding/json"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    CardsResponse
}

type NewGameJsonRequest struct {
	Players int `json:"players"`
}

type CardsResponse struct {
	CurrentDeck  []string
	PlayersCards map[string][]string
	Flop      []string
}

func (app *AppConfig) NewGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var playersRequest NewGameJsonRequest
    error := decoder.Decode(&playersRequest)
    if error != nil {
        fmt.Println("Error occured while decoding the data: ", error)
        return
    }

	game := model.CreateGame(playersRequest.Players)
	game.CreateRedisSession(w, r)
	payload := JsonResponse {
		Error:   false,
		Message: fmt.Sprintf("Success game start"),
		Data:    CardsResponse{},
	}
	payload.Data.CurrentDeck  = game.Cards.CurrentDeck
	payload.Data.PlayersCards = game.Cards.PlayersCards
	payload.Data.Flop         = game.Cards.Preflop
	log.Println("game created")
	app.writeJson(w, http.StatusOK, payload)
}

// func (app *AppConfig) OpenPreFlop(w http.ResponseWriter, r *http.Request) {
// 	game := model
// 	game.CreateRedisSession(w, r)
// 	payload := JsonResponse {
// 		Error:   false,
// 		Message: fmt.Sprintf("Preflop successfuly opened"),
// 		Data:    CardsResponse{},
// 	}
// 	payload.Data.CurrentDeck  = game.Cards.CurrentDeck
// 	payload.Data.PlayersCards = game.Cards.PlayersCards
// 	payload.Data.Flop         = game.Cards.Preflop
// 	log.Println("game created")
// 	app.writeJson(w, http.StatusOK, payload)
// }