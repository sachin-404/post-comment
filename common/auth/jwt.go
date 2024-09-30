package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sachin-404/post-comment/config"
	"time"
)

type JWTClaims struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(name string, userID int) (string, error) {
	claims := JWTClaims{
		Name:   name,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKey := config.GetApiConfig().JwtSecretKey
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUserIDFromContext(c echo.Context) (int, error) {
	user := c.Get("user")
	if user == nil {
		return 0, errors.New("no user found in context")
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	return claims.UserID, nil
}
