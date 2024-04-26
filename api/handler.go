package api

import (
	"github-readme-mountain/ui"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MountainHandler(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	totalCommit, err := fetchGrass(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch GitHub data"})
		return
	}

	print(totalCommit)

	svg := ui.GenerateCard()

	c.Data(http.StatusOK, "image/svg+xml", []byte(svg))
	// c.JSON(http.StatusOK, gin.H{"totalCommit": totalCommit})

}
