FROM golang:1.13

COPY project /app

WORKDIR /app

RUN apt-get update \
 && apt-get install -y \
    tree \
 && go build -o /go/bin/server ./cmd/server.go
