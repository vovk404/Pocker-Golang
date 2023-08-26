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
	game.Cards.createNewDeck()
	game.Cards.passCardsToPlayers(palyers)
	game.Cards.openPreflop()

	log.Println("cards set up finished")
	return game
}

// func GetCurrentGame(w http.ResponseWriter, r *http.Request) (Game, error) {
// 	sessionStore, err := session.Start(context.Background(), w, r)
// 	game := Game {
// 		Cards: Cards{},
// 	}
// 	if err != nil {
// 		return game, err
// 	}
// 	currentGame, ok := sessionStore.Get("CurrentGame")
// 	if ok {
// 		err = json.Unmarshal(currentGame, game)
// 		if err != nil {
// 		   return game, err
// 	    }
//     }

// 	return game, nil
// }

func (game *Game) CreateRedisSession(w http.ResponseWriter, r *http.Request) {
	sessionStore, err := session.Start(context.Background(), w, r)
	if err != nil {
		log.Println("Error during session start")
		return 
	}

	json, err := json.Marshal(game)
    if err != nil {
		log.Println("Error during marshaling game")
        return 
    }

	sessionStore.Set("CurrentGame", json)
	err = sessionStore.Save()
	if err != nil {
		log.Println("Error during saving game into the session")
		return 
	}
}