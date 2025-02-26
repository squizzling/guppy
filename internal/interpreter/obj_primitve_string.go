package interpreter

type ObjectString struct {
	Object

	Value string
}

func NewObjectString(s string) Object {
	return &ObjectString{
		Value: s,
	}
}

func (os *ObjectString) String(i *Interpreter) (string, error) {
	return os.Value, nil
}

var _ = FlowStringable(&ObjectString{})
