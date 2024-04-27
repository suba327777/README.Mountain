package object

type UserData struct {
	DailyCommitsSince1Month             map[string]int
	TotalPullRequestContributions       int
	TotalPullRequestReviewContributions int
	TotalRepositoryContributions        int
}
