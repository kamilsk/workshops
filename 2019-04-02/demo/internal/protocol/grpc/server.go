package grpc

import (
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func Must(options ...Option) *grpc.Server {
	srv, err := New(options...)
	if err != nil {
		panic(err)
	}
	return srv
}

func New(options ...Option) (*grpc.Server, error) {
	srv := grpc.NewServer(
		middleware.WithUnaryServerChain(
			recovery.UnaryServerInterceptor(),
		),
	)
	for i, configure := range options {
		if err := configure(srv); err != nil {
			return nil, errors.Wrapf(err, "grpc server: invalid %d option", i)
		}
	}
	return srv, nil
}
