package ui

import (
	"fmt"
	"strings"
)

// wip
func mountainColor(todayCommit int) string {
	var color string
	if todayCommit >= 20 {
		color = "#00cc00"
	} else if todayCommit >= 10 {
		green := 204 + (todayCommit-10)*5
		color = fmt.Sprintf("#%02x%02x%02x", 255, green, 0)
	} else if todayCommit >= 1 {
		green := 153 + todayCommit*5
		color = fmt.Sprintf("#%02x%02x%02x", 255, green, 0)
	} else {
		color = "#ffffff"
	}
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1000 1000" height="20" width="20"><path fill="%s" d="M860,880H80c0-173.49,24.31-307,75.31-427.75,43.1-102,121.1-216.84,196-292.25h44.28c-19.43,64.11-53,213.82-59.95,295.3,38.81-100.94,112-209.91,195.67-295.3h44.28c-21.52,73.42-53.59,213.49-60,295.3,35-102.95,119.9-215.9,195.67-295.3h44.28c-23.28,82.4-53.93,213.87-60,295.3C734.48,354.16,814.13,236.08,891.34,160h44.28C875.88,390.82,860,523.39,860,880ZM120.14,840H820c.42-262.13,11.61-406.94,57.64-607.64-99.49,117.66-163,244-194,388.11H644.12c6.11-136.68,22.56-252.89,53.55-388.12-99.5,117.67-163,244-194,388.12H464.12c6.11-136.68,22.56-252.89,53.55-388.12-99.5,117.67-163,244-194,388.12H284.12c6.11-136.68,22.56-252.89,53.55-388.12C275.38,306,226.21,386.82,192.19,467.75,145.81,578.06,122.18,700,120.14,840Z"></path></svg>`, color)
}

func generateMountain(count int) string {
	var results strings.Builder

	x := 0
	y := 0

	for i := 1; i <= count; i++ {
		x += 20
		if i == 8 {
			x -= 130
			y -= 18
			results.WriteString(generateTrianglePath(x, y))
		} else if i == 14 {
			x -= 110
			y -= 15
			results.WriteString(generateTrianglePath(x, y))
		} else if i == 19 {
			x -= 90
			y -= 15
			results.WriteString(generateTrianglePath(x, y))
		} else if i == 23 {
			x -= 70
			y -= 15
			results.WriteString(generateTrianglePath(x, y))
		} else if i == 26 {
			x -= 50
			y -= 15
			results.WriteString(generateTrianglePath(x, y))
		} else if i == 28 {
			x -= 30
			y -= 15
			results.WriteString(generateTrianglePath(x, y))
		} else if i == 29 {
			x -= 50
			y -= 5
			results.WriteString(generateStarPath(x, y))
		} else if i == 30 {
			x -= 30
			y += 20
			results.WriteString(generateStarPath(x, y))
		} else if i == 31 {
			x += 30
			y -= 30
			results.WriteString(generateFlagPath(x, y))
		} else {
			results.WriteString(generateTrianglePath(x, y))
		}

	}
	return results.String()
}
