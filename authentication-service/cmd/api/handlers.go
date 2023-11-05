package main

import (
	"authentication/data"
	"authentication/services"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool      `json:"error"`
	Message string    `json:"message"`
	User    data.User `json:"data,omitempty"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPaylod struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var redisRequest services.RedisLoginRequest

	err := app.readJSON(w, r, &requestPaylod)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	user, err := app.Repo.GetByEmail(requestPaylod.Email)
	/**
	TODO - delete it
	var newUser = data.User {
		Email: "andriy.v@overdose.digital",
	    FirstName: "Andriy",
     	LastName: "Vovk",
	    Password: "Qwerty123!",
	    Active: 1,
	}
	userId, err := app.Repo.Insert(newUser)
	log.Println("User created with id:", userId)
	**/
	if err != nil {
		app.errorJson(w, errors.New("user with such email is not present"), http.StatusUnauthorized)
		return
	}

	valid, err := app.Repo.PasswordMatches(requestPaylod.Password, *user)
	if err != nil || !valid {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}
	//add redis session
	redisRequest.Id    = user.ID
	redisRequest.Email = user.Email
	err = services.CreateRedisSession(redisRequest)
	if err != nil {
		app.errorJson(w, errors.New("Error during redis session creation"), http.StatusUnauthorized)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		User:    *user,
	}
	log.Println("Succesfully logged-in: ", user.Email)
	app.writeJson(w, http.StatusAccepted, payload)
}
