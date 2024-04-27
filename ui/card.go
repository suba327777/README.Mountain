package ui

import (
	"fmt"
)

func startSVG() string {
	return `<svg xmlns="http://www.w3.org/2000/svg">`
}

func rect(width, height int, bgColor string) string {
	return fmt.Sprintf(`<rect x="10" y="10" width="%d" height="%d" fill="%s" rx="5" ry="5" stroke="black" stroke-width="2"/>`, width, height, bgColor)
}

func title(username string) string {
	return fmt.Sprintf(`<text x="110" y="30" font-size="15" dominant-baseline="middle" text-anchor="middle" fill="white">%s's mountain</text>`, username)
}

func endSVG() string {
	return `</svg>`
}
func GenerateCard(username string) string {

	width := 340
	height := 200
	bgColor := "#141321"

	svg := startSVG()
	svg += rect(width, height, bgColor)
	svg += title(username)

	svg += endSVG()
	fmt.Println(svg)

	return svg
}
