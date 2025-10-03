package filter

import (
	"fmt"

	"guppy/pkg/interpreter"
)

type FFIFilter struct {
	interpreter.Object
}

func (f FFIFilter) Params(i *interpreter.Interpreter) (*interpreter.Params, error) {
	return &interpreter.Params{
		Params: []interpreter.ParamDef{
			{Name: "field"},
			{Name: "term", Default: interpreter.NewObjectNone()},
		},
		StarParam: "terms",
		KWParams: []interpreter.ParamDef{
			{Name: "match_missing", Default: interpreter.NewObjectBool(false)},
		},
	}, nil
}

func (f FFIFilter) resolveTerms(i *interpreter.Interpreter) ([]string, error) {
	var terms []string

	if objTerm, err := i.Scope.GetArg("term"); err != nil {
		return nil, err
	} else if strTerm, ok := objTerm.(*interpreter.ObjectString); ok {
		terms = append(terms, strTerm.Value)
	} else if _, ok = objTerm.(*interpreter.ObjectNone); !ok {
		return nil, fmt.Errorf("term is not *interpreter.ObjectString or *interpreter.ObjectNone")
	} else {
		// nothing
	}

	if v, err := i.Scope.GetArg("terms"); err != nil {
		return nil, err
	} else {
		switch v := v.(type) {
		case *interpreter.ObjectTuple:
			for _, o := range v.Items {
				if term, err := i.DoString(o); err != nil {
					return nil, err
				} else {
					terms = append(terms, term)
				}
			}
		default:
			return nil, fmt.Errorf("unhandled term type: %T", v)
		}
		return terms, nil
	}
}

func (f FFIFilter) Call(i *interpreter.Interpreter) (interpreter.Object, error) {
	if term, err := interpreter.ArgAsString(i, "field"); err != nil {
		return nil, err
	} else if terms, err := f.resolveTerms(i); err != nil {
		return nil, err
	} else if matchMissing, err := interpreter.ArgAsBool(i, "match_missing"); err != nil {
		return nil, err
	} else {
		return NewKV(term, terms, matchMissing), nil
	}
}

type kv struct {
	interpreter.Object

	key          string
	values       []string
	matchMissing bool
}

func NewKV(key string, values []string, matchMissing bool) Filter {
	return &kv{
		Object:       newFilterObject(),
		key:          key,
		values:       values,
		matchMissing: matchMissing,
	}
}

func (fkv *kv) RenderFilter() string {
	term := "*("
	for _, v := range fkv.values {
		term = term + "'" + v + "'"
	}
	term += ")"
	matchMissing := ""
	if fkv.matchMissing {
		matchMissing = ", match_missing=True"
	}
	return fmt.Sprintf("filter('%s', %s%s)", fkv.key, term, matchMissing)
}
