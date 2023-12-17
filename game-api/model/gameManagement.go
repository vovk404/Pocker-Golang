package model

import (
	"log"
	"net/http"
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
	err := Insert(&game, 3)
	if err != nil {
		log.Println("Could not save a new game ", err)
	}

	log.Println("cards set up finished")
	return game
}

func GetCurrentGame(w http.ResponseWriter, r *http.Request) (Game, error) {

	game := Game {
		Cards: Cards{},
	}

	return game, nil
}
