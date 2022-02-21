package main

import (
	"context"
	"log"
	"net/http"

	pb "grpc-ent/gen/car"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcServerEndpoint string = "localhost:8083"

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterCarServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	log.Println("http server listen and serve at port :8081")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
