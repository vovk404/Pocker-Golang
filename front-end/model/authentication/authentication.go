package authentication

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"errors"
)

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginJsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Login(entry LoginRequest) error {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	// call the service
	request, err := http.NewRequest("POST", "http://localhost:5442/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Couldn`t create an authentication service")
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusAccepted {
		return errors.New("Wrong login or password")
	}

	//creare a variable we'll read response.Body into
	var jsonFromService LoginJsonResponse

	//decode the json
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		return err
	}
	if jsonFromService.Error == true {
		return errors.New("Wrong login or password")
	}

	return nil
}
