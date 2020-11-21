package controllers

import (
	"net/http"

	m "github.com/fullstacktf/Narrativas-Backend/api/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Register : Endpoint that allows user register
func Register(c *gin.Context) {
	var newUser m.User

	if err := c.ShouldBindWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	_, err := newUser.Register()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "user created succesfully"})
	}
}

// Login : endpoint that allows user log in
func Login(c *gin.Context) {
	var userData m.User

	if err := c.ShouldBindWith(&userData, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid json provided"})
		return
	}

	token, err := userData.Login()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
