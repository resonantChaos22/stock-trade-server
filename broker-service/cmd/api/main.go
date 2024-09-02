package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	WEB_PORT = "8080"
)

func main() {
	r := gin.New()

	//	midlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//	routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	r.Run(fmt.Sprintf(":%s", WEB_PORT))
}
