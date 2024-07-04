package utils

import (
	"fmt"
	"os"
)

func WriteSvgFile(svg []byte) error {
	directory := "./mountain-output"
	filePath := directory + "/mountain.svg"

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, 0o755)
		if err != nil {
			return fmt.Errorf("Error creating directory: %v", err)
		}
	}

	svgFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		return fmt.Errorf("Error opening mountain.svg: %v", err)
	}
	defer svgFile.Close()

	_, err = svgFile.Write(svg)
	if err != nil {
		return fmt.Errorf("Error writing to SVG file: %v", err)
	}

	fmt.Println("SVG has been saved successfully.")
	return nil
}
