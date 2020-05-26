// +build tools

package main

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow"
	_ "golang.org/x/net/context"
)
