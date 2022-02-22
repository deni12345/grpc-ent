package main

import (
	"context"
	"log"

	pb "grpc-ent/gen/car"

	"google.golang.org/grpc"
)

func main() {
	server := ":8083"
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCarServiceClient(conn)
	ctx := context.Background()

	// t := time.Now().In(time.UTC)
	// createAt, _ := ptypes.TimestampProto(t)
	createRequest := &pb.GetCarInfoRequest{
		Id: 1,
	}
	res, err := c.GetCarInfo(ctx, createRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)
}
