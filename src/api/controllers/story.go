package controllers

import (
	"net/http"

	model "github.com/fullstacktf/Narrativas-Backend/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Get : returns all the stories
func Get(c *gin.Context) {
	var stories model.Stories
	stories.Get()
	c.JSON(http.StatusOK, gin.H{"data": stories})
}

// GetStory : endpoint that returns a story by ID
func GetStory(c *gin.Context) {
	var story model.Story
	id := c.Params.ByName("id")
	story.Get(id)
	c.JSON(http.StatusOK, gin.H{"data": story})
}

// DeleteStory : endpoint that deletes a story by ID
func DeleteStory(c *gin.Context) {
	id := c.Params.ByName("id")
	message := "Story with id " + id + " was deleted."
	c.String(http.StatusOK, message)
}

// PostStory : endpoint that creates a story
func PostStory(c *gin.Context) {
	var story model.Story
	err := c.ShouldBindWith(&story, binding.JSON)
	if err != nil {
		err := story.Insert()

		if err != nil {
			message := "Story created"
			c.String(http.StatusOK, message)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	}

}

// PatchStory : endpoint that modify a story
func PatchStory(c *gin.Context) {
	message := "Story modified"
	c.String(http.StatusOK, message)
}
