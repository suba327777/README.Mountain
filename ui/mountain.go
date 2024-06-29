package ui

import (
	"strings"
)

func generateMountain(count int, themeName, triangleMountainColor string) string {
	var results strings.Builder

	x := 0
	y := 0

	adjustments := map[int]struct {
		dx int
		dy int
		fn func(int, int) string
	}{
		8:  {dx: -130, dy: -18, fn: func(x, y int) string { return generateTrianglePath(x, y, triangleMountainColor) }},
		14: {dx: -110, dy: -15, fn: func(x, y int) string { return generateTrianglePath(x, y, triangleMountainColor) }},
		19: {dx: -90, dy: -15, fn: func(x, y int) string { return generateTrianglePath(x, y, triangleMountainColor) }},
		23: {dx: -70, dy: -15, fn: func(x, y int) string { return generateTrianglePath(x, y, triangleMountainColor) }},
		26: {dx: -50, dy: -15, fn: func(x, y int) string { return generateTrianglePath(x, y, triangleMountainColor) }},
		28: {dx: -30, dy: -15, fn: func(x, y int) string { return generateTrianglePath(x, y, triangleMountainColor) }},
		29: {dx: -50, dy: -5, fn: func(x, y int) string {
			switch themeName {
			case "default", "solarized":
				return generateBonusIconPath(x, y, Sun)
			case "dark", "onedark", "solarized_dark":
				return generateBonusIconPath(x, y, Star)
			case "sakura":
				return generateBonusIconPath(x, y, Sakura)
			case "maple":
				return generateBonusIconPath(x, y, maple)
			default:
				return ""
			}
		}},
		30: {dx: -30, dy: 20, fn: func(x, y int) string {
			switch themeName {
			case "dark":
				return generateBonusIconPath(x, y, Star)
			case "sakura":
				return generateBonusIconPath(x, y, Sakura)
			case "maple":
				return generateBonusIconPath(x, y, maple)
			default:
				return ""
			}
		}},
		31: {dx: 30, dy: -30, fn: func(x, y int) string { return generateBonusIconPath(x, y, Flag) }},
	}

	for i := 1; i <= count; i++ {
		x += 20
		if adj, ok := adjustments[i]; ok {
			x += adj.dx
			y += adj.dy
			results.WriteString(adj.fn(x, y))
		} else {
			results.WriteString(generateTrianglePath(x, y, triangleMountainColor))
		}
	}

	return results.String()
}
