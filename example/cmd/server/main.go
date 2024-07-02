package main

import (
	"github.com/lrayt/small-sparrow"
	"github.com/lrayt/small-sparrow/application"
	"github.com/lrayt/small-sparrow/example/app/handler"
	"log"
	"path/filepath"
)

type ExampleServer struct {
	HttpHandler *handler.HttpHandler
}

func NewExampleServer(httpHandler *handler.HttpHandler) *ExampleServer {
	return &ExampleServer{HttpHandler: httpHandler}
}

func main() {
	rootPath, pathErr := filepath.Abs("")
	if pathErr != nil {
		log.Fatalf("获取项目工作路径失败,err:%s\n", pathErr.Error())
	}
	rootPath = filepath.Join(rootPath, "example")
	if err := sparrow.Init("light-boot", "0.0.1", application.WithWorkerDir(rootPath)); err != nil {
		log.Fatalf("init app err:%s\n", err.Error())
	}
	//svc, cleanup, err := InitExampleServer()
	//defer cleanup()
	//if err != nil {
	//	log.Fatalf("服务启动失败,err:%s\n", err.Error())
	//}
	////sparrow.AddHandler(svc.HttpHandler)
	//sparrow.Setup()
}
