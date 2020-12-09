package controllers

import (
	"net/http"

	m "github.com/fullstacktf/Narrativas-Backend/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Register(c *gin.Context) {
	var newUser m.User

	if err := c.ShouldBindWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err := newUser.Register()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusCreated)
	}
}

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
