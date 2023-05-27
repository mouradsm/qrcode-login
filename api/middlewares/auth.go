package middlewares

import (
	"github.com/gin-gonic/gin"
	"qrcode-login/auth"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if (tokenString) == "" {
			context.JSON(401, gin.H{"error": "request does not contain a access token"})
			context.Abort()
			return
		}

		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
