package redis

import "github.com/dengpju/higo-throw/throw"

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
		return this
	}
	return this
}

func (this *Reply) Error() {
	if this.Err != nil {
		throw.Throw(this.Err, 0)
	}
}

func (this *Reply) ToString() string {
	this.Error()
	return this.Result.(string)
}

func (this *Reply) ToStrings() []string {
	this.Error()
	return this.Result.([]string)
}

func (this *Reply) ToInt() int {
	this.Error()
	return this.Result.(int)
}

func (this *Reply) ToSlice() []interface{} {
	this.Error()
	return this.Result.([]interface{})
}

func (this *Reply) Iterable() *Iterator {
	return NewIterator(this.Result.([]string))
}
