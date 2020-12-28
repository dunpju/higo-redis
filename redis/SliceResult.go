package redis

type SliceResult struct {
	*Reply
}

func NewSliceResult(result []string, err error) *StringResult {
	return &StringResult{&Reply{Result:result, Err: err}}
}
