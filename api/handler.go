package api

import (
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

	svg := generateSVG()

	c.Data(http.StatusOK, "image/svg+xml", []byte(svg))
	// c.JSON(http.StatusOK, gin.H{"totalCommit": totalCommit})

}

// wip
func generateSVG() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" width="200" height="100">
            <rect x="10" y="10" width="180" height="80" fill="lightblue" stroke="black" stroke-width="2"/>
            <text x="100" y="50" text-anchor="middle" alignment-baseline="middle" fill="black">Hello, SVG!</text>
        </svg>`
}
