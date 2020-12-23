package router

import (
	"time"

	"github.com/fullstacktf/Narrativas-Backend/controllers"
	mw "github.com/fullstacktf/Narrativas-Backend/middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Token", "Content-Length", "Accept", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}))

	router.GET("/characters/", mw.IsSignedIn, controllers.GetCharacters)
	router.GET("/characters/:id", mw.IsSignedIn, controllers.GetCharacter)
	router.POST("/characters/", mw.IsSignedIn, controllers.PostCharacter)
	router.DELETE("/characters/:id", mw.IsSignedIn, controllers.DeleteCharacter)
	router.PUT("/characters/", mw.IsSignedIn, controllers.PutCharacter)
	router.POST("/characters/:id/sections", mw.IsSignedIn, controllers.PostSection)
	router.POST("/characters/:id/sections/:id2/fields", mw.IsSignedIn, controllers.PostSectionField)

	router.GET("/stories/", mw.IsSignedIn, controllers.Get)
	router.GET("/stories/:id", mw.IsSignedIn, controllers.GetStory)
	router.POST("/stories/", mw.IsSignedIn, controllers.PostStory)
	router.POST("/stories/:id/events/", mw.IsSignedIn, controllers.PostEvent)
	router.POST("/stories/:id/events/relations", mw.IsSignedIn, controllers.PostEventRelation)
	router.DELETE("/stories/:id", mw.IsSignedIn, controllers.DeleteStory)

	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)

	router.PUT("/upload/images/character", mw.IsSignedIn, controllers.UploadCharacter)
	router.PUT("/upload/images/story", mw.IsSignedIn, controllers.UploadStory)

	router.Static("/static", "./images")

	return router
}
