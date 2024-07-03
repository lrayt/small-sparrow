// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lrayt/small-sparrow/example/app/handler"
)

// Injectors from wire.go:

func InitExampleServer() (*ExampleServer, func(), error) {
	httpHandler := handler.NewHttpHandler()
	exampleServer, err := NewExampleServer(httpHandler)
	if err != nil {
		return nil, nil, err
	}
	return exampleServer, func() {
	}, nil
}

// wire.go:

// HandlerProvider 获取参数
var HandlerProvider = wire.NewSet(handler.NewHttpHandler)