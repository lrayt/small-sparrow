package database

import (
	"context"
	"database/sql"
	"github.com/lrayt/small-sparrow/builder"
	"github.com/lrayt/small-sparrow/core"
	"gorm.io/gorm"
	"time"
)

type DBManager struct {
	GormDB *gorm.DB
}

func (p *DBManager) Init() error {
	var (
		err     error
		options = new(builder.DBOptions)
	)
	// cfg
	err = core.GConfigs().PackConf("database.scp-db", options)
	if err != nil {
		return err
	}
	// gorm db
	p.GormDB, err = builder.CreateGormDB(options)
	if err != nil {
		return err
	}
	// sql db
	var sqlDB *sql.DB
	sqlDB, err = p.GormDB.DB()
	if err != nil {
		return err
	}
	// ping
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return sqlDB.PingContext(ctx)
}

func (p DBManager) Close() error {
	db, err := p.GormDB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func NewDBManager() *DBManager {
	return new(DBManager)
}
