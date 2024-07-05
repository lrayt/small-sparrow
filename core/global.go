package core

import (
	"github.com/lrayt/small-sparrow/core/abstract"
	"github.com/lrayt/small-sparrow/core/runtime"
)

// GConfigs 全局配置
func GConfigs() abstract.ConfigProvider {
	return app.ConfigProvider
}

func IsProdEnv() bool {
	return app.Env.RunEnv == runtime.RunProdEnv
}
func IsTestEnv() bool {
	return app.Env.RunEnv == runtime.RunTestEnv
}
func IsLocalEnv() bool {
	return app.Env.RunEnv == runtime.RunLocalEnv
}

func GWorkDir() {

}

func GResourceDir() {

}
