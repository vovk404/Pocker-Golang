package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"errors"
)

type RedisLoginRequest struct {
	Id int `json:"id"`
	Email string `json:"email"`
}

func CreateRedisSession(entry RedisLoginRequest) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	request, err := http.NewRequest("POST", "http://session-redis-service:4111/login", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Could not create customer session request")
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Response erro from session redis service: ", err.Error())
		return err
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusOK {
		log.Println("Response erro from session redis service, status is not OK")
		return errors.New("Error during customer session creation")
	}

	response.Cookies()

	return nil
}