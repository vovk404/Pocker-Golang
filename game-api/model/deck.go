package model

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Cards struct {
	CurrentDeck []string
	PlayersCards map[string][]string
	Flop []string
}

type CardsJson struct {
	CurrentDeck []string `json:"CurrentDeck"`
	PlayersCards map[string][]string `json:"PlayersCards"`
	Flop []string `json:"Flop"`
}

func (cards *Cards) CreateNewDeck() {
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

	deck = cards.ShuffleDeck(deck)
	cards.CurrentDeck = deck
}

func (cards *Cards) ShuffleDeck(deck []string) []string {
	rand.Shuffle(len(deck), func(i, j int) {
        deck[i], deck[j] = deck[j], deck[i]
    })

	return deck
}

func (cards *Cards) PassCardsToPlayers(players int) {
	playersCards := map[string][]string{}
	for i := 1; i <= players; i++ {
		playersCards["player_" + fmt.Sprint(i)] = []string{
			cards.GetOneCardFromDeck(),
			cards.GetOneCardFromDeck(),
		}
	}
	cards.PlayersCards = playersCards
}

func (cards *Cards) OpenPreflop() {
	cards.Flop = cards.CurrentDeck[:3]
	cards.CurrentDeck = cards.CurrentDeck[3:]
}

func (cards *Cards) GetOneCardFromDeck() string {
	card := cards.CurrentDeck[0]
	cards.CurrentDeck = cards.CurrentDeck[1:]

	return card
}