package story

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStory : endpoint that returns a story by ID
func GetStory(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}