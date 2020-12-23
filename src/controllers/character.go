package controllers

import (
	"net/http"
	"strconv"

	m "github.com/fullstacktf/Narrativas-Backend/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetCharacter(c *gin.Context) {
	var character m.Character

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	character.ID = uint(id)
	userid, _ := c.Get("user_id")

	err = character.Get(userid.(uint))

	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.JSON(http.StatusOK, character)
}

func GetCharacters(c *gin.Context) {
	var characters m.Characters

	userid, _ := c.Get("user_id")

	if err := characters.Get(userid.(uint)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	c.JSON(http.StatusOK, characters)
}

func DeleteCharacter(c *gin.Context) {
	var character m.Character

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	character.ID = uint(id)

	userid, _ := c.Get("user_id")

	if err := character.Delete(userid.(uint)); err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Status(http.StatusOK)
}

func PostCharacter(c *gin.Context) {
	var character m.Character

	if err := c.ShouldBindWith(&character, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	userid, _ := c.Get("user_id")
	character.UserID = userid.(uint)

	err := character.Insert()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": character.ID})
}

func PutCharacter(c *gin.Context) {
	var character m.Character

	if err := c.ShouldBindWith(&character, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	userid, _ := c.Get("user_id")
	character.UserID = userid.(uint)

	err := character.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, character)
	}
}

func PostSection(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var section m.CharacterSection

	if err := c.ShouldBindWith(&section, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	section.CharacterID = uint(id)

	err = section.Insert()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, section)
}

func PostSectionField(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id2"), 10, 64)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var section m.CharacterSectionField

	if err := c.ShouldBindWith(&section, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	section.SectionID = uint(id)

	err = section.Insert()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, section)
}
