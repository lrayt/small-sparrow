//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lrayt/small-sparrow/example/Internal/database"
)

var InternalProvider = wire.NewSet(
	database.NewDBProvider,
)

func InitServer() (*Server, func(), error) {
	panic(wire.Build(InternalProvider, NewServer))
}
