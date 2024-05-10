package main

import (
	"fmt"
	"os"

	"github-readme-mountain/api"
)

func main() {

	svg, err := api.MountainHandler()
	if err != nil {
		fmt.Printf("Error calling MountainHandler: %v\n", err)
		return
	}

	// ファイルパスを設定
	directory := "./mountain-output"
	filePath := "./mountain-output/mountain.svg"

	// ディレクトリが存在しない場合は作成
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, 0o755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
	}

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
