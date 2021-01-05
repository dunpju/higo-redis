package redis

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

const (
	SERILIZER_JSON = "json"
	SERILIZER_GOB  = "gob"
)

type DbGetterFunc func() interface{}

type SimpleCache struct {
	Operation *StringOperation
	Expire    *Parameter
	DbGetter  DbGetterFunc
	Serilizer string
	Policy    CachePolicy
}

func NewSimpleCache(operation *StringOperation, expire *Parameter, serilizer string, policy CachePolicy) *SimpleCache {
	policy.SetOperation(operation)
	return &SimpleCache{Operation: operation, Expire: expire, Serilizer: serilizer, Policy: policy}
}

func (this *SimpleCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, this.Expire)
}

func(this *SimpleCache) GetCacheForObject(key string,obj interface{})  interface{} {
	ret:=this.GetCache(key)
	if ret==nil{
		return nil
	}
	if this.Serilizer==SERILIZER_JSON{
		err:=json.Unmarshal([]byte(ret.(string)),obj)
		if err!=nil{
			return nil
		}
	}else if   this.Serilizer==SERILIZER_GOB{
		var buf =&bytes.Buffer{}
		buf.WriteString(ret.(string))
		dec:=gob.NewDecoder(buf)
		if dec.Decode(obj)!=nil{
			return nil
		}
	}
	return nil
}

func(this *SimpleCache) GetCache(key string) (ret interface{}){
	if this.Policy!=nil{ //检查策略
		this.Policy.Before(key)
	}
	if this.Serilizer==SERILIZER_JSON{
		f:= func()  string {
			obj:= this.DbGetter()
			b,err:=json.Marshal(obj)
			if err!=nil{
				return ""
			}
			return string(b)
		}
		ret=this.Operation.Get(key).DefaultFunc(f)
	}else if this.Serilizer==SERILIZER_GOB {
		f := func() string {
			obj:= this.DbGetter()
			var buf= &bytes.Buffer{}
			enc := gob.NewEncoder(buf)
			if err := enc.Encode(obj); err != nil {
				return ""
			}
			return buf.String()
		}
		ret = this.Operation.Get(key).DefaultFunc(f)

	}
	if ret.(string) == "" && this.Policy != nil {
		this.Policy.IfNil(key, "")
	} else {
		this.SetCache(key, ret)
	}
	return
}
