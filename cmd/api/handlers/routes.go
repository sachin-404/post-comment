package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sachin-404/post-comment/cmd/api/handlers/post"
	"github.com/sachin-404/post-comment/cmd/api/handlers/user"
	"github.com/sachin-404/post-comment/config"
	"github.com/sachin-404/post-comment/internal/repo"
	"github.com/sachin-404/post-comment/internal/service/post_service"
	"github.com/sachin-404/post-comment/internal/service/user_service"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {
	userRepo := repo.NewUserRepo()
	userService := user_service.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	postRepo := repo.NewPostRepo()
	postService := post_service.NewPostService(postRepo)
	postHandler := post.NewPostHandler(postService)

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

	// JWT Middleware
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(config.GetApiConfig().JwtSecretKey),
		TokenLookup: "header:Authorization",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(user_service.JWTClaims)
		},
	})

	// Post routes
	postGroup := api.Group("/post", jwtMiddleware)
	postGroup.POST("", postHandler.CreatePost)
	postGroup.GET("/:id", postHandler.GetPost)
	postGroup.DELETE("/:id", postHandler.DeletePost)
}
