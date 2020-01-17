FROM golang:1.13

COPY project /app

RUN go get github.com/myitcv/gobin \
 && gobin github.com/izumin5210/gex/cmd/gex@v0.6.0 \
 && gobin github.com/twitchtv/retool@v1.3.7 \
 && curl -sSfL https://bit.ly/-chicken- \
  | sh -s -- -b $(go env GOPATH)/bin v0.0.8

WORKDIR /app
