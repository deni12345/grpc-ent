package api

import (
	"context"
	"grpc-ent/ent"
	pb "grpc-ent/gen/car"
	"log"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *routeServer) GetCarsInfo(ctx context.Context, _ *emptypb.Empty) (*pb.GetCarsInfoResponse, error) {
	entCars, err := s.DB.Car.Query().All(ctx)
	if err != nil {
		log.Printf("failed to get all of car's infos %v", err)
		return nil, err
	}
	cars, err := ListCars(entCars)
	if err != nil {
		log.Printf("failed to convert from ent.car list to pb.car list %v", err)
		return nil, err
	}

	return &pb.GetCarsInfoResponse{
		Car: cars,
	}, nil
}

func ListCars(entCars []*ent.Car) ([]*pb.Car, error) {
	var cars []*pb.Car
	for _, car := range entCars {
		createTime, err := ptypes.TimestampProto(car.ManufactureDate)

		if err != nil {
			log.Printf("failed to convert manufactureDate: %v", err)
			return nil, err
		}

		carItem := &pb.Car{
			Id:         int64(car.ID),
			Model:      car.Model,
			CreateTime: createTime,
			Cost:       car.Cost,
		}
		cars = append(cars, carItem)
	}
	return cars, nil
}
