package model

import "time"

type ServiceRouter struct {
	ID               string
	ThemeCode        string
	ServiceCode      string
	ServiceId        string
	AccessUUID       string
	RegionCode       string
	TargetUrl        string
	AlarmRule        string
	ServiceType      int8
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	ServiceName      string
	ThemeName        string
	ThemeType        string
	AccessRegionCode string
	SystemName       string
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
