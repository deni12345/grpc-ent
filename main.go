package main

import (
	"context"
	"grpc-ent/api"
	"grpc-ent/db"
	"log"
)

func main() {
	ctx := context.Background()
	db, err := db.Open()
	defer db.Close()

	if err != nil {
		log.Fatalf("failed to open connection db :%v", err)
	}

	s := api.NewRouteServer(db)
	api.RunServer(ctx, s, "8083")

}
