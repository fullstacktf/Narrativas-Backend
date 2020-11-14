package router

import (
	controllers "github.com/fullstacktf/Narrativas-Backend/api/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter : initialice router
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/character/:id", controllers.GetCharacter)
	router.POST("/character/", controllers.PostCharacter)
	router.DELETE("/character/:id", controllers.DeleteCharacter)
	router.PATCH("/character/:id", controllers.PatchCharacter)

	router.GET("/story/:id", controllers.GetStory)
	router.POST("/story/", controllers.PostStory)
	router.DELETE("/story/:id", controllers.DeleteStory)

	return router
}
