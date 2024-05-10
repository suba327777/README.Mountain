package ui

import (
	"strings"
)

func generateMountain(count int, themeName, triangleMountainColor string) string {
	var results strings.Builder

	x := 0
	y := 0

	for i := 1; i <= count; i++ {
		x += 20
		if i == 8 {
			x -= 130
			y -= 18
			results.WriteString(generateTrianglePath(x, y, triangleMountainColor))
		} else if i == 14 {
			x -= 110
			y -= 15
			results.WriteString(generateTrianglePath(x, y, triangleMountainColor))
		} else if i == 19 {
			x -= 90
			y -= 15
			results.WriteString(generateTrianglePath(x, y, triangleMountainColor))
		} else if i == 23 {
			x -= 70
			y -= 15
			results.WriteString(generateTrianglePath(x, y, triangleMountainColor))
		} else if i == 26 {
			x -= 50
			y -= 15
			results.WriteString(generateTrianglePath(x, y, triangleMountainColor))
		} else if i == 28 {
			x -= 30
			y -= 15
			results.WriteString(generateTrianglePath(x, y, triangleMountainColor))
		} else if i == 29 {
			x -= 50
			y -= 5
			switch themeName {
			case "default":
				results.WriteString(generateSunPath(x, y))
			case "dark":
				results.WriteString(generateStarPath(x, y))
			}
		} else if i == 30 {
			x -= 30
			y += 20
			switch themeName {
			case "default":

			case "dark":
				results.WriteString(generateStarPath(x, y))
			}
		} else if i == 31 {
			x += 30
			y -= 30
			results.WriteString(generateFlagPath(x, y))
		} else {
			results.WriteString(generateTrianglePath(x, y, triangleMountainColor))
		}

	}
	return results.String()
}
