// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lrayt/small-sparrow/example/Internal/database"
)

// Injectors from wire.go:

func InitServer() (*Server, func(), error) {
	dbProvider, err := database.NewDBProvider()
	if err != nil {
		return nil, nil, err
	}
	server := NewServer(dbProvider)
	return server, func() {
	}, nil
}

// wire.go:

var InternalProvider = wire.NewSet(database.NewDBProvider)
