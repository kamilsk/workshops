GO_PROTO      = v1.3.1
PROTO_GATEWAY = v1.8.5

COMMIT        = $(shell git rev-parse --short HEAD || echo 'none')
DATE          = $(shell date -u +%FT%X%Z || echo 'unknown')
VERSION       = $(shell git describe --tags 2>&- || echo 'dev' | cut -d - -f 1)

.PHONY: deps
deps: os-deps proto-deps go-deps

.PHONY: go-deps
go-deps:
	@cd demo && go mod tidy && go mod vendor && go mod verify

.PHONY: os-deps
os-deps:
	@brew bundle

.PHONY: proto-deps
proto-deps:
	@( \
	  export GO111MODULE=off; \
	  go get -d -u github.com/golang/protobuf/protoc-gen-go; \
	  git -C "$$(go env GOPATH)"/src/github.com/golang/protobuf checkout $(GO_PROTO); \
	  go install github.com/golang/protobuf/protoc-gen-go; \
	)
	@( \
	  export GO111MODULE=off; \
	  go get -d -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway; \
	  git -C "$$(go env GOPATH)"/src/github.com/grpc-ecosystem/grpc-gateway checkout $(PROTO_GATEWAY); \
	  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway; \
	  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger; \
	)


.PHONY: step-1
step-1:
	@rm -rf demo/*
	@git checkout fd9fe8dae0b4efb95bd96b0f58614b84f2f151c0 -- demo
	@make view

.PHONY: step-2
step-2:
	@rm -rf demo/*
	@git checkout fbe4d303757f50437e2a8d293036667ce162b707 -- demo
	@mkdir -p demo/internal/protocol/grpc/protobuf
	@protoc -Idemo/api/protobuf-spec \
	        --go_out=plugins=grpc,logtostderr=true:demo/internal/protocol/grpc/protobuf \
	        sequence.proto maintenance.proto
	@git diff fd9fe8dae0b4efb95bd96b0f58614b84f2f151c0.. demo/api
	@make format view

.PHONY: step-3
step-3:
	@rm -rf demo/*
	@git checkout d8cfbf295eddb6622f8fa690efcd945cf464b06d -- demo
	@make view wait
	@cd demo && go run cmd/server/main.go &
	@sleep 2
	@cd demo && go run cmd/client/main.go
	@ps aux | grep '[g]o run cmd/server' | awk '{print $$2}' | xargs kill -9 $1
	@ps aux | grep '[g]o-build.*exe/main' | awk '{print $$2}' | xargs kill -9 $1

.PHONY: step-4
step-4:
	@rm -rf demo/*
	@git checkout d3e1c969b9528b55cb83eedd3c751736ee91f1f8 -- demo
	@mkdir -p demo/internal/protocol/grpc/gateway
	@protoc -Idemo/api/protobuf-spec \
	        -I"$(GOPATH)"/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis \
	        --go_out=plugins=grpc,logtostderr=true:demo/internal/protocol/grpc/protobuf \
	        --grpc-gateway_out=logtostderr=true,import_path=gateway:demo/internal/protocol/grpc/gateway \
	        sequence.proto maintenance.proto
	@git checkout d3e1c969b9528b55cb83eedd3c751736ee91f1f8 -- demo/internal/protocol/grpc/gateway
	@make format
	@git diff d8cfbf295eddb6622f8fa690efcd945cf464b06d.. demo/api
	@make view wait
	@cd demo && go run cmd/server/main.go

.PHONY: step-5
step-5:
	@rm -rf demo/*
	@git checkout b2b1f7adcebcd482cc0d72ddcefddf066303d309 -- demo
	@mkdir -p demo/api/swagger-spec
	@protoc -Idemo/api/protobuf-spec \
	        -I"$(GOPATH)"/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis \
	        --go_out=plugins=grpc,logtostderr=true:demo/internal/protocol/grpc/protobuf \
	        --grpc-gateway_out=logtostderr=true,import_path=gateway:demo/internal/protocol/grpc/gateway \
	        --swagger_out=logtostderr=true,allow_merge=true,merge_file_name=service:demo/api/swagger-spec \
	        sequence.proto maintenance.proto
	@git checkout d3e1c969b9528b55cb83eedd3c751736ee91f1f8 -- demo/internal/protocol/grpc/gateway
	@make format
	@git diff d3e1c969b9528b55cb83eedd3c751736ee91f1f8.. demo/api
	@make view wait
	@cd demo && docker-compose -f deployments/docker-compose.yml up -d
	@sleep 2
	@open http://spec.127.0.0.1.xip.io
	@make wait
	@demo/bin/client -host service.127.0.0.1.xip.io:443
	@make wait
	@cd demo && docker-compose -f deployments/docker-compose.yml down

.PHONY: step-6
step-6:
	@rm -rf demo/*
	@git checkout 6dbc56957da30ddbbd9935dd71424e50007e24c7 -- demo
	@mkdir -p demo/scripts/py
	@( \
	  cd demo; \
	  python3 -m venv .virtenv; \
	  touch scripts/py/__init__.py; \
	  source .virtenv/bin/activate; \
	  pip install -r requirements.txt; \
	  python -m grpc_tools.protoc \
	         -Iapi/protobuf-spec \
	         -I"$(GOPATH)"/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis \
	         --python_out=scripts/py \
	         --grpc_python_out=scripts/py \
	         sequence.proto; \
	  python scripts/py/client.py; \
	  deactivate; \
	)
	@make wait


.PHONY: audit
audit:
	@ps aux | grep '[g]o'

.PHONY: build
build:
	@( \
	  cd demo; \
	  export _commit="-X main.commit=${COMMIT}"; \
	  export _date="-X main.date=${DATE}"; \
	  export _version="-X main.version=${VERSION}"; \
	  go build -ldflags "-s -w $${_commit} $${_date} $${_version}" \
	           -o bin/server \
	           cmd/server/main.go; \
	  go build -ldflags "-s -w $${_commit} $${_date} $${_version}" \
	           -o bin/client \
	           cmd/client/main.go; \
	  docker build -f build/package/Dockerfile \
	               -t $$(go list -m) \
	               --build-arg COMMIT=$(COMMIT) \
	               --build-arg DATE=$(DATE) \
	               --build-arg VERSION=$(VERSION) \
	               .; \
	)

.PHONY: format
format:
	@goimports -local workshops -ungroup -w demo

.PHONY: view
view:
	@tree demo

.PHONY: wait
wait:
	@read -p 'press ↵ to continue' next
