FROM golang:1.13

COPY project /app

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
  | sh -s -- -b $(go env GOPATH)/bin v1.22.2

WORKDIR /app
