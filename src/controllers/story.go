package controllers

import (
	"net/http"
	"strconv"

	m "github.com/fullstacktf/Narrativas-Backend/models"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var stories m.Stories
	userid, _ := c.Get("user_id")

	stories.Get(userid.(uint))
	c.JSON(http.StatusOK, stories)
}

func GetStory(c *gin.Context) {
	var story m.Story
	userid, _ := c.Get("user_id")
	story.UserID = userid.(uint)

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
	var story m.Story
	userid, _ := c.Get("user_id")
	story.ID = uint(id)

	if err = story.Delete(userid.(uint)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}

func PostStory(c *gin.Context) {
	var story m.Story
	userid, _ := c.Get("user_id")
	story.UserID = userid.(uint)

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
	var event m.Event

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
	var relation m.EventRelation

	if err := c.BindJSON(&relation); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := relation.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, relation)
}
