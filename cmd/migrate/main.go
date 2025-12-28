package main

import (
	"context"
	"log"
	"os"

	"gym-api/internal/database"
)

func main() {
	if os.Getenv("APP_ENV") != "local" {
		log.Println("Migration is local")
	} else {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("DATABASE_URL environment variable is not set")
		}
		client := database.NewEntClient(dsn)
		defer client.Close()

		log.Println("Running database migrations...")

		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("migration failed: %v", err)
		}

		log.Println("Migrations finished successfully")
	}
}
