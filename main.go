package main

import (
	"fmt"
	"github-readme-mountain/api"
	"os"

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

	svg, err := api.MountainHandler()

	if err != nil {
		fmt.Printf("Error calling MountainHandler: %v\n", err)
		return
	}

	// ファイルパスを設定
	directory := "./output"
	filePath := "./output/output.svg"

	os.Mkdir(directory, os.ModePerm)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(svg)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("SVG has been saved successfully.")

}
