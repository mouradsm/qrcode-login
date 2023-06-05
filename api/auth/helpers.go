package auth

import (
	"errors"
	"qrcode-login/database"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func GenerateRefreshToken() (tokenString string, err error) {

	//I'm returning a new JWT but this is not really necessary
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)

	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := parseToken(signedToken)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	user := database.Instance.Where("email = ?", claims.Email)
	if user == nil {
		err = errors.New("can't find user data")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}

func DecodeToken(signedToken string) (claims jwt.Claims) {
	token, err := parseToken(signedToken)

	if err != nil {
		err = errors.New("couldn't parse claims")
		return
	}

	return token.Claims.(*JWTClaim)
}

func parseToken(signedToken string) (token *jwt.Token, err error) {

	return jwt.ParseWithClaims(strings.Split(signedToken, "Bearer ")[1], &JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
}
