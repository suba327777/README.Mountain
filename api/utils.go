package api

func dailyCommitsPeriodCount(dailyCommitsPeriodCount map[string]int) int {
	count := 0
	for _, commits := range dailyCommitsPeriodCount {
		if commits >= 1 {
			count++
		}
	}
	return count
}

func commitsPeriodCount(commitsPerodCount map[string]int) int {
	count := 0
	for _, commits := range commitsPerodCount {
		count += commits
	}

	return count
}
