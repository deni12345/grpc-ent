package main

import (
	"context"
	"log"
	"time"

	pb "grpc-ent/gen/car"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

func main() {
	server := "localhost:8083"
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCarServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	t := time.Now().In(time.UTC)
	createAt, _ := ptypes.TimestampProto(t)
	createRequest := &pb.CreateCarInfoRequest{
		Car: &pb.Car{
			Model:      "toyota",
			CreateTime: createAt,
			Cost:       1500000,
		},
	}
	res, err := c.CreateCarInfo(ctx, createRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)
}
