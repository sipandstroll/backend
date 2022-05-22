package main

import (
	"github.com/gin-gonic/gin"
	"helloworld/config"
	"helloworld/middleware"
	"net/http"
)

func main() {
	// initialize gin Engine
	r := gin.Default()

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	r.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})

	r.Use(middleware.AuthMiddleware)

	r.GET("/helloAuth", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "user"})
	})

	r.Run(":3000")

	_, _ = r, firebaseAuth
}
