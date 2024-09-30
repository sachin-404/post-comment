package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sachin-404/post-comment/internal/service/user_service"
)

func GetUserIDFromContext(c echo.Context) (int, error) {
	user := c.Get("user")
	if user == nil {
		return 0, errors.New("no user found in context")
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*user_service.JWTClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	return claims.UserID, nil
}
