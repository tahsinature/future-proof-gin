package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tahsinature/future-proof-gin/pkg/services"
)

var authService = new(services.AuthService)

func Auth(c *gin.Context) {
	tokenAuth, err := authService.ExtractTokenMetadata(c.Request)
	if err != nil {
		// Token either expired or not valid
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}

	userID, err := authService.FetchAuth(tokenAuth)
	if err != nil {
		// Token does not exists in Redis (User logged out or expired)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}

	// To be called from GetUserID()
	c.Set("userID", userID)
	c.Next()
}
