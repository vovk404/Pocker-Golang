package main

import (
	"fmt"
	"net/http"
	"log"
)

const webPort = "80"

type AppConfig struct {
	
}

func main() {
	log.Println("Starting game-api service")

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