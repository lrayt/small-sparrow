package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"testing"
	"time"
)

type ServiceRouter struct {
	ID          string //  `gorm:"column:ID;primaryKey;type:char(32)" json:"id"`
	ThemeCode   string // `gorm:"column:THEME_CODE;index;type:varchar(32);not null;comment:服务主题编码" json:"theme_code"`
	ServiceCode string //   `gorm:"column:SERVICE_CODE;index;type:varchar(32);not null;comment:服务编码" json:"service_code"`
	ServiceId   string //  `gorm:"column:SERVICE_ID;index;type:varchar(32);not null;comment:服务Id" json:"service_id"`
	AccessUUID  string //  `gorm:"column:ACCESS_UUID;index;type:varchar(32);not null;comment:接入方唯一码" json:"access_uuid"`
	RegionCode  string //  `gorm:"column:REGION_CODE;index;type:varchar(45);not null;comment:行政区划代码" json:"region_code"`
	TargetUrl   string //  `gorm:"column:TARGET_URL;type:varchar(256);not null;comment:目标地址" json:"target_url"`
	AlarmRule   string //   `gorm:"column:ALARM_CONDITION;type:varchar(32);comment:告警规则json形式" json:"alarm_rule"`
	ServiceType int8   //  `gorm:"column:SERVICE_TYPE;comment:服务类型" json:"service_type"`
	//State            ServiceState //  `gorm:"column:STATE;comment:状态" json:"state"`
	//Method           HttpMethod   //  `gorm:"column:METHOD;not null;comment:Http方法" json:"method"`
	CreatedAt        *time.Time // `gorm:"column:CREATED_AT;autoCreateTime" json:"-"`
	UpdatedAt        *time.Time // `gorm:"column:UPDATED_AT;autoUpdateTime" json:"-"`
	ServiceName      string     `json:"service_name"`       // 服务名称
	ThemeName        string     `json:"theme_name"`         // 服务主题名称
	ThemeType        string     `json:"theme_type"`         // 服务主题类型
	AccessRegionCode string     `json:"access_region_code"` // 调用方行政区划代码
	SystemName       string     `json:"system_name"`        // 调用方系统名称
}

func (t ServiceRouter) TableName() string {
	return "public.service_router"
}
func (t ServiceRouter) CacheKey() string {
	return t.ThemeCode + "-" + t.ServiceCode + "-" + t.AccessUUID
}
func (t ServiceRouter) Key() string {
	return "Index-" + t.ThemeCode + "-" + t.ServiceCode + "-" + t.AccessUUID
}
func (t ServiceRouter) LockKey() string {
	return "Lock-" + t.ThemeCode + "-" + t.ServiceCode + "-" + t.AccessUUID
}

var rdb *redis.Client
var dbm *gorm.DB

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.2.54:6379",
		Password: "123456",
		DB:       1,
		PoolSize: 50,
	})
	str, err := rdb.Ping().Result()
	log.Println("ping:", str, err)
	dsn := "host=192.168.2.54 user=system password=gtis dbname=test port=54321 sslmode=disable TimeZone=Asia/Shanghai client_encoding=UTF8"
	if db, err1 := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{}); err1 != nil {
		log.Fatalf("open database err:%s\n", err1.Error())
	} else {
		if sqlDB, err2 := db.DB(); err2 == nil {
			if err3 := sqlDB.Ping(); err3 != nil {
				log.Fatalf("ping err:%s\n", err3.Error())
			}
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)
		}
		dbm = db.Debug()
	}
}

func GetData(r *ServiceRouter) (*ServiceRouter, error) {
	var router = new(ServiceRouter)
	if data, err := rdb.Get(r.CacheKey()).Bytes(); err == nil && len(data) > 0 {
		if err1 := json.Unmarshal(data, router); err1 == nil {
			return router, nil
		}
	}

	lockAcquired, setErr := rdb.SetNX(router.LockKey(), 1, 10*time.Second).Result()
	if setErr != nil {
		return nil, setErr
	}

	if lockAcquired {
		if err := dbm.Model(&ServiceRouter{}).Where("theme_code = ? AND service_code = ? AND access_uuid = ?", r.ThemeCode, r.ServiceCode, r.AccessUUID).First(router).Error; err != nil {
			return nil, err
		}
		data, err1 := json.Marshal(router)
		if err1 != nil {
			return nil, fmt.Errorf("set cache err:%s\n", err1.Error())
		}
		if err := rdb.Set(router.CacheKey(), string(data), 0).Err(); err != nil {
			log.Printf("set cache err:%s\n", err.Error())
		}
		if _, err := rdb.Incr(router.Key()).Result(); err != nil {
			log.Printf("Incr cache err:%s\n", err.Error())
		}
	} else {
		//time.Sleep(time.Second * 3)
		return GetData(r)
	}
	return router, nil

}

func TestCache(t *testing.T) {
	// PT3202009100001
	// S22440000000000w202010100001
	// 697dd23d323141eb9509b8ec4687a71e
	//data, err := GetData(&ServiceRouter{
	//	ThemeCode:   "PT3202009100001",
	//	ServiceCode: "S22440000000000w202010100001",
	//	AccessUUID:  "697dd23d323141eb9509b8ec4687a71e",
	//})
	//t.Log(data, "123", err)
	//time.Sleep(10 * time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			data, err := GetData(&ServiceRouter{
				ThemeCode:   "PT3202009100001",
				ServiceCode: "S22440000000000w202010100001",
				AccessUUID:  "697dd23d323141eb9509b8ec4687a71e",
			})
			log.Println(data, "err:", err)
			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(100 * time.Second)
}
