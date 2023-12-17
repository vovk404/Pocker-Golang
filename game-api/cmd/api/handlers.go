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

type GetCurrentGameRequest struct {
	UserId string `json:"userId"`
}

type CardsResponse struct {
	CurrentDeck  []string
	PlayersCards map[string][]string
	Flop      []string
}

func (app *AppConfig) NewGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var playersRequest NewGameJsonRequest
    if err := decoder.Decode(&playersRequest);  err != nil {
		log.Println("Error occured while decoding the data in authentication service: ")
        return
	}

	game := model.CreateGame(playersRequest.Players)

	payload := JsonResponse {
		Error:   false,
		Message: fmt.Sprintf("Success game start"),
		Data:    CardsResponse{},
	}
	payload.Data.PlayersCards = game.Cards.PlayersCards
	log.Println("game created")
	app.writeJson(w, http.StatusOK, payload)
}

func (app *AppConfig) OpenPreFlop(w http.ResponseWriter, r *http.Request) {
	
	payload := JsonResponse {
		Error:   false,
		Message: fmt.Sprintf("Preflop successfuly opened"),
		Data:    CardsResponse{},
	}
	// TODO 
	game, err := model.GetCurrentGame(w, r)
	if err != nil {
		log.Panicln("Can`t get current game")
	}
	game.Cards.OpenPreflop()

	payload.Data.Flop = game.Cards.Flop
	app.writeJson(w, http.StatusOK, payload)
}

func (app *AppConfig) GetCurrentUserGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var currentGameRequest GetCurrentGameRequest
    if err := decoder.Decode(&currentGameRequest);  err != nil {
		log.Println("Error occured while decoding the data in request: ")
        return
	}

	// TODO
}