package api

import (
	"context"
	"grpc-ent/ent"
	"grpc-ent/ent/car"
	pb "grpc-ent/gen/car"
	"log"

	"github.com/golang/protobuf/ptypes"
)

func (s *routeServer) GetCarInfo(ctx context.Context, params *pb.GetCarInfoRequest) (*pb.GetCarInfoResponse, error) {
	var carInfo *ent.Car = nil
	if params.Model != nil {
		carInfo = s.DB.Car.Query().Where(car.ID(int(params.Id)), car.Model(*params.Model)).OnlyX(ctx)
	} else {
		carInfo = s.DB.Car.Query().Where(car.ID(int(params.Id))).OnlyX(ctx)
	}
	createTime, err := ptypes.TimestampProto(carInfo.ManufactureDate)

	if err != nil {
		log.Printf("failed to convert manufactureDate: %v", err)
		return nil, err
	}

	return &pb.GetCarInfoResponse{
		Car: &pb.Car{
			Model:      carInfo.Model,
			CreateTime: createTime,
		},
	}, nil
}
