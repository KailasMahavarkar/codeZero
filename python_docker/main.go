package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"python/config"
	"python/routes"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Global variable to hold the configuration
var AppConfig *config.Config

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	e := echo.New()

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading environment variables")
	}

	log.SetOutput(os.Stderr)
	// Apply rate limiter middleware
	rateLimiterConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	// Apply rate limiter middleware
	e.Use(middleware.RateLimiterWithConfig(rateLimiterConfig))

	// Apply CORS middleware
	e.Use(middleware.CORS())

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	// Assign the configuration to the global variable
	AppConfig = config

	// // spawn a goroutine to clear the cache every 5 minutes
	// go func() {
	// 	for {
	// 		time.Sleep(5 * time.Minute)
	// 	}
	// }()

	// Register routes
	routes.RegisterRoutes(e, AppConfig)

	// Start the server
	e.Start(getPort())
	log.Println("Server Started!!!")
}
