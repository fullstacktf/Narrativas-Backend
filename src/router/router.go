package router

import (
	"github.com/fullstacktf/Narrativas-Backend/api/controllers"
	mw "github.com/fullstacktf/Narrativas-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/characters/", mw.IsSignedIn, controllers.GetCharacters)
	router.GET("/characters/:id", mw.IsSignedIn, controllers.GetCharacter)
	router.POST("/characters/", mw.IsSignedIn, controllers.PostCharacter)
	router.DELETE("/characters/:id", mw.IsSignedIn, controllers.DeleteCharacter)
	router.PUT("/characters/", mw.IsSignedIn, controllers.PutCharacter)

	router.GET("/story/", mw.IsSignedIn, controllers.Get)
	router.GET("/story/:id", mw.IsSignedIn, controllers.GetStory)
	router.POST("/story/", mw.IsSignedIn, controllers.PostStory)
	router.DELETE("/story/:id", mw.IsSignedIn, controllers.DeleteStory)
	router.PATCH("/story/:id", mw.IsSignedIn, controllers.PatchStory)

	// router.POST("/story/:id", mw.IsSignedIn, controllers.PostEvent)

	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)

	router.POST("/upload/images/character", mw.IsSignedIn, controllers.UploadCharacter)
	router.POST("/upload/images/story", mw.IsSignedIn, controllers.UploadStory)

	router.Static("/static", "./images")

	return router
}
