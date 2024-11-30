package syncOnce

import "sync"

type ReadOnly struct {
	value int
	once  sync.Once
}

func (readOnly *ReadOnly) SetValue(val int) {
	readOnly.once.Do(func() {
		readOnly.value = val
	})
}

func (readOnly *ReadOnly) GetValue() int {
	return readOnly.value
}
