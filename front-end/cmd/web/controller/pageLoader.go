package controller

import (
	"log"
	"net/http"
	"strconv"
	"front-end/cmd/web/model/authentication"
	"front-end/cmd/web/model/view"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
		view.RenderTemplate(w, "login_page.gohtml", map[string]interface{}{})
		return
	}
		//add redis session
	var redisRequest authentication.RedisLoginRequest
	redisRequest.Id    = user.ID
	redisRequest.Email = user.Email
	err, cookie := authentication.CreateRedisSession(redisRequest)
	if err != nil || cookie.Name == "" {
		log.Println("Error during session creation")
		view.RenderTemplate(w, "login_page.gohtml", map[string]interface{}{})
		return
 	}
	http.SetCookie(w, &cookie)
		
	view.RenderTemplate(w, "start.game.gohtml", map[string]interface{}{})
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	err := authentication.Logout(r)
	if err != nil {
		log.Println("Error during logout: ", err.Error())
		return
	}
		
	view.RenderTemplate(w, "login_page.gohtml", map[string]interface{}{})
	return
}

func OpenMainPage(w http.ResponseWriter, r *http.Request) {
	err := authentication.CheckRedisSession(r)
	if err != nil {
		view.RenderTemplate(w, "login_page.gohtml", map[string]interface{}{})
		return
	}
	view.RenderTemplate(w, "game/new_game.gohtml", map[string]interface{}{})
}

func OpenGamePage(w http.ResponseWriter, r *http.Request) {
	err := authentication.CheckRedisSession(r)
	if err != nil {
			//TODO add cookies messages
		log.Println("Redis session wasn`t found: ", err.Error())
		view.RenderTemplate(w, "login_page.gohtml", map[string]interface{}{})
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
	view.RenderTemplate(w, "game/new_game.gohtml", data)
	return
}