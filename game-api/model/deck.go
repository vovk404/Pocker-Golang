package model

import (
	"strconv"
	"math/rand"
)

type Cards struct {}

func (cards *Cards) getNewDeck() []string {
	deck := []string{}
	for i := 2; i <= 10; i++ {
		element := strconv.FormatInt(int64(i), 10)
		clubs    := element + "_clubs"
		diamonds := element + "_diamonds"
		hearts   := element + "_hearts"
		spades   := element + "_spades"
		deck = append(deck, clubs, diamonds, hearts, spades);
	}
	suitesOfCards := []string{"jack", "quieen", "king", "ace"}

	for _, val := range suitesOfCards {
		clubs    := val + "_clubs"
		diamonds := val + "_diamonds"
		hearts   := val + "_hearts"
		spades   := val + "_spades"
		deck = append(deck, clubs, diamonds, hearts, spades);
	}

	return cards.shuffleDeck(deck)
}

func (cards *Cards) shuffleDeck(deck []string) []string {
	rand.Shuffle(len(deck), func(i, j int) {
        deck[i], deck[j] = deck[j], deck[i]
    })

	return deck
}