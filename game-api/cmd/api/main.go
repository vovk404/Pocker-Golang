package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/go-session/session/v3"
	"github.com/go-session/redis/v3"
)

const webPort = "80"

type AppConfig struct {
	
}

func main() {
	log.Println("Starting game-api service")
	session.InitManager(
		session.SetStore(redis.NewRedisStore(&redis.Options{
			Addr: "redis:6379",
			Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
			DB: 0,
		})),
	)
	log.Println("Session init")

	//TODO connect to DB
	// conn := connectToDB()
	// if conn == nil {
	// 	log.Panic("Can`t connect to Postgres")
	// }

	//set up config
	app := AppConfig{}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}