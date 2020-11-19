package router

import (
	controllers "github.com/fullstacktf/Narrativas-Backend/api/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter : initialice router
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/characters/", controllers.GetCharacters)
	router.GET("/characters/:id", controllers.GetCharacter)
	router.POST("/characters/", controllers.PostCharacter)
	router.DELETE("/characters/:id", controllers.DeleteCharacter)
	router.PATCH("/characters/:id", controllers.PatchCharacter)

	router.GET("/stories/:id", controllers.GetStory)
	router.POST("/stories/", controllers.PostStory)
	router.DELETE("/stories/:id", controllers.DeleteStory)

	router.POST("/auth/register", controllers.Register)

	return router
}
