package controllers

import (
	"net/http"
	"strconv"

	model "github.com/fullstacktf/Narrativas-Backend/api/models"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var stories model.Stories
	// userid, _ := c.Get("user_id")

	//stories.Get(userid.(uint))
	stories.Get(1)
	c.JSON(http.StatusOK, gin.H{"stories": stories})
}

func GetStory(c *gin.Context) {
	var story model.Story
	// userid, _ := c.Get("user_id")
	// story.UserID = userid.(uint)
	story.UserID = 1
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	story.Get(uint(id))

	c.JSON(http.StatusOK, gin.H{"story": story})
}

func DeleteStory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	var story model.Story
	// c.Get("user_id")
	story.ID = uint(id)
	userid := 1

	if err = story.Delete(uint(userid)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}

func PostStory(c *gin.Context) {
	var story model.Story
	// userid, _ := c.Get("user_id")
	// story.UserID = userid.(uint)
	story.UserID = 1

	if err := c.BindJSON(&story); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := story.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"story": story})
}

func PostEvent(c *gin.Context) {
	var event model.Event

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	if err := c.BindJSON(&event); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.StoryID = uint(id)

	if err := event.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func PostEventRelation(c *gin.Context) {
	var relation model.EventRelation

	if err := c.BindJSON(&relation); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := relation.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event_relation": relation})
}
