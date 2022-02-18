package api

import (
	"context"
	"grpc-ent/ent"
	pb "grpc-ent/gen/car"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

type routeServer struct {
	pb.UnimplementedCarServiceServer
	DB *ent.Client
}

func NewRouteServer(db *ent.Client) *routeServer {
	return &routeServer{
		DB: db,
	}
}
func RunServer(ctx context.Context, s *routeServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	pb.RegisterCarServiceServer(server, s)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
