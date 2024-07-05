package builder

type DBDriver string

const (
	DBDriverMysql     DBDriver = "mysql"
	DBDriverPostgresL          = "postgres"
)

type DBOptions struct {
	Driver      DBDriver
	User        string
	Password    string
	Host        string
	Port        uint
	DBName      string
	SchemaName  string
	Charset     string
	MaxIdleConn int // 最大空闲连接数
	MaxOpenConn int // 最大连接数
	MaxLifetime int // 最大连接时长
}

func (o *DBOptions) SetDefault() {
	if o.MaxIdleConn <= 0 {
		o.MaxIdleConn = 10
	}
	if o.MaxOpenConn <= 0 {
		o.MaxOpenConn = 100
	}
	if o.MaxLifetime <= 0 {
		o.MaxLifetime = 30
	}
}
