package database

import (
	"log"

	"gym-api/internal/ent"

	_ "github.com/lib/pq"
)

func NewEntClient(dsn string) *ent.Client {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection: %v", err)
	}

	//if errCreated := client.Schema.Create(context.Background()); errCreated != nil {
	//	log.Fatalf("failed creating schema: %v", err)
	//}

	return client
}
