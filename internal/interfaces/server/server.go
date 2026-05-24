package server

import (
	"mogi/internal/infrastructure/postgres"
	"mogi/internal/interfaces/handler"
	"mogi/internal/usecase/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewServer initializes and returns Echo server
func NewServer(db *postgres.Database) *echo.Echo {
	server := echo.New()

	// Middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// Initialize repositories and usecases
	usersRepo := postgres.NewUserRepository(db)
	usersUC := users.NewUseCase(usersRepo)

	// Initialize handlers
	usersHandler := handler.NewUserHandler(usersUC)

	// Setup routes
	setupRoutes(server, usersHandler)

	return server
}
