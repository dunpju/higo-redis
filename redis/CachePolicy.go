package redis

import "regexp"

type CachePolicy interface {
	Before(key string)
}

// 穿透策略
type CrossPolicy struct {
	KeyRegx string
}

func NewCrossPolicy(keyRegx string) *CrossPolicy {
	return &CrossPolicy{KeyRegx: keyRegx}
}

func (this *CrossPolicy)Before(key string)  {
	if regexp.MustCompile(this.KeyRegx).MatchString(key) {
		panic("error cache key")
	}
}