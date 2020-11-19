package controllers

import (
	"net/http"

	models "github.com/fullstacktf/Narrativas-Backend/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetCharacter : endpoint that returns a character by ID
func GetCharacter(c *gin.Context) {
	// var characters models.Character
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetCharacters : endpoint that returns a character by ID
func GetCharacters(c *gin.Context) {
	var characters models.Characters
	// TO DO: SET ID DEPENDING JWT TOKEN

	if err := characters.Get(1); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

// DeleteCharacter : endpoint that deletes a character by ID
func DeleteCharacter(c *gin.Context) {
	id := c.Params.ByName("id")
	message := "Character with id " + id + " was deleted."
	c.String(http.StatusOK, message)
}

// PostCharacter : endpoint that creates a character
func PostCharacter(c *gin.Context) {
	var newCharacter models.Character

	if err := c.ShouldBindWith(&newCharacter, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	_, err := newCharacter.Insert()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Character created succesfully"})
	}
}

// PatchCharacter : endpoint that updates a character
func PatchCharacter(c *gin.Context) {
	id := c.Params.ByName("id")
	message := "Character with id " + id + " was updated."
	c.String(http.StatusOK, message)
}
