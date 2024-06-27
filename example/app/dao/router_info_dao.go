package dao

import "github.com/lrayt/small-sparrow/example/app/model"

type RouterInfoDao struct {
}

func (d RouterInfoDao) FindRouterInfo(map[string]interface{}) (*model.ServiceRouter, error) {
	return nil, nil
}
