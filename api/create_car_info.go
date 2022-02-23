package api

import (
	"context"

	v1 "grpc-ent/gen/car"
	"log"

	"github.com/golang/protobuf/ptypes"
)

func (s *routeServer) CreateCarInfo(ctx context.Context, params *v1.CreateCarInfoRequest) (*v1.CreateCarInfoResponse, error) {
	log.Println(params)
	manufactureDate := params.CreateTime.AsTime()

	if err := s.DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resource: %v", err)
	}

	car, err := s.DB.Car.
		Create().
		SetModel(params.Model).
		SetCost(params.Cost).
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

	return &v1.CreateCarInfoResponse{
		Car: &v1.Car{
			Id:         int64(car.ID),
			Model:      car.Model,
			CreateTime: createTime,
			Cost:       car.Cost,
		},
	}, nil
}
