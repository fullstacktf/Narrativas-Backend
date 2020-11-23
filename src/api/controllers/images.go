package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/fullstacktf/Narrativas-Backend/common"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, path string) {

	form, _ := c.MultipartForm()
	files := form.File["file"]
	file, _ := c.FormFile("file")
	fileExtension := filepath.Ext(file.Filename)
	allowedExtensions := []string{".jpg", ".png"}

	if !common.StringInSlice(fileExtension, allowedExtensions) {
		c.Status(415)
		return
	}

	uuid, err := common.GenerateUUID()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for _, file := range files {
		err := c.SaveUploadedFile(file, "./"+path+uuid+fileExtension)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.Status(http.StatusOK)
}

func UploadCharacter(c *gin.Context) {
	UploadImage(c, "images/characters/")
}

func UploadStory(c *gin.Context) {
	UploadImage(c, "images/stories/")
}
