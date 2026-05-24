package main

import (
	"fmt"
	"mogi/internal/infrastructure/postgres"
	"mogi/internal/interfaces/server"
	"mogi/internal/shared/config"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := postgres.NewDatabase(cfg.Database.GetDSN())
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		panic(err)
	}

	// Initialize server
	srv := server.NewServer(db)

	// Start server
	address := ":" + cfg.Server.Port
	fmt.Println("Server starting on port", cfg.Server.Port)
	if err := srv.Start(address); err != nil {
		panic(err)
	}
}

