package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qrcode-login/auth"
	"qrcode-login/database"
	"qrcode-login/models"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func UserInfo(context *gin.Context) {
	var user models.User

	claims := auth.DecodeToken(context.GetHeader("Authorization"))
	record := database.Instance.Where("email = ?", claims.(*auth.JWTClaim).Email).First(&user)

	if record.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Can't find user data"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
	//context.JSON(http.StatusOK, user)

}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User or password not found"})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	refreshToken, err := auth.GenerateRefreshToken()

	context.JSON(http.StatusOK, TokenResponse{
		Token:        tokenString,
		RefreshToken: refreshToken,
	})

}
