package interpreter

import (
	"fmt"
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

func (od *ObjectDict) get(key Object, def Object) (Object, error) {
	if len(od.items) > 0 {
		return nil, fmt.Errorf("can't read from dict with data")
	}
	return def, nil
}

var _ = FlowCall(methodDictGet{})

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
		return self.get(key, def)
	}
}
