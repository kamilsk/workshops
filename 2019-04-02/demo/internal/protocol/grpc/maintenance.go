package grpc

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kamilsk/platform/pkg/pointer"
	helper "github.com/kamilsk/platform/pkg/protobuf"

	"workshops/service-declarative-definition/internal/protocol/grpc/protobuf"
)

func NewMaintenanceServer() protobuf.MaintenanceServer {
	return &maintenance{}
}

type maintenance struct{}

func (*maintenance) Status(context.Context, *empty.Empty) (*protobuf.StatusResponse, error) {
	log.Println("status called")
	return &protobuf.StatusResponse{
		Increment: &protobuf.StatusResponse_IncrementStatus{
			Id: 1, Value: 1,
			UpdatedAt: helper.Timestamp(pointer.ToTime(time.Now())),
		},
		Fibonacci: &protobuf.StatusResponse_FibonacciStatus{
			Id: 1, Value: 1,
			UpdatedAt: helper.Timestamp(pointer.ToTime(time.Now())),
		},
		Unique: &protobuf.StatusResponse_UniqueStringStatus{
			Id: 1, Value: "",
			UpdatedAt: helper.Timestamp(pointer.ToTime(time.Now())),
		},
	}, nil
}
