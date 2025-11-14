package interpreter

import (
	"fmt"
	"strings"

	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

// TODO: Proper interface
type ObjectList struct {
	itypes.Object

	Items []itypes.Object
}

func NewObjectList(items ...itypes.Object) *ObjectList {
	return &ObjectList{
		Object: itypes.NewObject(map[string]itypes.Object{
			"__add__":       methodListAdd{Object: itypes.NewObject(nil)},
			"__subscript__": methodListSubscript{Object: itypes.NewObject(nil)},
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
		sb.WriteString(itypes.Repr(item))
	}
	sb.WriteString(")")
	return sb.String()
}

type methodListAdd struct {
	itypes.Object
}

func (mla methodListAdd) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return itypes.BinaryParams, nil
}

func (mla methodListAdd) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[*ObjectList](i, "self"); err != nil {
		return nil, err
	} else if right, err := i.GetArg("right"); err != nil {
		return nil, err
	} else {
		switch right := right.(type) {
		case *ObjectList:
			items := make([]itypes.Object, 0, len(self.Items)+len(right.Items))
			items = append(items, self.Items...)
			items = append(items, right.Items...)
			return NewObjectList(items...), nil
		default:
			return nil, fmt.Errorf("methodListAdd: unknown type %T", right)
		}
	}
}

type methodListSubscript struct {
	itypes.Object
}

func (mls methodListSubscript) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "start"},
		},
	}, nil
}

func (mls methodListSubscript) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[*ObjectList](i, "self"); err != nil {
		return nil, err
	} else if start, err := itypes.ArgAs[*primitive.ObjectInt](i, "start"); err != nil {
		return nil, err
	} else if len(self.Items) < start.Value+1 || start.Value < 0 {
		// TODO: Does flow support x[-1] for last item?
		return nil, fmt.Errorf("index out of range (%d < %d)", len(self.Items), start.Value+1)
	} else {
		return self.Items[start.Value], nil
	}
}
