package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sachin-404/post-comment/common/database"
	"github.com/sachin-404/post-comment/config"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	cfg, err := config.LoadApiConfig()
	if err != nil {
		fmt.Println("failed to load config" + err.Error())
		log.Fatalf("Failed to load configuration: %v", err.Error())
	}
	if cfg == nil {
		log.Fatalf("Failed to load configuration")
		os.Exit(1)
	}

	// Initialize database
	database.Init(cfg.DatabaseDSN)

	// Create a new Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Setup routes
	//handlers.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
