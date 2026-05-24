package server

import (
	"mogi/internal/interfaces/handler"

	"github.com/labstack/echo/v4"
)

// setupRoutes configures all application routes
func setupRoutes(server *echo.Echo, usersHandler *handler.UserHandler) {
	// Health check / Welcome
	server.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Welcome",
		})
	})

	// User routes
	server.GET("/users", usersHandler.GetAll)
	server.GET("/users/:id", usersHandler.GetByID)
	server.POST("/users", usersHandler.Create)
	server.PUT("/users/:id", usersHandler.Update)
	server.DELETE("/users/:id", usersHandler.Delete)
}
