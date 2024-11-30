package abstraction

type ReadOnly interface {
	GetValue() int
}

type dataStruct struct {
	value int
}

func (data dataStruct) GetValue() int {
	return data.value
}

func NewReadOnlyData(num int) ReadOnly {
	return dataStruct{value: num}
}
