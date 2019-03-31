package gateway

import (
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kamilsk/platform/pkg/safe"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func Must(target string, options ...Option) *server {
	srv, err := New(target, options...)
	if err != nil {
		panic(err)
	}
	return srv
}

func New(target string, options ...Option) (*server, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "gateway server: connect to %s", target)
	}
	srv := &server{conn, runtime.NewServeMux()}
	for i, configure := range options {
		if err := configure(srv); err != nil {
			return nil, errors.Wrapf(err, "gateway server: invalid %d option", i)
		}
	}
	return srv, nil
}

type server struct {
	conn *grpc.ClientConn
	mux  *runtime.ServeMux
}

func (server *server) Serve(listener net.Listener) error {
	defer safe.Close(listener, func(err error) { log.Fatal(err) })
	return (&http.Server{Handler: server.mux}).Serve(listener)
}
