package main

import (
	"fmt"
	"os"

	"github-readme-mountain/api"

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
	filePath := "./output/output.svg"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
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
