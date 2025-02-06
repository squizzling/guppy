package interpreter

type ObjectNone struct {
	Object
}

func NewObjectNone() Object {
	return &ObjectNone{}
}
