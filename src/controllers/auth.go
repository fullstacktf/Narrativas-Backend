package controllers

import (
	"net/http"

	m "github.com/fullstacktf/Narrativas-Backend/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Register(context *gin.Context) {
	var newUser m.User

	if err := context.ShouldBindWith(&newUser, binding.JSON); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err := newUser.Register()

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		context.Status(http.StatusCreated)
	}
}

func Login(context *gin.Context) {
	var userData m.User

	if err := context.ShouldBindWith(&userData, binding.JSON); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid json provided"})
		return
	}

	token, err := userData.Login()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"token": token})
	}
}
