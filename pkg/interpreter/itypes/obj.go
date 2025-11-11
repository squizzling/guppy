package itypes

import (
	"fmt"
)

type flowObject struct {
	members map[string]Object
}

func NewObject(attributes map[string]Object) Object {
	return &flowObject{
		members: attributes,
	}
}

func (f *flowObject) Member(i Interpreter, obj Object, memberName string) (Object, error) {
	if f.members == nil {
		return nil, fmt.Errorf("object of type %T does not support member lookup for %s", obj, memberName)
	} else if member, ok := f.members[memberName]; !ok {
		return nil, fmt.Errorf("object of type %T does not contain member %s", obj, memberName)
	} else {
		return NewLValue(obj, member), nil
	}
}
