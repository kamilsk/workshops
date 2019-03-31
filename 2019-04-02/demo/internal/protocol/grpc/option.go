package grpc

import (
	"google.golang.org/grpc"

	"workshops/service-declarative-definition/internal/protocol/grpc/protobuf"
)

type Option func(*grpc.Server) error

func WithMaintenanceServer(srv protobuf.MaintenanceServer) Option {
	return func(server *grpc.Server) error {
		protobuf.RegisterMaintenanceServer(server, srv)
		return nil
	}
}

func WithSequenceServer(srv protobuf.SequenceServer) Option {
	return func(server *grpc.Server) error {
		protobuf.RegisterSequenceServer(server, srv)
		return nil
	}
}
