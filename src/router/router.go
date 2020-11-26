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

	router.GET("/story/", controllers.Get)
	router.GET("/story/:id", controllers.GetStory)
	router.POST("/story/", controllers.PostStory)
	router.DELETE("/story/:id", controllers.DeleteStory)
	router.PATCH("/story/:id", controllers.PatchStory)

	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)

	return router
}
