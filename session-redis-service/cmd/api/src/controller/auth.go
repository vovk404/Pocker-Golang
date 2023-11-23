package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RedisLoginRequest struct {
	Id int `json:"id"`
	Email string `json:"email"`
}

type Authentication struct{}

func (auth *Authentication) Login(c *gin.Context) {
	//var redisRequest RedisLoginRequest
	session := sessions.Default(c)
	request := c.Request
	decoder := json.NewDecoder(request.Body)
    var createSessionRequest RedisLoginRequest
    if err := decoder.Decode(&createSessionRequest);  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong Params, session is not created",
		})
		return
	}
	
	session.Set("id", createSessionRequest.Id)
	session.Set("email", createSessionRequest.Email)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully",
	})
}

func (auth *Authentication) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}