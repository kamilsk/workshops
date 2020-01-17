// +build tools

package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "golang.org/x/tools/cmd/goimports"
)

//go:generate go build -v -o=../bin/mockgen github.com/golang/mock/mockgen
//go:generate go build -v -o=../bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go build -v -o=../bin/goimports golang.org/x/tools/cmd/goimports
