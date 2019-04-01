FROM golang:latest AS build

WORKDIR /src

ADD . .

ARG COMMIT
ARG DATE
ARG VERSION

ENV GOOS        linux
ENV GOARCH      amd64
ENV CGO_ENABLED 0
ENV GO111MODULE on

RUN go env \
 && export _commit="-X main.commit=${COMMIT}" \
 && export _date="-X main.date=${DATE}" \
 && export _version="-X main.version=${VERSION}" \
 && go build -ldflags "-s -w ${_commit} ${_date} ${_version}" \
             -mod vendor \
             -o /go/bin/service \
             ./cmd/server/main.go


FROM alpine:latest AS service

LABEL maintainer="Kamil Samigullin <kamil@samigullin.info>"

RUN adduser -D -H -u 1000 service
USER service

COPY --from=build /go/bin/service /usr/local/bin/

ENV BIND=0.0.0.0 \
    HTTP_PORT=8080 \
    GRPC_PORT=9001

EXPOSE ${HTTP_PORT} \
       ${GRPC_PORT}

ENTRYPOINT [ "service" ]
CMD        [ "-h" ]
