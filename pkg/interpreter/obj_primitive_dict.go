package interpreter

import (
	"errors"
	"fmt"
	"strings"

	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type DictItem struct {
	Key   itypes.Object
	Value itypes.Object
}

type ObjectDict struct {
	itypes.Object

	items []DictItem
}

func NewObjectDict(items []DictItem) itypes.Object {
	return &ObjectDict{
		Object: itypes.NewObject(map[string]itypes.Object{
			"get":           methodDictGet{itypes.NewObject(nil)},
			"__subscript__": methodDictSubscript{itypes.NewObject(nil)},
		}),

		items: items,
	}
}

func NewObjectDictFromMap(items map[string]itypes.Object) itypes.Object {
	var itemList []DictItem
	for key, value := range items {
		itemList = append(itemList, DictItem{
			Key:   primitive.NewObjectString(key),
			Value: value,
		})
	}
	return NewObjectDict(itemList)
}

func (od *ObjectDict) AsMapStringString() (map[string]string, error) {
	m := make(map[string]string)
	for idx, item := range od.items {
		sKey := ""
		sValue := ""
		switch key := item.Key.(type) {
		case *primitive.ObjectString:
			sKey = key.Value
		default:
			return nil, fmt.Errorf("dict idx %d (%s) is %T not *interpreter.ObjectString", idx, itypes.Repr(key), key)
		}
		switch value := item.Value.(type) {
		case *primitive.ObjectString:
			sValue = value.Value
		default:
			return nil, fmt.Errorf("dict idx %d (%s) is %T not *interpreter.ObjectString", idx, sKey, value)
		}
		m[sKey] = sValue
	}
	return m, nil
}

func (od *ObjectDict) tryGet(key itypes.Object, def itypes.Object) (itypes.Object, error) {
	if obj, err := od.mustGet(key); err != nil {
		return nil, err
	} else if obj == nil {
		return def, nil
	} else {
		return obj, nil
	}
}

func (od *ObjectDict) mustGet(key itypes.Object) (itypes.Object, error) {
	if len(od.items) == 0 {
		return nil, nil
	}

	switch key := key.(type) {
	case *primitive.ObjectString:
		return od.mustGetString(key.Value)
	default:
		return nil, fmt.Errorf("requested key is %T", key)
	}
}

func (od *ObjectDict) mustGetString(key string) (itypes.Object, error) {
	for _, item := range od.items {
		if itemKey, ok := item.Key.(*primitive.ObjectString); ok {
			if itemKey.Value == key {
				return item.Value, nil
			}
		}
	}
	return nil, nil
}

func (od *ObjectDict) Repr() string {
	var sb strings.Builder
	sb.WriteString("{")
	for idx, item := range od.items {
		if idx > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(itypes.Repr(item.Key))
		sb.WriteString(": ")
		sb.WriteString(itypes.Repr(item.Value))
	}
	sb.WriteString("}")
	return sb.String()
}

var _ = itypes.FlowCall(methodDictGet{})
var _ = itypes.FlowCall(methodDictSubscript{})

type methodDictGet struct {
	itypes.Object
}

func (mdg methodDictGet) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "key"},
			{Name: "default", Default: primitive.NewObjectNone()},
		},
	}, nil
}

func resolveDictKey(i itypes.Interpreter) (itypes.Object, error) {
	if key, err := i.GetArg("key"); err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

func resolveDictDefault(i itypes.Interpreter) (itypes.Object, error) {
	if key, err := i.GetArg("default"); err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

func (mdg methodDictGet) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[*ObjectDict](i, "self"); err != nil {
		return nil, err
	} else if key, err := resolveDictKey(i); err != nil {
		return nil, err
	} else if def, err := resolveDictDefault(i); err != nil {
		return nil, err
	} else {
		return self.tryGet(key, def)
	}
}

type methodDictSubscript struct {
	itypes.Object
}

func (mds methodDictSubscript) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "self"},
			{Name: "start"},
		},
	}, nil
}

func (mds methodDictSubscript) resolveDictKey(i itypes.Interpreter) (itypes.Object, error) {
	if key, err := i.GetArg("start"); err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

func (mds methodDictSubscript) Call(i itypes.Interpreter) (itypes.Object, error) {
	if self, err := itypes.ArgAs[*ObjectDict](i, "self"); err != nil {
		return nil, err
	} else if start, err := mds.resolveDictKey(i); err != nil {
		return nil, err
	} else if value, err := self.mustGet(start); err != nil {
		return nil, err
	} else if value == nil {
		return nil, errors.New("key not found in dict")
	} else {
		return value, nil
	}
}
