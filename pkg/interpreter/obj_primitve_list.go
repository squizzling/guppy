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

func NewObjectList(items ...Object) *ObjectList {
	return &ObjectList{
		Object: NewObject(map[string]Object{
			"__add__":       methodListAdd{Object: NewObject(nil)},
			"__subscript__": methodListSubscript{Object: NewObject(nil)},
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

type methodListSubscript struct {
	Object
}

func (mls methodListSubscript) Params(i *Interpreter) (*Params, error) {
	return &Params{
		Params: []ParamDef{
			{Name: "self"},
			{Name: "start"},
		},
	}, nil
}

func (mls methodListSubscript) Call(i *Interpreter) (Object, error) {
	if self, err := ArgAs[*ObjectList](i, "self"); err != nil {
		return nil, err
	} else if start, err := ArgAs[*ObjectInt](i, "start"); err != nil {
		return nil, err
	} else if len(self.Items) < start.Value+1 || start.Value < 0 {
		// TODO: Does flow support x[-1] for last item?
		return nil, fmt.Errorf("index out of range (%d < %d)", len(self.Items), start.Value+1)
	} else {
		return self.Items[start.Value], nil
	}
}
