package redis

import (
	"github.com/dengpju/higo-throw/throw"
)

type StringResult struct {
	Result string
	Err    error
}

func (this *StringResult) Output(out *interface{}) {
	*out = this.Result
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

func (this *StringResult) Unwrap() ResultType {
	if this.Err != nil {
		throw.Throw(this.Err, 0)
	}
	return this
}

func (this *StringResult) Default(v string) ResultType {
	if this.Err != nil {
		this.Result = v
		return this
	}
	return this
}

