package redis

type DbGetterFunc func() string

type SimpleCache struct {
	Operation *StringOperation
	Expire    *Parameter
	DbGetter  DbGetterFunc
}

func NewSimpleCache(operation *StringOperation, expire *Parameter) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire}
}

func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, this.Expire)
}

func (this *SimpleCache) GetCache(key string) (ret string) {
	ret = this.Operation.Get(key).DefaultFunc(this.DbGetter).String()
	this.SetCache(key, ret)
	return
}
