package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	handleRoutes()

	fmt.Println("Starting front end service on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}

func handleRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "start.game.gohtml", map[string]interface{}{})
	})
	http.HandleFunc("/new_game", func(w http.ResponseWriter, r *http.Request) {
		participants := r.PostFormValue("participants")
		if participants == "" {
			//log.Panic("Participants wasn`t passed")
			http.Redirect(w, r, "http://localhost/", http.StatusSeeOther)
		}
		data := map[string]interface{}{"participants": r.PostFormValue("participants")}
		render(w, "game/new_game.gohtml", data)
	})
}

func render(w http.ResponseWriter, t string, data map[string]interface{}) {

	partials := []string{
		fmt.Sprintf("./cmd/web/templates/%s", t),
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	tmpl, err := template.ParseFiles(partials...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}