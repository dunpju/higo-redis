package tests

import (
	"github.com/dengpju/higo-redis/redis"
	"reflect"
	"testing"
)

func TestAttributes_Find(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		this redis.Attributes
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Find(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAttribute(t *testing.T) {
	type args struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.NewAttribute(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPoolConfigure(t *testing.T) {
	type args struct {
		Attr []*redis.Attribute
	}
	tests := []struct {
		name string
		args args
		want *redis.PoolConfigure
	}{
		// TODO: Add test cases.
		{
			name: "NewPoolConfigure",
			args: args{Attr: []*redis.Attribute{
				redis.PoolHost("192.168.8.99"),
				redis.PoolPort(6379),
				redis.PoolAuth("1qaz2wsx"),
				redis.PoolDb(0),
				redis.PoolMaxConnections(10),
				redis.PoolMaxIdle(3),
				redis.PoolMaxIdleTime(60),
				redis.PoolMaxConnLifetime(10),
				redis.PoolWait(true)},
			},
			want: &redis.PoolConfigure{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.NewPoolConfigure(tt.args.Attr...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoolConfigure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolAuth(t *testing.T) {
	type args struct {
		auth string
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolAuth(tt.args.auth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolDb(t *testing.T) {
	type args struct {
		db int
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolDb(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolHost(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolHost(tt.args.host); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolMaxConnLifetime(t *testing.T) {
	type args struct {
		maxConnLifetime int
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolMaxConnLifetime(tt.args.maxConnLifetime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolMaxConnLifetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolMaxConnections(t *testing.T) {
	type args struct {
		maxConnections int
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolMaxConnections(tt.args.maxConnections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolMaxConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolMaxIdle(t *testing.T) {
	type args struct {
		maxIdle int
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolMaxIdle(tt.args.maxIdle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolMaxIdle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolMaxIdleTime(t *testing.T) {
	type args struct {
		maxIdleTime int
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolMaxIdleTime(tt.args.maxIdleTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolMaxIdleTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolPort(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolPort(tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolWait(t *testing.T) {
	type args struct {
		wait bool
	}
	tests := []struct {
		name string
		args args
		want *redis.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redis.PoolWait(tt.args.wait); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoolWait() = %v, want %v", got, tt.want)
			}
		})
	}
}
