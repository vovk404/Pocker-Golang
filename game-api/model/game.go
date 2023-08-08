package model

import (
	"log"
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