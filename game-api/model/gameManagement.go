package model

import (
	"log"
	"net/http"
	"encoding/json"
)

type Game struct {
	Cards Cards
}

func CreateGame(palyers int) *Game {
	game := Game {
		Cards: Cards{},
	}
	game.Cards.CreateNewDeck()
	game.Cards.PassCardsToPlayers(palyers)
	//TODO 
	savedGame, _ := SaveCurrentGame(game)
	return savedGame
}

func GetCurrentGame(w http.ResponseWriter, r *http.Request) (*Game, error) {
	customerGame, err := GetByUserId(3)
	
	if err != nil {
		log.Println("Customer game load error")
		return nil, err
	}
	game := Game{
		Cards: Cards{},
	}
	err = json.Unmarshal([]byte(customerGame.Data), &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func SaveCurrentGame(game Game) (*Game, error) {
	err := SaveGame(&game, 3)
	if err != nil {
		log.Println("Could not save a new game ", err)
	}

	log.Println("cards set up finished")
	return &game, nil
}
