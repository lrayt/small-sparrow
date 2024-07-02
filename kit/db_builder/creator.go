package db_builder

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func CreateGormDB(options *Options) (*gorm.DB, error) {
	options.SetDefault() // 设置默认值
	var dialect gorm.Dialector
	switch options.Driver {
	case DBDriverMysql:
		var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", options.User, options.Password, options.Host, options.DBName, options.Charset)
		dialect = mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		})

	case DBDriverPostgresL:
		var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai client_encoding=%s", options.Host, options.User, options.Password, options.DBName, options.Port, options.Charset)
		dialect = postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		})
	default:
		return nil, fmt.Errorf("unkown db driver:%s\n", options.Driver)
	}

	var database, openErr = gorm.Open(dialect)
	if openErr != nil {
		return nil, openErr
	}
	// 获取数据库
	var db, dbErr = database.DB()
	if dbErr != nil {
		return nil, dbErr
	}
	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(options.MaxIdleConn)                                 // 设置空闲连接池中连接的最大数量
	db.SetMaxOpenConns(options.MaxOpenConn)                                 // 设置打开数据库连接的最大数量。
	db.SetConnMaxLifetime(time.Minute * time.Duration(options.MaxLifetime)) // 设置了连接可复用的最大时间。
	return database, nil
}
