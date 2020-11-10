package main

import (
	character "github.com/fullstacktf/Narrativas-Backend/api/character"
	story "github.com/fullstacktf/Narrativas-Backend/api/story"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/character/:id", character.GetCharacter)
	r.POST("/character/", character.PostCharacter)
	r.DELETE("/character/:id", character.DeleteCharacter)
	r.PATCH("/character/:id", character.PatchCharacter)

	r.GET("/story/:id", story.GetStory)

	r.Run(":10000")
}
