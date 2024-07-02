package database

import (
	sparrow "github.com/lrayt/small-sparrow"
	"github.com/lrayt/small-sparrow/kit/db_builder"
	"gorm.io/gorm"
)

type DBProvider struct {
	DB *gorm.DB
}

func NewDBProvider() (*DBProvider, error) {
	options := new(db_builder.Options)
	if err := sparrow.GConfigs().PackConf("database.scp-db", options); err != nil {
		return nil, err
	}
	if db, err := db_builder.CreateGormDB(options); err != nil {
		return nil, err
	} else {
		return &DBProvider{DB: db}, nil
	}
}
