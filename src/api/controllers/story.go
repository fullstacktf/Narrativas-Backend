package controllers

import (
	"net/http"
	"strconv"

	model "github.com/fullstacktf/Narrativas-Backend/api/models"
	"github.com/gin-gonic/gin"
)

// Get : returns all the stories
func Get(c *gin.Context) {
	var stories model.Stories
	userid, _ := c.Get("user_id")

	stories.Get(userid.(uint))
	c.JSON(http.StatusOK, gin.H{"data": stories})
}

// GetStory : endpoint that returns a story by ID
func GetStory(c *gin.Context) {
	var story model.Story
	userid, _ := c.Get("user_id")
	story.UserID = userid.(uint)

	id := c.Params.ByName("id")
	story.Get(id)
	c.JSON(http.StatusOK, gin.H{"data": story})
}

// DeleteStory : endpoint that deletes a story by ID
func DeleteStory(c *gin.Context) {
	id := c.Params.ByName("id")
	var story model.Story
	c.Get("user_id")
	// storyID, _ := strconv.Atoi(id)

	err := story.Get(id)
	if err == nil {
		err := story.Delete(id)
		if err == nil {
			message := "Story with id " + id + " was deleted."
			c.JSON(http.StatusOK, gin.H{"story": story})
			c.String(http.StatusOK, message)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		message := "Error: Story don't exist"
		c.String(http.StatusOK, message)
	}
}

// PostStory : endpoint that creates a story
func PostStory(c *gin.Context) {
	var story model.Story
	userid, _ := c.Get("user_id")
	story.UserID = userid.(uint)

	err := c.BindJSON(&story)
	if err == nil {
		err := story.Insert()
		if err == nil {
			message := "Story created"
			c.JSON(http.StatusOK, gin.H{"story": story})
			c.String(http.StatusOK, message)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// PostEvent : Add a new event
func PostEvent(c *gin.Context) {
	c.Get("user_id")

	id := c.Params.ByName("id")
	var event model.Event

	storyID, err := strconv.Atoi(id)
	event.StoryID = storyID

	err = c.BindJSON(&event)
	if err == nil {
		err := event.InsertEvent()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"event": event})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// PatchStory : endpoint that modify a story
func PatchStory(c *gin.Context) {

	id := c.Params.ByName("id")
	var story model.Story
	userid, _ := c.Get("user_id")
	story.UserID = userid.(uint)
	storyID, err := strconv.Atoi(id)
	story.ID = storyID

	err = c.BindJSON(&story)
	if err == nil {
		err := story.Update()
		if err == nil {
			message := "Story with " + id + " was modified"
			c.JSON(http.StatusOK, gin.H{"story": story})
			c.String(http.StatusOK, message)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
