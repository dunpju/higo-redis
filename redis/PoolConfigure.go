package redis

type PoolConfigure struct {
	Host int
	Port int
	Auth string
	Db int
	MaxConnections int
	MaxIdle int
	MaxIdleTime int
}
