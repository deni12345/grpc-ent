package api

import (
	"context"
	pb "grpc-ent/gen/car"
	"log"

	"github.com/golang/protobuf/ptypes"
)

func (s *routeServer) CreateCarInfo(ctx context.Context, params *pb.CreateCarInfoRequest) (*pb.CreateCarInfoResponse, error) {
	log.Println(params.Car.CreateTime)
	cost := float64(params.Car.Cost)
	manufactureDate := params.Car.CreateTime.AsTime()

	if err := s.DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resource: %v", err)
	}

	car, err := s.DB.Car.
		Create().
		SetModel(params.Car.Model).
		SetCost(cost).
		SetManufactureDate(manufactureDate).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(car.ManufactureDate)

	if err != nil {
		log.Printf("failed to convert manufactureDate: %v", err)
		return nil, err
	}

	return &pb.CreateCarInfoResponse{
		Car: &pb.Car{
			Id:         int64(car.ID),
			Model:      car.Model,
			CreateTime: createTime,
			Cost:       float32(car.Cost),
		},
	}, nil
}
