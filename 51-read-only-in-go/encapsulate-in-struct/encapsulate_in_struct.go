package encapsulate

type ReadyOnly struct {
	value int
}

func NewReadOnly(val int) ReadyOnly {
	return ReadyOnly{value: val}
}

func (roValue ReadyOnly) GetValue() int {
	return roValue.value
}
