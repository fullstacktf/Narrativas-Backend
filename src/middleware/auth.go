package middleware

import (
	"net/http"

	"github.com/fullstacktf/Narrativas-Backend/common"
	"github.com/gin-gonic/gin"
)

func IsSignedIn(c *gin.Context) {
	token := c.Request.Header["Token"]

	if len(token) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
	}

	for _, n := range common.ActiveTokens {
		if token[0] == n.Token {
			c.Set("user_id", n.ID)
			return
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
