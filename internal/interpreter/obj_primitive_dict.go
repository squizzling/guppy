package interpreter

import (
	"fmt"
	"strings"
)

type DictItem struct {
	Key   Object
	Value Object
}

type ObjectDict struct {
	Object

	items []DictItem
}

func NewObjectDict(items []DictItem) Object {
	return &ObjectDict{
		Object: NewObject(map[string]Object{
			"get": methodDictGet{NewObject(nil)},
			"__subscript__": methodDictSubscript{NewObject(nil)},
		}),

		items: items,
	}
}

func NewObjectDictFromMap(items map[string]Object) Object {
	var itemList []DictItem
	for key, value := range items {
		itemList = append(itemList, DictItem{
			Key:   NewObjectString(key),
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
		case *ObjectString:
			sKey = key.Value
		default:
			return nil, fmt.Errorf("dict idx %d (%s) is %T not *interpreter.ObjectString", idx, Repr(key), key)
		}
		switch value := item.Value.(type) {
		case *ObjectString:
			sValue = value.Value
		default:
			return nil, fmt.Errorf("dict idx %d (%s) is %T not *interpreter.ObjectString", idx, sKey, value)
		}
		m[sKey] = sValue
	}
	return m, nil
}

func (od *ObjectDict) tryGet(key Object, def Object) (Object, error) {
	if obj, err := od.mustGet(key); err != nil {
		return nil, err
	} else if obj == nil {
		return def, nil
	} else {
		return obj, nil
	}
}

func (od *ObjectDict) mustGet(key Object) (Object, error) {
	if len(od.items) > 0 {
		return nil, fmt.Errorf("can't read from dict with data")
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
		sb.WriteString(Repr(item.Key))
		sb.WriteString(": ")
		sb.WriteString(Repr(item.Value))
	}
	sb.WriteString("}")
	return sb.String()
}

var _ = FlowCall(methodDictGet{})
var _ = FlowCall(methodDictSubscript{})

type methodDictGet struct {
	Object
}

func (mdg methodDictGet) Params(i *Interpreter) (*Params, error) {
	return &Params{
		Params: []ParamDef{
			{Name: "self"},
			{Name: "key"},
			{Name: "default", Default: NewObjectNone()},
		},
	}, nil
}

func resolveDictKey(i *Interpreter) (Object, error) {
	if key, err := i.Scope.GetArg("key"); err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

func resolveDictDefault(i *Interpreter) (Object, error) {
	if key, err := i.Scope.GetArg("default"); err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

func (mdg methodDictGet) Call(i *Interpreter) (Object, error) {
	if self, err := ArgAs[*ObjectDict](i, "self"); err != nil {
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
	Object
}

func (mds methodDictSubscript) Params(i *Interpreter) (*Params, error) {
	return &Params{
		Params: []ParamDef{
			{Name: "self"},
			{Name: "start"},
		},
	}, nil
}

func (mds methodDictSubscript) resolveDictKey(i *Interpreter) (Object, error) {
	if key, err := i.Scope.GetArg("start"); err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

func (mds methodDictSubscript) Call(i *Interpreter) (Object, error) {
	if self, err := ArgAs[*ObjectDict](i, "self"); err != nil {
		return nil, err
	} else if start, err := mds.resolveDictKey(i); err != nil {
		return nil, err
	} else {
		return self.mustGet(start)
	}
}
