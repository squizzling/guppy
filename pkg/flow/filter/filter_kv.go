package filter

import (
	"fmt"
	"strings"

	"guppy/pkg/interpreter/ffi"
	"guppy/pkg/interpreter/ftypes"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
)

type ffiFilter struct {
	/**
	f = filter()                                              #TypeError: <'filter function'> missing required argument 'field'
	f = filter(None)                                          #TypeError: <'filter function'> unsupported type 'none' object for argument 'field'
	f = filter(5)                                             #TypeError: <'filter function'> unsupported type 'long' object for argument 'field'
	f = filter('a')                                           #ValueError: missing required value for filter term
	f = filter('a', None)                                     #TypeError: <'filter function'> unsupported type 'none' object for argument 'term'
	f = filter('a', 'b', None)                                #TypeError: <'filter function'> unsupported type 'none' object for argument 'terms'
	f = filter('a', 'b', 5)                                   #TypeError: <'filter function'> unsupported type 'long' object for argument 'terms'
	f = filter('a', 'b', terms='c')                           #: class sf.analytics.program.parsing.values.ScalarValue cannot be cast to class sf.analytics.program.parsing.values.TupleValue (sf.analytics.program.parsing.values.ScalarValue and sf.analytics.program.parsing.values.TupleValue are in unnamed module of loader 'app')
	f = filter('a', 'b', terms=5)                             #TypeError: <'filter function'> unsupported type 'long' object for argument 'terms'
	f = filter('a', 'b', terms=['c'])                         #TypeError: <'filter function'> unsupported type 'list' object for argument 'terms'
	f = filter('a', 'b', terms=('c',))                        #TypeError: <'filter function'> unsupported type 'tuple' object for argument 'terms'
	f = filter('a', 'b', ['c'])                               #TypeError: <'filter function'> unsupported type 'list' object for argument 'terms'
	f = filter('a', 'b', *['c'])                              #Ok
	f = filter('a', 'b', 'c', match_missing=None)             #TypeError: <'filter function'> unsupported type 'none' object for argument 'match_missing'
	f = filter('a', 'b', 'c', match_missing=5)                #Ok
	f = filter('a', 'b', 'c', match_missing=5.5)              #TypeError: <'filter function'> unsupported type 'double' object for argument 'match_missing'
	f = filter('a', 'b', 'c', match_missing='d')              #Ok
	f = filter('a', 'b', 'c', match_missing=True)             #Ok
	f = filter('a', 'b', 'c', match_missing=[])               #TypeError: <'filter function'> unsupported type 'list' object for argument 'match_missing'
	f = filter('a', 'b', 'c', match_missing=())               #TypeError: <'filter function'> unsupported type 'tuple' object for argument 'match_missing'
	f = filter('a', 'b', 'c', match_missing=const(1))         #TypeError: <'filter function'> unsupported type <stream of DOUBLE> for argument 'match_missing'
	f = filter('a', 'b', 'c', match_missing=filter('d', 'e')) #TypeError: <'filter function'> unsupported type 'filter' object for argument 'match_missing'
	f = filter({})                                            #: cannot create filter from an empty dictionary
	f = filter({'a': 'b'})                                    #Ok
	f = filter(field={'a': 'b'})                              #Ok
	f = filter({'a': 'b'}, 'c')                               #TypeError: cannot specify additional parameters to dictionary filters
	f = filter({'a': 'b'}, match_missing=True)                #TypeError: cannot specify additional parameters to dictionary filters
	f = filter({'a': ['b']})                                  #Ok
	f = filter({'a': ('b')})                                  #Ok
	f = filter({'a': ('b')}, match_missing=True)              #TypeError: cannot specify additional parameters to dictionary filters
	f = filter({'a': ('b')}, c=True)                          #TypeError: <'filter function'> got an unexpected keyword argument 'c'
	f = filter({'a': 'b'}, c=True)                            #TypeError: <'filter function'> got an unexpected keyword argument 'c'
	f = filter({'a': 'b'}, None)                              #TypeError: <'filter function'> unsupported type 'none' object for argument 'term'
	f = filter({'a': 'b'}, 'c', None)                         #TypeError: <'filter function'> unsupported type 'none' object for argument 'terms'
	f = filter({'a': 'b'}, 'c', 'd')                          #TypeError: cannot specify additional parameters to dictionary filters

	From this we can determine:
	- `field` is the only required parameter.  It can be a string or a dict
	- If `field` is a string, then term must be provided
	- If `field` is a dict, type checking still applies, but everything else must be Missing
	- We can't use any defaults, we need to manually resolve (or enforce absence) via Missing
	*/
	Field struct {
		String *primitive.ObjectString
		Dict   *primitive.ObjectDict
	} `ffi:"field"`
	Term         ftypes.ThingOrMissing[*primitive.ObjectString] `ffi:"term"`
	Terms        ftypes.ThingOrMissing[*primitive.ObjectTuple]  `ffi:"terms,star"`
	MatchMissing ftypes.ThingOrMissing[*primitive.ObjectBool]   `ffi:"match_missing,kw"`
}

