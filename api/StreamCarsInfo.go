package api

import (
	"context"
	"fmt"
	v1 "grpc-ent/gen/car"
	"log"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *routeServer) GetCarsInfo(params *emptypb.Empty, stream v1.CarService_GetCarsInfoServer) error {
	log.Println("hello from cars")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cars, err := s.DB.Car.Query().All(ctx)
	if err != nil {
		return fmt.Errorf("fail to query list of cars: ", err)
	}

	for _, car := range cars {
		createTime, err := ptypes.TimestampProto(car.ManufactureDate)
		if err != nil {
			return fmt.Errorf("failed to convert manufactureDate: %v", err)
		}

		if err := stream.Send(&v1.Car{
			Model:      car.Model,
			CreateTime: createTime,
			Cost:       car.Cost,
		}); err != nil {
			return fmt.Errorf("fail to send stream: %v", err)
		}
	}
	return nil
}
