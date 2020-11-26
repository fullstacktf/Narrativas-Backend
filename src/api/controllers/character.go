package controllers

import (
	"net/http"
	"strconv"

	models "github.com/fullstacktf/Narrativas-Backend/api/models"
	common "github.com/fullstacktf/Narrativas-Backend/common"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetCharacter(c *gin.Context) {
	var character models.Character

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	character.ID = uint(id)

	token := c.Request.Header["Token"]

	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}

	userid, err := common.IsSignedIn(token[0])

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	err = character.Get(uint(userid))

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, character)
}

func GetCharacters(c *gin.Context) {
	var characters models.Characters

	token := c.Request.Header["Token"]

	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}

	id, err := common.IsSignedIn(token[0])

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if err := characters.Get(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"characters": characters})
}

func DeleteCharacter(c *gin.Context) {
	var character models.Character

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	character.ID = uint(id)
	token := c.Request.Header["Token"]

	if len(token) == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	userid, err := common.IsSignedIn(token[0])

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if err := character.Delete(userid); err != nil {
		c.Status(http.StatusForbidden)
		return
	}

	c.Status(http.StatusOK)
}

func PostCharacter(c *gin.Context) {
	var character models.Character

	if err := c.ShouldBindWith(&character, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	token := c.Request.Header["Token"]

	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
	}

	err := character.Insert(token[0])

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Character created succesfully"})
	}
}

func PutCharacter(c *gin.Context) {
	var character models.Character

	if err := c.ShouldBindWith(&character, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	token := c.Request.Header["Token"]

	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
	}

	err := character.Update(token[0])

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Character updated succesfully"})
	}
}