func NewFFIFilter() itypes.FlowCall {
	return ffi.NewFFI(ffiFilter{
		Term:         ftypes.NewThingOrMissingNone[*primitive.ObjectString](),
		Terms:        ftypes.NewThingOrMissingNone[*primitive.ObjectTuple](),
		MatchMissing: ftypes.NewThingOrMissingNone[*primitive.ObjectBool](),
	})
}

func (f ffiFilter) Call(i itypes.Interpreter) (itypes.Object, error) {
	if f.Field.String != nil {
		return f.callString(i)
	} else {
		return f.callDict(i)
	}
}

func (f ffiFilter) callString(i itypes.Interpreter) (itypes.Object, error) {
	if f.Term.Missing != nil {
		panic("term is required")
	}

	var terms []string

	terms = append(terms, f.Term.Thing.Value)

	for _, objTerm := range f.Terms.Thing.Items {
		if term, ok := objTerm.(*primitive.ObjectString); ok {
			terms = append(terms, term.Value)
		} else {
			// TODO: Can we push this to FFI? Or have FFI generate the error string for us?  It would greatly enable
			//       consistent errors.
			return nil, fmt.Errorf("ffiFilter.callString: expecting *primitive.ObjectString, got %T", term)
		}
	}

	matchMissing := f.MatchMissing.Thing != nil && f.MatchMissing.Thing.Value
	return NewKV(f.Field.String.Value, terms, matchMissing), nil
}

func (f ffiFilter) callDict(i itypes.Interpreter) (itypes.Object, error) {
	if f.Term.Missing == nil {
		panic("cannot specify additional parameters to dictionary filters")
	}
	if f.Terms.Missing == nil {
		panic("cannot specify additional parameters to dictionary filters")
	}
	if f.MatchMissing.Missing == nil {
		panic("cannot specify additional parameters to dictionary filters")
	}
	panic("TODO")
}

/*type FFIFilter struct {
	itypes.Object
}

func (f FFIFilter) Params(i itypes.Interpreter) (*itypes.Params, error) {
	return &itypes.Params{
		Params: []itypes.ParamDef{
			{Name: "field"},
			{Name: "term", Default: primitive.NewObjectNone()},
		},
		StarParam: "terms",
		KWParams: []itypes.ParamDef{
			{Name: "match_missing", Default: primitive.NewObjectBool(false)},
		},
	}, nil
}

func (f FFIFilter) resolveTerms(i itypes.Interpreter) ([]string, error) {
	var terms []string

	if objTerm, err := i.GetArg("term"); err != nil {
		return nil, err
	} else if strTerm, ok := objTerm.(*primitive.ObjectString); ok {
		terms = append(terms, strTerm.Value)
	} else if _, ok = objTerm.(*primitive.ObjectNone); !ok {
		return nil, fmt.Errorf("term is not *interpreter.ObjectString or *interpreter.ObjectNone")
	} else {
		// nothing
	}

	if v, err := i.GetArg("terms"); err != nil {
		return nil, err
	} else {
		switch v := v.(type) {
		case *primitive.ObjectTuple:
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

func (f FFIFilter) Call(i itypes.Interpreter) (itypes.Object, error) {
	if term, err := interpreter.ArgAsString(i, "field"); err != nil {
		return nil, err
	} else if terms, err := f.resolveTerms(i); err != nil {
		return nil, err
	} else if matchMissing, err := itypes.ArgAs[*primitive.ObjectBool](i, "match_missing"); err != nil {
		return nil, err
	} else {
		return NewKV(term, terms, matchMissing.Value), nil
	}
}*/

type kv struct {
	itypes.Object

	key          string
	values       []string
	matchMissing bool
}

func NewKV(key string, values []string, matchMissing bool) Filter {
	return &kv{
		Object:       prototypeFilter,
		key:          key,
		values:       values,
		matchMissing: matchMissing,
	}
}

func (fkv *kv) Repr() string {
	var sb strings.Builder
	sb.WriteString("filter('")
	sb.WriteString(fkv.key)
	sb.WriteString("'")
	for _, v := range fkv.values {
		sb.WriteString(", '")
		sb.WriteString(v)
		sb.WriteString("'")
	}

	if fkv.matchMissing {
		sb.WriteString(", match_missing=True")
	}
	sb.WriteString(")")
	return sb.String()
}
