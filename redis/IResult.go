package redis

import (
	"github.com/dengpju/higo-throw/throw"
)

type IResult interface {
	Unwrap() *Reply
	Default(v string) *Reply
}

type Reply struct {
	Result interface{}
	Err    error
}

func (this *Reply) Unwrap() *Reply {
	return this
}

func (this *Reply) Default(v string) *Reply {
	if this.Err != nil {
		this.Result = v
		this.Err = nil
		return this
	}
	return this
}

func (this *Reply) Error() {
	if this.Err != nil {
		throw.Throw(this.Err, 0)
	}
}

func (this *Reply) String() string {
	this.Error()
	return this.Result.(string)
}

func (this *Reply) Strings() []string {
	this.Error()
	return this.Result.([]string)
}

func (this *Reply) Int() int {
	this.Error()
	return this.Result.(int)
}

func (this *Reply) Bool() bool {
	this.Error()
	return this.Result.(bool)
}

func (this *Reply) Slice() []interface{} {
	this.Error()
	return this.Result.([]interface{})
}

func (this *Reply) Iterable() *Iterator {
	return NewIterator(this.Result.([]string))
}
