package interpreter

type ObjectDouble struct {
	Object

	Value float64
}

func NewObjectDouble(f float64) Object {
	return &ObjectDouble{
		Value: f,
	}
}

func (od *ObjectDouble) Double(i *Interpreter) (float64, error) {
	return od.Value, nil
}
