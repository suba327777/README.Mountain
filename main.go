package main

import (
	"flag"
	"fmt"

	"github-readme-mountain/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getHello() string {
	return "hello world"
}

var (
	userName string
)

func init() {
	flag.StringVar(&userName, "userName", "userName", "Input your user name")
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
	flag.Parse()
	fmt.Println(flag.Args(), validateArgs())
}

func validateArgs() error {
	flag.Parse()
	if !validateUserName() {
		return fmt.Errorf("error: Invalid user name %s", userName)
	}
	return nil
}

// TODO: Implement user check logic
func validateUserName() bool {
	return userName == "niwaniwa"
}
