package ui

import (
	"fmt"
)

func startSVG(width, height int, viewBox string) string {
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="%s" >`, width, height, viewBox)
}

func rect(width, height int, bgColor string) string {
	return fmt.Sprintf(`<rect width="%d" height="%d" fill="%s" rx="5" ry="5" stroke="white" stroke-width="1"/>`, width, height, bgColor)
}

func title(username string, mountain string) string {
	return fmt.Sprintf(`
    <g>
      <g transform="translate(5,8)">%s</g>
        <text x="45" y="26" font-size="16" dominant-baseline="middle" text-anchor="start" fill="white">%s</text>
    </g>
    `, mountain, username)
}

func leftInfo(width int, tree, climber, Backpack string) string {
	return fmt.Sprintf(`
  <g>
    <g>
      <g transform="translate(5,65)">%s</g>
      <text x="35" y="75" font-size="11" dominant-baseline="middle" text-anchor="start" fill="white">test: %d</text>
    </g>
    <g>
      <g transform="translate(5,110)">%s</g>
      <text x="35" y="120" font-size="11" dominant-baseline="middle" text-anchor="start" fill="white">test: %d</text>
    </g>
    <g>
      <g transform="translate(5,155)">%s</g>
      <text x="35" y="165" font-size="11" dominant-baseline="middle" text-anchor="start" fill="white">test: %d</text>
    </g>
  </g>
  `, tree, width, climber, width, Backpack, width)
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
	svg += leftInfo(width, Tree, Climber, Backpack)

	svg += endSVG()

	return svg
}
