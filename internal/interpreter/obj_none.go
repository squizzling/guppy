package interpreter

import (
	"errors"
)

type ObjectNone struct{}

func NewObjectNone() Object {
	return &ObjectNone{}
}

func (on *ObjectNone) Repr() string {
	return "None"
}

func (on *ObjectNone) Member(i *Interpreter, obj Object, memberName string) (Object, error) {
	return nil, errors.New("None doesn't support member")
}
