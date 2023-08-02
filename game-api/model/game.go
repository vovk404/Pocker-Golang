package model

import (
	"log"
)

type Game struct {
	Cards Cards
}

func CreateGame(palyers int) Game {
	
	//create a new deck of cards
	game := Game {
		Cards: Cards{},
	}
	deck := game.Cards.getNewDeck()
	log.Println("Shuffled deck: ", deck)

	//give 2 cards to each players

	//open preflop

	return game
}