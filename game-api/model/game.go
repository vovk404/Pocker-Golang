package model

import (
	"log"
	"encoding/json"
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
	game.createRedisSession()

	log.Println("cards set up finished")
	return game
}

func (game *Game) createRedisSession() {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB: 0,
	})
	pong, err := client.Ping().Result()
	log.Println(pong, err)

	json, err := json.Marshal(game)
    if err != nil {
        log.Println("Error during marshaling game object", err)
    }

	err = client.Set("CurrentGame", json, 0).Err()
	if err != nil {
		log.Println("Error during saving the game to the redis cache: ", err)
	}
}