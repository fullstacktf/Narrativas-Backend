package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStory : endpoint that returns a story by ID
func GetStory(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// DeleteStory : endpoint that deletes a story by ID
func DeleteStory(c *gin.Context) {
	id := c.Params.ByName("id")
	message := "Story with id " + id + " was deleted."
	c.String(http.StatusOK, message)
}

// PostStory : endpoint that creates a story
func PostStory(c *gin.Context) {
	message := "Story created"
	c.String(http.StatusOK, message)
}