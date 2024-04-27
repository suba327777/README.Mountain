package main

import (
	"fmt"

	"github-readme-mountain/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getHello() string {
	return "hello world"
}

func main() {
	fmt.Println(getHello())

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("read error: %v", err)
	}

	r := gin.Default()

	r.GET("/mountain", api.MountainHandler)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("server err", err)
	}
}
