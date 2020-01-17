BINPATH   = $(PWD)/bin
GO        = GOPRIVATE=$(GOPRIVATE) GOFLAGS=$(GOFLAGS) go
GOFLAGS   =
GOPRIVATE = go.avito.ru
LOCAL     = $(MODULE)
MODULE    = $(shell $(GO) list -m)
PATHS     = $(shell $(GO) list ./... 2> /dev/null | sed -e "s|$(MODULE)/\{0,1\}||g")
SHELL     = /bin/bash -euo pipefail

export PATH := $(BINPATH):$(PATH)

.PHONY: deps
deps: tools
	@$(GO) mod tidy

.PHONY: tools
tools:
	@cd tools && $(GO) mod tidy && $(GO) generate tools.go


.PHONY: generate
generate: avito-generate go-generate

.PHONY: avito-generate
avito-generate:
	@avito codegen

.PHONY: go-generate
go-generate:
	@$(GO) generate ./...


.PHONY: format
format:
	@goimports -local $(LOCAL) -ungroup -w $(PATHS)


.PHONY: lint
lint:
	@golangci-lint run ./...


.PHONY: test
test:
	@$(GO) test -race ./...
