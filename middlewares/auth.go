package middlewares

import (
	"net/http"
	"strings"

	"bojana.dev/api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authorizationHeader := context.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	token := strings.Split(authorizationHeader, " ")[1]
	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
