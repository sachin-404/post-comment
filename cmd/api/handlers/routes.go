package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sachin-404/post-comment/cmd/api/handlers/user"
	"github.com/sachin-404/post-comment/internal/repo"
	"github.com/sachin-404/post-comment/internal/service/user_service"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {
	userRepo := repo.NewUserRepo()
	userService := user_service.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	api := e.Group("/api")

	// Health check
	api.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// User routes
	userGroup := api.Group("/user")
	userGroup.POST("/signup", userHandler.SignUp)
	userGroup.POST("/login", userHandler.Login)
}
