package user_service

import (
	"github.com/sachin-404/post-comment/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	jwtSecretKey := config.GetApiConfig().JwtSecretKey
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//func ValidateToken(tokenString string) (int, error) {
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//
//	if err != nil {
//		return 0, err
//	}
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		userID := int(claims["user_id"].(float64))
//		return userID, nil
//	}
//
//	return 0, jwt.ErrSignatureInvalid
//}
