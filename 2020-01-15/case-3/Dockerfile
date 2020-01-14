FROM golang:1.13

COPY project /app

RUN curl -sSfL https://bit.ly/-chicken- \
  | sh -s -- -b $(go env GOPATH)/bin v0.0.8
