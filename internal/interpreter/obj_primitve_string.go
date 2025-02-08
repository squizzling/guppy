package interpreter

type ObjectString struct {
	Object

	s string
}

func NewObjectString(s string) Object {
	return &ObjectString{
		s: s,
	}
}

func (os *ObjectString) String(i *Interpreter) (string, error) {
	return os.s, nil
}

var _ = FlowStringable(&ObjectString{})
