package ui

import "fmt"

func formatNumber(num int) string {
	if num >= 1000 {
		return fmt.Sprintf("%.1fk", float64(num)/1000)
	} else {
		return fmt.Sprintf("%d", num)
	}
}
