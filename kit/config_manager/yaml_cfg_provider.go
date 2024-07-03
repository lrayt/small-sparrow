package config_manager

import (
	"errors"
	"fmt"
	"github.com/lrayt/small-sparrow/core/runtime"
	"github.com/spf13/viper"
	"path/filepath"
)

type YamlConfigProvider struct {
	medium *viper.Viper
}

func NewYamlConfigProvider(globalEnv *runtime.Env) (*YamlConfigProvider, error) {
	cfgFile := filepath.Join(globalEnv.WorkDir, "resource", "conf", fmt.Sprintf("skeleton-%s-conf.yaml", globalEnv.RunEnv))
	cfgContainer := &YamlConfigProvider{medium: viper.New()}
	cfgContainer.medium.SetConfigType("yaml")
	cfgContainer.medium.SetConfigFile(cfgFile)
	if err := cfgContainer.medium.ReadInConfig(); err != nil {
		return nil, errors.New(fmt.Sprintf("加载配置文件[%s]失败,err: %s", cfgFile, err))
	} else {
		fmt.Printf("加载配置文件[%s]\n", cfgFile)
	}
	return cfgContainer, nil
}

func (y YamlConfigProvider) GetValue(key string) interface{} {
	return y.medium.Get(key)
}

func (y YamlConfigProvider) PackConf(cfgId string, obj interface{}) error {
	return y.medium.UnmarshalKey(cfgId, obj)
}

func (y YamlConfigProvider) PackConfToMap(cfgId string) map[string]interface{} {
	return y.medium.GetStringMap(cfgId)
}

func (y YamlConfigProvider) GetIntSlice(key string) []int {
	return y.medium.GetIntSlice(key)
}
