package redis

type BoolResult struct {
	*Reply
}

func NewBoolResult(result bool, err error) *BoolResult {
	return &BoolResult{&Reply{Result:result, Err:err}}
}


