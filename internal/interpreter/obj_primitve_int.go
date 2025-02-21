package interpreter

import (
	"strconv"
)

type ObjectInt struct {
	Object

	Value int
}

func NewObjectInt(i int) Object {
	return &ObjectInt{
		Value: i,
	}
}

func (oi *ObjectInt) Int(i *Interpreter) (int, error) {
	return oi.Value, nil
}

func (oi *ObjectInt) String(i *Interpreter) (string, error) {
	return strconv.Itoa(oi.Value), nil
}

var _ = FlowStringable(&ObjectInt{})
var _ = FlowIntable(&ObjectInt{})
