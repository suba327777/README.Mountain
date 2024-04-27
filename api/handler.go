package api

import (
	"net/http"
	"time"

	"github-readme-mountain/ui"

	"github.com/gin-gonic/gin"
)

func MountainHandler(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	toDate := time.Now()
	fromDate := toDate.AddDate(0, -1, 0)
	user, err := getUserData(username, fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch GitHub data"})
		return
	}

	print(user.DailyCommitsSince1Month)

	svg := ui.GenerateCard(username)

	c.Data(http.StatusOK, "image/svg+xml", []byte(svg))
	c.JSON(http.StatusOK, gin.H{"totalCommit": user})
}
