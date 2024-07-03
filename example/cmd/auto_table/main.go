package main

import (
	sparrow "github.com/lrayt/small-sparrow"
	"github.com/lrayt/small-sparrow/application"
	"github.com/lrayt/small-sparrow/example/Internal/database"
	"github.com/lrayt/small-sparrow/example/app/model"
	"log"
	"path/filepath"
)

var (
	AppName = "small_sparrow"
	Version = "0.1.0"
)

type Server struct {
	dbm *database.DBProvider
}

func NewServer(dbm *database.DBProvider) *Server {
	return &Server{dbm: dbm}
}

func (s Server) AutoTable() {
	if err := s.dbm.DB.AutoMigrate(model.OrderInfo{}); err != nil {
		log.Fatalf("AutoMigrate model.OrderInfo Err:%s\n", err.Error())
	}
}

func main() {
	rootPath, pathErr := filepath.Abs("")
	if pathErr != nil {
		log.Fatalf("获取项目工作路径失败,err:%s\n", pathErr.Error())
	}
	rootPath = filepath.Join(rootPath, "example")
	if err := sparrow.Init(AppName, Version, application.WithWorkerDir(rootPath)); err != nil {
		log.Fatalf("环境初始化失败,err:%s\n", err.Error())
	}
	svc, clean, err := InitServer()
	defer clean()
	if err != nil {
		log.Fatalf("服务初始化失败,err:%s\n", err.Error())
	}
	svc.AutoTable()
}
