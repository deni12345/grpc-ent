package api

import (
	"context"
	pb "grpc-ent/gen/car"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *routeServer) CreateCarInfo(ctx context.Context, params *pb.CreateCarInfoRequest) (*pb.CreateCarInfoResponse, error) {
	car, err := s.DB.Car.
		Create().
		SetModel(params.Car.Model).
		SetCost(float64(params.Car.Cost)).
		SetManufactureDate(params.Car.CreateTime.AsTime()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCarInfoResponse{
		Car: &pb.Car{
			Id:         int64(car.ID),
			Model:      car.Model,
			CreateTime: timestamppb.New(car.ManufactureDate),
			Cost:       float32(car.Cost),
		},
	}, nil
}
