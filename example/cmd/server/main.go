package main

import (
	"github.com/lrayt/small-sparrow/core"
	"github.com/lrayt/small-sparrow/example/Internal/database"
	"github.com/lrayt/small-sparrow/example/app/handler"
	"log"
	"path/filepath"
)

type ExampleServer struct {
	HttpHandler *handler.HttpHandler
	dbm         *database.DBManager
}

func NewExampleServer(httpHandler *handler.HttpHandler, dbm *database.DBManager) (*ExampleServer, error) {
	rootPath, pathErr := filepath.Abs("")
	if pathErr != nil {
		log.Fatalf("获取项目工作路径失败,err:%s\n", pathErr.Error())
	}
	rootPath = filepath.Join(rootPath, "example")
	if err := core.InitApp(
		core.WithHandler(httpHandler),
		core.WithWorkerDir(rootPath),
		core.WithStarter(dbm),
	); err != nil {
		return nil, err
	}
	return &ExampleServer{HttpHandler: httpHandler}, nil
}

func (s ExampleServer) Setup() {
	core.SetupApp()
}

func main() {
	svc, cleanup, err := InitExampleServer()
	defer cleanup()
	if err != nil {
		log.Fatalf("服务启动失败,err:%s\n", err.Error())
	}
	svc.Setup()
}
