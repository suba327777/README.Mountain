package api

func commitCountDailySince1Month(DailyCommitsSince1Month map[string]int) int {
	count := 0
	for _, commits := range DailyCommitsSince1Month {
		if commits >= 1 {
			count++
		}
	}
	return count
}
