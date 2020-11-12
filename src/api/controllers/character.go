package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCharacter : endpoint that returns a character by ID
func GetCharacter(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// DeleteCharacter : endpoint that deletes a character by ID
func DeleteCharacter(c *gin.Context) {
	id := c.Params.ByName("id")
	message := "Character with id " + id + " was deleted."
	c.String(http.StatusOK, message)
}

// PostCharacter : endpoint that creates a character
func PostCharacter(c *gin.Context) {
	message := "Character created"
	c.String(http.StatusOK, message)
}

// PatchCharacter : endpoint that updates a character
func PatchCharacter(c *gin.Context) {
	id := c.Params.ByName("id")
	message := "Character with id " + id + " was updated."
	c.String(http.StatusOK, message)
}
