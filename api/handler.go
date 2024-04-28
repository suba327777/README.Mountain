package api

import (
	"errors"
	"os"
	"time"

	"github-readme-mountain/ui"
)

func MountainHandler() ([]byte, error) {

	username := os.Getenv("USERNAME")
	if len(username) == 0 {
		return nil, errors.New("USERNAME is not set")
	}

	toDate := time.Now()
	fromDate := toDate.AddDate(0, -1, 0)
	user, err := getUserData(username, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	dailyCommitsSince1MonthCount := commitCountDailySince1Month(user.DailyCommitsSince1Month)

	svg := ui.GenerateCard(username, dailyCommitsSince1MonthCount)

	return []byte(svg), nil
}
