package main

import (
	character "github.com/fullstacktf/Narrativas-Backend/api/character"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/character/:id", character.GetCharacter)
	r.POST("/character/", character.PostCharacter)
	r.DELETE("/character/:id", character.DeleteCharacter)
	r.PATCH("/character/:id", character.PatchCharacter)

	r.Run(":10000")
}
