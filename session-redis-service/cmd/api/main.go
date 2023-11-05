package main

import (
	"session-redis-service/cmd/api/src/middleware"
	"session-redis-service/cmd/api/src/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Authentication controller.AuthenticationInterface
}

func main() {
	//set up config
	app := Config{
		Authentication: &controller.Authentication{},
	}
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "redis:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.POST("/login", app.Authentication.Login)
	r.GET("/logout", app.Authentication.Logout)
	auth := r.Group("/auth")
	auth.Use(middleware.Authentication())
	{
		auth.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Everything is ok",
			})
		})
	}
	r.Run(":4111")
}