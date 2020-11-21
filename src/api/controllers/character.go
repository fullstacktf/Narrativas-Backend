package controllers

import (
	"net/http"
	"strconv"

	models "github.com/fullstacktf/Narrativas-Backend/api/models"
	common "github.com/fullstacktf/Narrativas-Backend/common"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetCharacter : endpoint that returns a character by ID
func GetCharacter(c *gin.Context) {
	//	id := c.Params.ByName("id")

	var test = models.Test{
		Character: models.Character{
			ID:        1,
			UserID:    1,
			Name:      "Prueba",
			Biography: "123",
			Image:     "testtt.png",
			CharacterSection: []models.CharacterSection{
				{
					ID:          1,
					CharacterID: 1,
					Title:       "prueba",
					CharacterSectionField: []models.CharacterSectionField{
						{
							ID:          1,
							SectionID:   1,
							Name:        "Titulo 1",
							Value:       "255",
							Description: "prueba",
						},
						{
							ID:          2,
							SectionID:   1,
							Name:        "Titulo 2",
							Value:       "333",
							Description: "prueba",
						},
					},
				},
				{
					ID:          2,
					CharacterID: 1,
					Title:       "prueba2",
					CharacterSectionField: []models.CharacterSectionField{
						{
							ID:          3,
							SectionID:   1,
							Name:        "Titulo 1",
							Value:       "255",
							Description: "prueba",
						},
					},
				},
			},
		},
	}

	c.JSON(http.StatusOK, test)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := characters.Get(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

func DeleteCharacter(c *gin.Context) {
	var character models.Character

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := character.Delete(userid); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "character deleted succesfully"})
}

func PostCharacterAux(c *gin.Context) {
	var newCharacter models.Character

	if err := c.ShouldBindWith(&newCharacter, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	token := c.Request.Header["Token"]

	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
	}

	_, err := newCharacter.Insert(token[0])

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Character created succesfully"})
	}
}

func PostCharacter(c *gin.Context) {
	var newCharacter models.Test

	if err := c.ShouldBindWith(&newCharacter, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	c.JSON(http.StatusOK, newCharacter)

	// token := c.Request.Header["Token"]
	//
	// if len(token) == 0 {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
	// }
	//
	err := newCharacter.Insert()

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
