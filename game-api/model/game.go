package model

import (
	"log"
)

type Game struct {
	Cards Cards
	CurrentDeck []string
}

func CreateGame(palyers int) Game {
	
	//create a new deck of cards
	game := Game {
		Cards: Cards{},
	}
	deck := game.Cards.getNewDeck()
	game.CurrentDeck = deck
	log.Println("Shuffled deck: ", deck)

	//give 2 cards to each players

	//open preflop

	return game
}