package redis

import "github.com/dengpju/higo-throw/throw"

type StringResult struct {
	Result string
	Err    error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

func (this *StringResult) Unwrap() string {
	if this.Err != nil {
		throw.Throw(this.Err, 0)
	}
	return this.Result
}

func (this *StringResult) Default(v string) string {
	if this.Err != nil {
		return v
	}
	return this.Result
}
