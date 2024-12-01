package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func main() {

	server := gin.Default()

	server.GET("/ping", handleRoute)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func handleRoute(c *gin.Context) {
	response := Response{
		Message: "Hello world!",
	}

	c.json(200, response)
}
