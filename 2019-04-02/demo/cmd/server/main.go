package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/kamilsk/platform/pkg/safe"

	"workshops/service-declarative-definition/internal/protocol/grpc"
	"workshops/service-declarative-definition/internal/protocol/grpc/gateway"
)

var (
	restPort = flag.Uint("rest", 8080, "tcp port for http listening")
	grpcPort = flag.Uint("grpc", 9001, "tcp port for grpc listening")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	restListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *restPort))
	if err != nil {
		log.Fatal(err)
	}
	defer safe.Close(restListener, func(err error) { log.Fatal(err) })

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	defer safe.Close(grpcListener, func(err error) { log.Fatal(err) })

	server := grpc.Must(
		grpc.WithMaintenanceServer(grpc.NewMaintenanceServer()),
		grpc.WithSequenceServer(grpc.NewSequenceServer()),
	)

	go func() {
		time.Sleep(time.Second)
		server := gateway.Must(fmt.Sprintf(":%d", *grpcPort),
			gateway.WithMaintenanceHandler(ctx),
			gateway.WithSequenceHandler(ctx),
		)
		log.Println("start rest gateway at", *restPort)
		if err := server.Serve(restListener); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("start grpc server at", *grpcPort)
	if err := server.Serve(grpcListener); err != nil {
		log.Fatal(err)
	}
}
