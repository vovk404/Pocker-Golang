package model

import (
	"log"
	"fmt"
	"github.com/go-redis/redis"
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
	crateRedisSession()

	log.Println("cards set up finished")
	return game
}

func crateRedisSession() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}