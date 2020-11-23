package controllers

import (
	"fmt"
	"net/http"

	"github.com/fullstacktf/Narrativas-Backend/common"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, path string) {

	form, _ := c.MultipartForm()
	files := form.File["file"]

	uuid, err := common.GenerateUUID()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for _, file := range files {
		err := c.SaveUploadedFile(file, "./"+path+uuid+".png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func UploadCharacter(c *gin.Context) {
	UploadImage(c, "images/characters/")
}

func UploadStory(c *gin.Context) {
	UploadImage(c, "images/stories/")
}
