package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/fullstacktf/Narrativas-Backend/common"

	"github.com/gin-gonic/gin"
)

// curl localhost:9090/upload/images/character -X PUT -H "Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDc5NTA3NTYsInVzZXJfaWQiOjR9.0kjZFKUI6nCYEh36ehBCQCd4JnkYN4yRKjSyMyCOAuU" --form image=@Yo.jpg

func UploadImage(c *gin.Context, path string) {
	fmt.Println("Upload image")
	form, _ := c.MultipartForm()
	files := form.File["image"]
	file, _ := c.FormFile("image")
	fileExtension := filepath.Ext(file.Filename)
	allowedExtensions := []string{".jpg", ".png"}

	if !common.StringInSlice(fileExtension, allowedExtensions) {
		c.Status(http.StatusUnsupportedMediaType)
		return
	}

	uuid, err := common.GenerateUUID()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	filename := uuid + fileExtension
	for _, file := range files {
		err := c.SaveUploadedFile(file, "./"+path+filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"image": filename})
}

func UploadCharacter(c *gin.Context) {
	UploadImage(c, "images/characters/")
}

func UploadStory(c *gin.Context) {
	UploadImage(c, "images/stories/")
}
