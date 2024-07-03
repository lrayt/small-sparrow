//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lrayt/small-sparrow/example/app/handler"
)

//var InternalProvider = wire.NewSet(
//	database.NewDBProvider,
//	database.NewCacheProvider,
//	http_manager.NewGinHttpProvider,
//	message.NewMQProvider,
//)
//
//// DaoProvider 数据库操作
//var DaoProvider = wire.NewSet(
//	dao.NewSceneThemeDao,
//	dao.NewSubjectDao,
//	dao.NewServiceThemeDao,
//	dao.NewServiceAccessDao,
//	dao.NewAccessBuilderDao,
//	dao.NewServiceDao,
//	dao.NewAlarmInfoDao,
//	dao.NewAlarmPushDao,
//	dao.NewAlarmRuleDao,
//	dao.NewStatisticsDao,
//	dao.NewRelationshipsAlarmServiceDao,
//	dao.NewRelationshipsThemeServiceDao,
//	dao.NewRelationshipsServiceAccessDao,
//	dao.NewRelationshipsSubjectApiDao,
//	dao.NewXZQDao,
//	dao.NewAPIDao,
//	dao.NewAttachInfoDao,
//	dao.NewServerInfoDao,
//	dao.NewSysDictDao,
//	dao.NewRouterInfoDao,
//)
//
//// ServiceProvider 业务处理
//var ServiceProvider = wire.NewSet(
//	service.NewSceneThemeService,
//	service.NewProxyService,
//	service.NewSubjectService,
//	service.NewServiceThemeService,
//	service.NewServiceAccessService,
//	service.NewAccessBuilderService,
//	service.NewServiceInfoService,
//	service.NewRelationshipsServiceAccess,
//	service.NewRelationshipsThemeService,
//	service.NewRelationshipsSubjectApi,
//	service.NewXZQService,
//	service.NewAttachInfoService,
//	service.NewAlarmInfoService,
//	service.NewStatisticsService,
//	service.NewSysDictService,
//)

// HandlerProvider 获取参数
var HandlerProvider = wire.NewSet(
	handler.NewHttpHandler,
)

func InitExampleServer() (*ExampleServer, func(), error) {
	panic(wire.Build(HandlerProvider, NewExampleServer))
}
