package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHello() string {
	return "hello world"
}

func mountainHandler(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	c.String(200, "Hello %s!", username)
}

func main() {
	fmt.Println(getHello())

	r := gin.Default()

	r.GET("/mountain", mountainHandler)

	fmt.Println("server is runnning")
	r.Run(":8080")
}
