package ui

import (
	"fmt"
)

func startSVG() string {
	return `<svg xmlns="http://www.w3.org/2000/svg">`
}

func rect(width, height int, bgColor string) string {
	return fmt.Sprintf(`<rect x="10" y="10" width="%d" height="%d" fill="%s" rx="5" ry="5" stroke="white" stroke-width="1"/>`, width, height, bgColor)
}

func title(username string, mountain string) string {
	return fmt.Sprintf(`
    <text x="70" y="30" font-size="15" dominant-baseline="middle" text-anchor="middle" fill="white">%s's</text>
    <g transform="translate(130,12)">%s</g>
    `, username, mountain)
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
	svg += title(username, Mountain)

	svg += endSVG()
	fmt.Println(svg)
	fmt.Println(Mountain)

	return svg
}
