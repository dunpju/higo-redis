package redis

const (
	HOST               = "Host"
	PORT               = "Port"
	AUTH               = "Auth"
	DB_NUM             = "Db"
	MAX_CONNECTIONS    = "MaxConnections"
	MAX_IDLE           = "MaxIdle"
	MAX_IDLE_TIME      = "MaxIdleTime"
	MAX_CONN_LIFE_TIME = "MaxConnLifetime"
	WAIT               = "Wait"
)

type PoolConfigure struct {
	Host            string
	Port            int
	Auth            string
	Db              int
	MaxConnections  int
	MaxIdle         int
	MaxIdleTime     int
	MaxConnLifetime int
	Wait            bool
}

func NewPoolConfigure(Attr ...*Attribute) *PoolConfigure {
	pool := &PoolConfigure{Host: "127.0.0.1", Port: 6379, Db: 0}
	host := Attributes(Attr).Find(HOST)
	if host != nil {
		pool.Host = host.(string)
	}
	port := Attributes(Attr).Find(PORT)
	if port != nil {
		pool.Port = port.(int)
	}
	auth := Attributes(Attr).Find(AUTH)
	if auth != nil {
		pool.Auth = auth.(string)
	}
	db := Attributes(Attr).Find(DB_NUM)
	if db != nil {
		pool.Db = db.(int)
	}
	maxConnections := Attributes(Attr).Find(MAX_CONNECTIONS)
	if maxConnections != nil {
		pool.MaxConnections = maxConnections.(int)
	}
	maxIdle := Attributes(Attr).Find(MAX_IDLE)
	if maxIdle != nil {
		pool.MaxIdle = maxIdle.(int)
	}
	maxIdleTime := Attributes(Attr).Find(MAX_IDLE_TIME)
	if maxIdleTime != nil {
		pool.MaxIdleTime = maxIdleTime.(int)
	}
	maxConnLifetime := Attributes(Attr).Find(MAX_IDLE_TIME)
	if maxConnLifetime != nil {
		pool.MaxConnLifetime = maxConnLifetime.(int)
	}
	wait := Attributes(Attr).Find(WAIT)
	if wait != nil {
		pool.Wait = wait.(bool)
	}
	return pool
}

type Attribute struct {
	Name  string
	Value interface{}
}

func NewAttribute(name string, value interface{}) *Attribute {
	return &Attribute{Name: name, Value: value}
}

type Attributes []*Attribute

func (this Attributes) Find(name string) interface{} {
	for _, p := range this {
		if p.Name == name {
			return p.Value
		}
	}
	return nil
}

func PoolHost(host string) *Attribute {
	return NewAttribute(HOST, host)
}

func PoolPort(port int) *Attribute {
	return NewAttribute(PORT, port)
}

func PoolAuth(auth string) *Attribute {
	return NewAttribute(AUTH, auth)
}

func PoolDb(db int) *Attribute {
	return NewAttribute(DB_NUM, db)
}

func PoolMaxConnections(maxConnections int) *Attribute {
	return NewAttribute(MAX_CONNECTIONS, maxConnections)
}

func PoolMaxIdle(maxIdle int) *Attribute {
	return NewAttribute(MAX_IDLE, maxIdle)
}

func PoolMaxIdleTime(maxIdleTime int) *Attribute {
	return NewAttribute(MAX_IDLE_TIME, maxIdleTime)
}

func PoolMaxConnLifetime(maxConnLifetime int) *Attribute {
	return NewAttribute(MAX_CONN_LIFE_TIME, maxConnLifetime)
}

func PoolWait(wait bool) *Attribute {
	return NewAttribute(WAIT, wait)
}
