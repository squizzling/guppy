package interpreter

import (
	"fmt"

	"guppy/pkg/interpreter/itypes"
)

type flowObject struct {
	members map[string]itypes.Object
}

func NewObject(attributes map[string]itypes.Object) itypes.Object {
	return &flowObject{
		members: attributes,
	}
}

func (f *flowObject) Member(i itypes.Interpreter, obj itypes.Object, memberName string) (itypes.Object, error) {
	if f.members == nil {
		return nil, fmt.Errorf("object of type %T does not support member lookup for %s", obj, memberName)
	} else if member, ok := f.members[memberName]; !ok {
		return nil, fmt.Errorf("object of type %T does not contain member %s", obj, memberName)
	} else {
		return NewLValue(obj, member), nil
	}
}
