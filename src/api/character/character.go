package character

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCharacter(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func DeleteCharacter(c *gin.Context) {
	id := c.Params.ByName("id")
	message := "Character with id " + id + " was deleted."
	c.String(http.StatusOK, message)
}

func PostCharacter(c *gin.Context) {
	message := "Character created"
	c.String(http.StatusOK, message)
}

func PatchCharacter(c *gin.Context) {
	id := c.Params.ByName("id")
	message := "Character with id " + id + " was updated."
	c.String(http.StatusOK, message)
}
