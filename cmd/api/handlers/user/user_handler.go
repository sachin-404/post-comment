package user

import (
	"github.com/sachin-404/post-comment/common/auth"
	"github.com/sachin-404/post-comment/internal/service/user_service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *user_service.UserService
}

func NewUserHandler(userService *user_service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUp(c echo.Context) error {
	var req user_service.UserSignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.userService.CreateUser(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func (h *UserHandler) Login(c echo.Context) error {
	var req user_service.UserLoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	user, err := h.userService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	token, err := auth.GenerateToken(user.Name, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
