package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"workshops/service-declarative-definition/internal/protocol/grpc/protobuf"
)

var (
	port = flag.Uint("port", 9001, "tcp-port for listening")
)

func main() {
	flag.Parse()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), opts...)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := protobuf.NewMaintenanceClient(conn)
	resp, err := client.Status(ctx, &empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(resp)
}
