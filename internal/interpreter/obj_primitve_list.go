package interpreter

import (
	"fmt"
	"strings"
)

// TODO: Proper interface
type ObjectList struct {
	Object

	Items []Object
}

func NewObjectList(items ...Object) Object {
	return &ObjectList{
		Object: NewObject(map[string]Object{
			"__add__": methodListAdd{Object: NewObject(nil)},
		}),
		Items: items,
	}
}

func (ol *ObjectList) Repr() string {
	var sb strings.Builder
	sb.WriteString("list(")
	for idx, item := range ol.Items {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(Repr(item))
	}
	sb.WriteString(")")
	return sb.String()
}

type methodListAdd struct {
	Object
}

func (mla methodListAdd) Params(i *Interpreter) (*Params, error) {
	return BinaryParams, nil
}

func (mla methodListAdd) Call(i *Interpreter) (Object, error) {
	if self, err := ArgAs[*ObjectList](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.Scope.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right := right.(type) {
		case *ObjectList:
			items := make([]Object, 0, len(self.Items)+len(right.Items))
			items = append(items, self.Items...)
			items = append(items, right.Items...)
			return NewObjectList(items...), nil
		default:
			return nil, fmt.Errorf("methodListAdd: unknown type %T", right)
		}
	}
}
