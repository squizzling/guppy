package interpreter

import (
	"fmt"

	"github.com/squizzling/types/pkg/result"
)

type Object interface {
	// We include the root Object so the default behavior knows its object type, instead of just *flowObject
	Member(i *Interpreter, obj Object, memberName string) result.Result[Object]
}

type flowObject struct {
	members map[string]Object
}

func NewObject(attributes map[string]Object) Object {
	return &flowObject{
		members: attributes,
	}
}

func (f *flowObject) Member(i *Interpreter, obj Object, memberName string) result.Result[Object] {
	if f.members == nil {
		return result.Err[Object](fmt.Errorf("object of type %T does not support member lookup for %s", obj, memberName))
	} else if member, ok := f.members[memberName]; !ok {
		return result.Err[Object](fmt.Errorf("object of type %T does not contain member %s", obj, memberName))
	} else {
		return result.Ok(NewLValue(obj, member))
	}
}
