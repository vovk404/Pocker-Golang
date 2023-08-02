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
		clubs    := element + "clubs"
		diamonds := element + "diamonds"
		hearts   := element + "hearts"
		spades   := element + "spades"
		deck = append(deck, clubs, diamonds, hearts, spades);
	}
	suitesOfCards := []string{"jack", "quieen", "king", "ace"}

	for _, val := range suitesOfCards {
		clubs    := val + "clubs"
		diamonds := val + "diamonds"
		hearts   := val + "hearts"
		spades   := val + "spades"
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