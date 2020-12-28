package redis

type SliceResult struct {
	*Reply
}

func NewSliceResult(result []string, err error) *StringResult {
	return &StringResult{&Reply{Result:result, Err: err}}
}

func (this *SliceResult) Unwrap() *Reply {
	return this.Reply
}

func (this *SliceResult) Default(v string) *Reply {
	if this.Err != nil {
		this.Result = v
		return this.Reply
	}
	return this.Reply
}
