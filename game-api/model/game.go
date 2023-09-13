package model

import (
	"log"
	"encoding/json"
	"context"
	"net/http"
	"github.com/go-session/session/v3"
)

type Game struct {
	Cards Cards
}

func CreateGame(palyers int) Game {
	
	game := Game {
		Cards: Cards{},
	}
	game.Cards.CreateNewDeck()
	game.Cards.PassCardsToPlayers(palyers)

	log.Println("cards set up finished")
	return game
}

func GetCurrentGame(w http.ResponseWriter, r *http.Request) (Game, error) {

	game := Game {
		Cards: Cards{},
	}


	return game, nil
}

func (game *Game) SaveGameInotRedisSession(w http.ResponseWriter, r *http.Request) string {
	sessionStore, err := session.Start(context.Background(), w, r)
	sessionId := sessionStore.SessionID()
	log.Println("My session id in save: ", sessionId)

	if err != nil {
		log.Println("Error during session start")
		return sessionId
	}

	json, err := json.Marshal(game)
    if err != nil {
		log.Println("Error during marshaling game")
        return sessionId
    }

	sessionStore.Set("CurrentGame", json)
	err = sessionStore.Save()
	if err != nil {
		log.Println("Error during saving game into the session")
		return sessionId
	}

	return sessionId
}