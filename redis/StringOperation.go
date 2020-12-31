package redis

type StringOperation struct {
}

func NewStringOperation() *StringOperation {
	return &StringOperation{}
}

func (this *StringOperation) Set(key string, v interface{}, args ...*Parameter) string {
	return Redis.Set(key, v, args...)
}

func (this *StringOperation) Get(key string) *StringResult {
	return Redis.get(key)
}

func (this *StringOperation) Mget(key string, keys ...*Parameter) []string {
	return Redis.Mget(key, keys...)
}
