package redis

import "time"

const PARAM_EXPIRE  = "expire"

type Parameter struct {
	Name  string
	Value interface{}
}

func NewParameter(name string, value interface{}) *Parameter {
	return &Parameter{Name: name, Value: value}
}

type Parameters []*Parameter

func (this Parameters)Find(name string) interface{} {
	for _,p := range this {
		if p.Name == name {
			return p.Value
		}
	}
	return nil
}

func WithExpire(t time.Duration) *Parameter  {
	return NewParameter(PARAM_EXPIRE, t)
}