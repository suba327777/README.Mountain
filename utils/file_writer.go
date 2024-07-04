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
			return fmt.Errorf("Error creating directory: %v\n", err)
		}
	}

	svgFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		return fmt.Errorf("Error opening mountain.svg: %v\n", err)
	}
	defer svgFile.Close()

	_, err = svgFile.Write(svg)
	if err != nil {
		return fmt.Errorf("Error writing to SVG file: %v\n", err)
	}

	fmt.Println("SVG has been saved successfully.")
	return nil
}

func WriteReadmeFile() error {
	directory := "./mountain-output"
	readmePath := directory + "/README.md"

	readmeFile, err := os.OpenFile(readmePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		return fmt.Errorf("Error opening README.md file: %v\n", err)
	}
	defer readmeFile.Close()

	username := os.Getenv("USERNAME")
	theme := os.Getenv("THEME")
	branch, err := getHeadBranch()
	if err != nil {
		return fmt.Errorf("Error getting HEAD branch: %v\n", err)
	}

	startReadme := fmt.Sprintf("## %s\n\n", theme)
	showSvg := "![](./mountain.svg)\n\n"
	usageText := "### Now you can add this to your markdown\n"
	startCodeBlock := "```\n"
	imageUrl := fmt.Sprintf(`<img src="https://raw.githubusercontent.com/%s/%s/%s/mountain-output/mountain.svg" />`, username, username, branch)
	endCodeBlock := "\n```"
	readmeContent := startReadme + showSvg + usageText + startCodeBlock + imageUrl + endCodeBlock

	fmt.Println(readmeContent)

	_, err = readmeFile.Write([]byte(readmeContent))
	if err != nil {
		return fmt.Errorf("Error writing to README.md file: %v\n", err)
	}

	fmt.Println("README has been saved successfully.")
	return nil
}
