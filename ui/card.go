package ui

import (
	"fmt"
)

func startSVG(width, height int, viewBox string) string {
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="%s">`, width, height, viewBox)
}

func rect(width, height int, bgColor string) string {
	return fmt.Sprintf(`<rect width="%d" height="%d" fill="%s" rx="5" ry="5" stroke="white" stroke-width="1"/>`, width, height, bgColor)
}

func title(username string, mountain string) string {
	return fmt.Sprintf(`
    <g style="padding: 2em">
    <text x="0" y="30" font-size="15" dominant-baseline="middle" text-anchor="middle" fill="white">%s's</text>
    <g transform="translate(130,12)">%s</g>
    </g>
    `, username, mountain)
}

func endSVG() string {
	return `</svg>`
}

func GenerateCard(username string) string {
	width := 340
	height := 200
	viewBox := fmt.Sprintf("0 0 %d %d", width, height)
	bgColor := "#141321"

	svg := startSVG(width, height, viewBox)
	svg += rect(width, height, bgColor)
	svg += title(username, Mountain)

	svg += endSVG()

	return svg
}
