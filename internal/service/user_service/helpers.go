package user_service

import (
	"errors"
	"github.com/sachin-404/post-comment/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetApiConfig().JwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
