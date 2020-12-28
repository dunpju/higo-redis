package redis

type Iterator struct {
	value []string
	index int
}

func NewIterator(value []string) *Iterator {
	return &Iterator{value: value}
}

func (this *Iterator) HasNext() bool {
	if this.value == nil || len(this.value) == 0 {
		return false
	}
	return this.index < len(this.value)
}

func (this *Iterator) Next() (ret string) {
	ret = this.value[this.index]
	this.index = this.index + 1
	return
}
