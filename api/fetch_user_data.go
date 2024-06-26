package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"github-readme-mountain/object"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func getUserData(username string, from, to time.Time) (object.UserData, error) {
	token := os.Getenv("GITHUB_TOKEN")

	auth := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	httpClient := oauth2.NewClient(context.Background(), auth)

	client := githubv4.NewClient(httpClient)

	var query struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int
					Weeks              []struct {
						ContributionDays []struct {
							Date              ContributionDaysTime
							ContributionCount int
						}
					}
				} `graphql:"contributionCalendar"`
			} `graphql:"contributionsCollection(from: $from, to: $to)"`
		} `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String(username),
		"from":     githubv4.DateTime{Time: from.Local().UTC()},
		"to":       githubv4.DateTime{Time: to.Local().UTC()},
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println("Error request data:", err)
		return object.UserData{}, fmt.Errorf("error requesting data: %w", err)
	}

	dailyCommits := make(map[string]int)
	for _, week := range query.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			dailyCommits[day.Date.Format("2006-01-02")] = day.ContributionCount
		}
	}

	user := object.UserData{
		DailyCommitsPeriod: dailyCommits,
	}

	return user, nil
}

type ContributionDaysTime struct {
	time.Time
}

func (m *ContributionDaysTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	tt, err := time.Parse(`"2006-01-02"`, string(data))
	*m = ContributionDaysTime{tt}
	return err
}
