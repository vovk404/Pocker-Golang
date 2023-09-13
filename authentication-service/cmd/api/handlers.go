package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	log.Println("Trying to authenticate")
	var requestPaylod struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPaylod)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	user, err := app.Repo.GetByEmail(requestPaylod.Email)
	/**
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
		app.errorJson(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	valid, err := app.Repo.PasswordMatches(requestPaylod.Password, *user)
	if err != nil || !valid {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}
	log.Println("Succesfully logged-in: ", user.Email)
	app.writeJson(w, http.StatusAccepted, payload)
}
