package main

import (
	"log"
	"fmt"
	"net/http"
)

const webPort = "80"

type AppConfig struct {}

func main() {
	//set up config
	app := AppConfig{}
	app.routes()
	err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), nil)

	if err != nil {
		log.Panic(err)
	}
}