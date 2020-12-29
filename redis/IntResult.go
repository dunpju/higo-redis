package redis

type IntResult struct {
	*Reply
}

func NewIntResult(result int, err error) *BoolResult {
	return &BoolResult{&Reply{Result:result, Err:err}}
}


