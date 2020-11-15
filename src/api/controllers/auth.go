package controllers

import (
	"net/http"

	models "github.com/fullstacktf/Narrativas-Backend/api/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Register : Endpoint that allows user register
func Register(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request."})
		return
	}

	if newUser.Insert() {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Username or email already exists."})
	} else {
		c.JSON(http.StatusOK, gin.H{"Message": "User created."})
	}

	return
}
