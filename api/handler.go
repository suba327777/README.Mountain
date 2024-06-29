package api

import (
	"errors"
	"os"
	"time"

	"github-readme-mountain/ui"
)

func MountainHandler() ([]byte, error) {
	username := os.Getenv("USERNAME2")
	if len(username) == 0 {
		return nil, errors.New("USERNAME is not set")
	}

	toDate := time.Now()
	fromDate := toDate.AddDate(0, -1, 0)
	user, err := getUserData(username, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	dailyCommitsSince1MonthCount := dailyCommitsPeriodCount(user.DailyCommitsPeriod)

	fromDate = time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
	toDate = time.Date(time.Now().Year(), time.Now().Month()+1, 1, 0, 0, 0, -1, time.UTC)
	user, err = getUserData(username, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	endOfMonth := toDate.Day()
	dailyCommitsMonthCount := dailyCommitsPeriodCount(user.DailyCommitsPeriod)
	commitsMonthCount := commitsPeriodCount(user.DailyCommitsPeriod)

	svg := ui.GenerateCard(username, dailyCommitsSince1MonthCount, dailyCommitsMonthCount, endOfMonth, commitsMonthCount)

	return []byte(svg), nil
}
