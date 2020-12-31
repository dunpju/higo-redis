package redis

type StringResult struct {
	*Reply
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{&Reply{Result:result, Err: err}}
}

func NewStringsResult(result []string, err error) *StringResult {
	return &StringResult{&Reply{Result:result, Err: err}}
}

