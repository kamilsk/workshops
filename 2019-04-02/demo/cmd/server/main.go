package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kamilsk/platform/pkg/safe"

	"workshops/service-declarative-definition/internal/protocol/grpc"
)

var (
	port = flag.Uint("port", 9001, "tcp-port for listening")
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	defer safe.Close(listener, func(err error) { log.Fatal(err) })

	server, err := grpc.NewServer(
		grpc.WithMaintenanceServer(grpc.NewMaintenanceServer()),
		grpc.WithSequenceServer(grpc.NewSequenceServer()),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start server at", *port)
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
