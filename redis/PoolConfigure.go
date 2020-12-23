package redis

type PoolConfigure struct {
	Host string
	Port int
	Auth string
	Db int
	MaxConnections int
	MaxIdle int
	MaxIdleTime int
}
