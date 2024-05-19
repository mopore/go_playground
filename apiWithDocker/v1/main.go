package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

type User struct {
	Id int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func PostUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	log.Printf("Recieved user: %v\n", user)
	if user.Firstname == "" || user.Lastname == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "User created successfully!",
	})
}


func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", PostUser)
	}

	log.Println("Starting server...")
	router.Run(port)
}
