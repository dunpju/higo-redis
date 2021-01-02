package redis

import "encoding/json"

const SERILIZER_JSON  = "json"

type DbGetterFunc func() interface{}

type SimpleCache struct {
	Operation *StringOperation
	Expire    *Parameter
	DbGetter  DbGetterFunc
	Serilizer string
}

func NewSimpleCache(operation *StringOperation, expire *Parameter, serilizer string) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire, Serilizer: serilizer}
}

func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, this.Expire)
}

func (this *SimpleCache) GetCache(key string) (ret interface{}) {
	if this.Serilizer == SERILIZER_JSON {
		f := func() string {
			obj := this.DbGetter()
			b,err := json.Marshal(obj)
			if err!=nil {
				return ""
			}
			return string(b)
		}
		ret = this.Operation.Get(key).DefaultFunc(f).String()
		this.SetCache(key, ret)
	}

	return
}
