package api

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

func GetUserData(username string) (string, error) {

	apiURL := "https://api.github.com/graphql"

	token := os.Getenv("GITHUB_TOKEN")

	auth := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	httpClient := oauth2.NewClient(context.Background(), auth)

	client := graphql.NewClient(apiURL, httpClient)

	var query struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions int
			} `graphql:"contributionsCollection"`
		} `graphql:"user(login: $username)"`
	}
	variables := map[string]interface{}{
		"username": username,
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return "", err
	}
	fmt.Println(query.User.ContributionsCollection.TotalCommitContributions)
	return strconv.Itoa(query.User.ContributionsCollection.TotalCommitContributions), nil

}
