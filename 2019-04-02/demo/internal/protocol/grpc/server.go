package grpc

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewServer(options ...Option) (*grpc.Server, error) {
	srv := grpc.NewServer()
	for i, configure := range options {
		if err := configure(srv); err != nil {
			return nil, errors.Wrapf(err, "grpc server: invalid %d option", i)
		}
	}
	return srv, nil
}
