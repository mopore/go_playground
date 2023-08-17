package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)


func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	log.Println("Starting server...")
	router.Run(port)
}
