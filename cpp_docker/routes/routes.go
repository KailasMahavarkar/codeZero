package routes

import (
	"cpp/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(e *echo.Echo, config *config.Config) {
	e.GET("/ping", ping)
	e.POST("/api/v1/code_processor", codeProcessor)
}

func codeProcessor(c echo.Context) error {
	response := map[string]string{"message": "Code Processor"}
	return c.JSON(http.StatusOK, response)
}

func ping(c echo.Context) error {
	response := map[string]string{"message": "pong"}
	return c.JSON(http.StatusOK, response)
}
