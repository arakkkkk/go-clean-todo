package app

import (
	"context"
	"log"
	"todo/pkg/ent"
)

func init() {

	client, err := ent.Open()
	if err != nil {
		panic(err)
	}

	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
