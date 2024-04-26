package main

import (
	"fmt"
	"github-readme-mountain/api"
	"github.com/gin-gonic/gin"
)

func getHello() string {
	return "hello world"
}

func main() {
	fmt.Println(getHello())

	r := gin.Default()

	r.GET("/mountain", api.MountainHandler)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("server err", err)
	}
}
