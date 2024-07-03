package core

import "github.com/lrayt/small-sparrow/core/abstract"

// GConfigs 全局配置
func GConfigs() abstract.ConfigProvider {
	return app.ConfigProvider
}

// GRunEnv 运行环境
func GRunEnv() string {
	return string(app.Env.RunEnv)
}

func GWorkDir() {

}

func GResourceDir() {

}
