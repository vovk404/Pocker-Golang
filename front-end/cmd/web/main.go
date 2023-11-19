package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"front-end/model/authentication"
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
		render(w, "login_page.gohtml", map[string]interface{}{})
	})
	
	http.HandleFunc("/new_game", func(w http.ResponseWriter, r *http.Request) {
		err := authentication.CheckRedisSession(r)
		if err != nil {
			//TODO add cookies messages
			log.Println("Redis session wasn`t found: ", err.Error())
			render(w, "login_page.gohtml", map[string]interface{}{})
			return
		}
		players, err := strconv.ParseInt(r.PostFormValue("players"), 10, 64)
		if err != nil {
			//log.Panic("Players wasn`t passed")
			http.Redirect(w, r, "http://localhost/", http.StatusSeeOther)
		}
		var playersArray []int
		for i := 1; i <= int(players); i++ {
			playersArray = append(playersArray, i)
		}
		data := map[string]interface{}{"players": playersArray}
		render(w, "game/new_game.gohtml", data)
		return
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var loginRequest authentication.LoginRequest

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Could not parse form.", http.StatusBadRequest)
			return
		}
		loginRequest.Email = r.PostForm.Get("email")
		loginRequest.Password = r.PostForm.Get("password")

		err, user := authentication.Login(loginRequest)
		if err != nil {
			log.Println("Error during login: ", err.Error())
			render(w, "login_page.gohtml", map[string]interface{}{})
			return
		}
		//add redis session
		var redisRequest authentication.RedisLoginRequest
	    redisRequest.Id    = user.ID
	    redisRequest.Email = user.Email
	    err, cookie := authentication.CreateRedisSession(redisRequest)
	    if err != nil || cookie.Name == "" {
			log.Println("Error during session creation")
			render(w, "login_page.gohtml", map[string]interface{}{})
			return
    	}
		http.SetCookie(w, &cookie)
		
		render(w, "start.game.gohtml", map[string]interface{}{})
		return
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