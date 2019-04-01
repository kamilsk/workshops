package main

import (
	"context"
	"flag"
	"log"
	"os"
	"runtime"
	"text/template"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	helper "github.com/kamilsk/platform/pkg/protobuf"
	"github.com/kamilsk/platform/pkg/unsafe"
	"google.golang.org/grpc"

	"workshops/service-declarative-definition/internal/protocol/grpc/protobuf"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)

var (
	host   = flag.String("host", ":9001", "host to connect")
	status = `
Server Status:
  Increment:
    ID:        {{ .Increment.Id }}
    Value:     {{ .Increment.Value }}
    Update At: {{ format .Increment.UpdatedAt }}

  Fibonacci:
    ID:        {{ .Fibonacci.Id }}
    Value:     {{ .Fibonacci.Value }}
    Update At: {{ format .Fibonacci.UpdatedAt }}

  Unique String:
    ID:        {{ .Unique.Id }}
    Value:     {{ .Unique.Value }}
    Update At: {{ format .Unique.UpdatedAt }}
`
)

func main() {
	flag.Parse()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(*host, opts...)
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
	log.Printf("Version %s (commit: %s, build date: %s, go version: %s, compiler: %s, platform: %s/%s)\n",
		version, commit, date, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)

	spew.Fdump(os.Stdout, resp)
	tpl := template.New("status").Funcs(template.FuncMap{
		"format": func(ts *timestamp.Timestamp) string {
			return helper.Time(ts).Format(time.RFC822)
		},
	})
	unsafe.Ignore(template.Must(tpl.Parse(status)).Execute(os.Stderr, resp))
}
