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

func leftInfo(dailyCommitsMonthCount, endOfMonth, commitsYearCount int, tree, climber string) string {
	return fmt.Sprintf(`
  <g>
    <g>
      <g transform="translate(5,75)">%s</g>
      <text x="48" y="85" font-size="11" dominant-baseline="middle" text-anchor="start" fill="white">commit day</text>
      <text x="50" y="100" font-size="11" dominant-baseline="middle" text-anchor="start" fill="white">per month</text>
      <text x="125" y="90" font-size="10" dominant-baseline="middle" text-anchor="start" fill="white">%d/%d</text>
    </g>
    <g>
      <g transform="translate(5,140)">%s</g>
      <text x="40" y="150" font-size="11" dominant-baseline="middle" text-anchor="start" fill="white">Total Commits</text>
      <text x="60" y="165" font-size="11" dominant-baseline="middle" text-anchor="start" fill="white">(2024)</text>
      <text x="130" y="155" font-size="10" dominant-baseline="middle" text-anchor="start" fill="white">%d</text>

    </g>
  </g>
  `, tree, dailyCommitsMonthCount, endOfMonth, climber, commitsYearCount)
}

func rightInfo(grass string) string {
	return fmt.Sprintf(`
  <g>
    <rect x="175" y="55" width="150" height="130" fill="#141321" rx="5" ry="5" stroke="white" stroke-width="1"/> 
    <g transform="translate(158 167)">
      <g>%s</g>
    </g>
  </g>
  `, grass)
}

func endSVG() string {
	return `</svg>`
}

func GenerateCard(username string, dailyCommitsSince1MonthCount, dailyCommitsMonthCount, endOfMonth, commitsYearCount int) string {
	width := 340
	height := 200
	viewBox := fmt.Sprintf("0 0 %d %d", width, height)
	bgColor := "#141321"
	count := 31
	grassMountain := generateMountain(count)
	svg := startSVG(width, height, viewBox)
	svg += rect(width, height, bgColor)
	svg += title(username, Mountain)
	svg += leftInfo(dailyCommitsMonthCount, endOfMonth, commitsYearCount, Tree, Climber)
	svg += rightInfo(grassMountain)

	svg += endSVG()

	return svg
}
