package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthenticationInterface interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}