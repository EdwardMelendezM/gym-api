package database

import (
	"context"
	"log"

	"gym-api/internal/ent"

	_ "github.com/lib/pq"
)

func NewEntClient(dsn string) *ent.Client {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

	return client
}
