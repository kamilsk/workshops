package grpc

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"

	"workshops/service-declarative-definition/internal/protocol/grpc/protobuf"
)

func NewSequenceServer() protobuf.SequenceServer {
	return &sequence{}
}

type sequence struct{}

func (*sequence) Increment(context.Context, *empty.Empty) (*protobuf.IncrementNumber, error) {
	log.Println("increment called")
	return &protobuf.IncrementNumber{Id: 1, Value: 1}, nil
}

func (*sequence) Fibonacci(context.Context, *empty.Empty) (*protobuf.FibonacciNumber, error) {
	log.Println("fibonacci called")
	return &protobuf.FibonacciNumber{Id: 1, Value: 1}, nil
}

func (*sequence) Unique(context.Context, *empty.Empty) (*protobuf.UniqueString, error) {
	log.Println("unique called")
	return &protobuf.UniqueString{Id: 1, Value: "unique"}, nil
}
