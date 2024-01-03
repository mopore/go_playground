package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var myOS = os.Getenv("OS")

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, fmt.Sprintf("Hello, %s", myOS))
    })
    r.Run(":8080")
}
