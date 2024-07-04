package ui

import (
	"fmt"
	"os"
)

func startSVG() string {
	return fmt.Sprintf(`
  <svg xmlns="http://www.w3.org/2000/svg" width="340" height="200" viewBox="0 0 340 200" >
    <a xlink:href="https://github.com/suba327777/README.Mountain">
  `)
}

func rect(bgColor, borderColor string) string {
	return fmt.Sprintf(`<rect width="340" height="200" fill="%s" rx="5" ry="5" stroke="%s" stroke-width="1"/>`, bgColor, borderColor)
}

func title(mountainIcon, textColor, username string) string {
	return fmt.Sprintf(`
    <g>
      <g transform="translate(5,8)">%s</g>
        <text x="45" y="26" font-size="18" dominant-baseline="middle" text-anchor="start"  fill="%s">%s</text>
    </g>
    `, mountainIcon, textColor, username)
}

func leftInfo(dailyCommitsMonthCount, endOfMonth int, treeIcon, climberIcon, textColor, formatCommitsMonthCount string) string {
	return fmt.Sprintf(`
  <g>
    <g>
      <g transform="translate(5,75)">%s</g>
      <text x="43" y="85" font-size="13" dominant-baseline="middle" text-anchor="start"  fill="%s">Commit Day</text>
      <text x="47" y="100" font-size="13" dominant-baseline="middle" text-anchor="start"  fill="%s">Per Month</text>
      <text x="131" y="93" font-size="12" dominant-baseline="middle" text-anchor="start"  fill="%s">%d/%d</text>
    </g>
    <g>
      <g transform="translate(5,140)">%s</g> <text x="54" y="150" font-size="13" dominant-baseline="middle" text-anchor="start"  fill="%s">Commits</text>
      <text x="49" y="165" font-size="13" dominant-baseline="middle" text-anchor="start"  fill="%s">Per Month</text>
      <text x="143" y="157" font-size="12" dominant-baseline="middle" text-anchor="start"  fill="%s">%s</text>
    </g>
  </g>
  `, treeIcon, textColor, textColor, textColor, dailyCommitsMonthCount, endOfMonth, climberIcon, textColor, textColor, textColor, formatCommitsMonthCount)
}

func rightInfo(bgColor, borderColor, grass string) string {
	return fmt.Sprintf(`
  <g>
    <rect x="175" y="55" width="150" height="130" fill="%s" rx="5" ry="5" stroke="%s" stroke-width="1"/> 
    <g transform="translate(158 167)">
      <g>%s</g>
    </g>
  </g>
  `, bgColor, borderColor, grass)
}

func endSVG() string {
	return `
    </a>
  </svg>`
}

func GenerateCard(username string, dailyCommitsSince1MonthCount, dailyCommitsMonthCount, endOfMonth, commitsMonthCount int) string {
	themeName := os.Getenv("THEME")
	theme := getTheme(themeName)

	mountainIcon := changeIconColor(Mountain, theme.MountainIconColor)
	climberIcon := changeIconColor(Climber, theme.IconColor)
	treeIcon := changeIconColor(Tree, theme.IconColor)
	grassMountain := generateMountain(dailyCommitsSince1MonthCount, theme.Name, theme.TriangleMountainColor)
	formatCommitsMonthCount := formatNumber(commitsMonthCount)

	svg := startSVG()
	svg += rect(theme.BgColor, theme.BorderColor)
	svg += title(mountainIcon, theme.TitleColor, username)
	svg += leftInfo(dailyCommitsMonthCount, endOfMonth, treeIcon, climberIcon, theme.TextColor, formatCommitsMonthCount)
	svg += rightInfo(theme.BgColor, theme.BorderColor, grassMountain)
	svg += endSVG()

	return svg
}
