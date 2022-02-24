package main

import (
	"context"
	"io"
	"log"

	pb "grpc-ent/gen/car"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	server := ":8083"
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// c := pb.NewCarServiceClient(conn)
	ctx := context.Background()

	// t := time.Now().In(time.UTC)
	// createAt, _ := ptypes.TimestampProto(t)
	// createRequest := &pb.GetCarInfoRequest{
	// 	Id: 1,
	// }
	// res, err := c.GetCarInfo(ctx, createRequest)
	// if err != nil {
	// 	log.Fatalf("Create failed: %v", err)
	// }
	// log.Printf("Create result: <%+v>\n\n", res)

	c := pb.NewCarServiceClient(conn)
	in := emptypb.Empty{}
	stream, err := c.GetCarsInfo(ctx, &in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Printf("fail to stream cars: ", err)
			}
			log.Printf("Resp received: %s", resp.Model)
		}
	}()
	<-done
	log.Println("finished streaming")
}
