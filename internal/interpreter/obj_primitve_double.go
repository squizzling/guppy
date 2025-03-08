package interpreter

type ObjectDouble struct {
	Object

	Value float64
}

func NewObjectDouble(f float64) Object {
	return &ObjectDouble{
		Object: NewObject(nil),
		Value:  f,
	}
}
