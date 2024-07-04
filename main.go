package main

import (
	"fmt"

	"github-readme-mountain/api"
	"github-readme-mountain/utils"
)

func main() {
	svg, err := api.MountainHandler()
	if err != nil {
		fmt.Printf("Error calling MountainHandler: %v\n", err)
		return
	}

	err = utils.WriteSvgFile(svg)
	if err != nil {
		fmt.Printf("Error calling WriteSvgFile: %v\n", err)
		return
	}

}
