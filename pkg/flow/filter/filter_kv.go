package filter

import (
	"errors"
	"fmt"

	"github.com/squizzling/guppy/pkg/interpreter/ffi"
	"github.com/squizzling/guppy/pkg/interpreter/ftypes"
	"github.com/squizzling/guppy/pkg/interpreter/itypes"
	"github.com/squizzling/guppy/pkg/interpreter/primitive"
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
	Terms        *primitive.ObjectTuple                         `ffi:"terms,star"`
	MatchMissing ftypes.ThingOrMissing[*primitive.ObjectBool]   `ffi:"match_missing,kw"`
}

func NewFFIFilter() itypes.FlowCall {
	return ffi.NewFFI(ffiFilter{
		Term:         ftypes.NewThingOrMissingNone[*primitive.ObjectString](),
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
		return nil, fmt.Errorf("ffiFilter.callString: param `term` for ffiFilter.Term is missing, expecting *primitive.ObjectString")
	}

	var terms []string

	terms = append(terms, f.Term.Thing.Value)

	for _, objTerm := range f.Terms.Items {
		if term, ok := objTerm.(*primitive.ObjectString); ok {
			terms = append(terms, term.Value)
		} else {
			// TODO: Can we push this to FFI? Or have FFI generate the error string for us?  It would greatly enable
			//       consistent errors.
			return nil, fmt.Errorf("ffiFilter.callString: expecting *primitive.ObjectString, got %T", objTerm)
		}
	}

	matchMissing := f.MatchMissing.Thing != nil && f.MatchMissing.Thing.Value
	return NewFilterKeyValue(prototypeFilter, f.Field.String.Value, terms, matchMissing), nil
}

func (f ffiFilter) callDict(i itypes.Interpreter) (itypes.Object, error) {
	if f.Term.Missing == nil {
		return nil, errors.New("cannot specify additional parameters (`term`) to dictionary filters")
	}
	if len(f.Terms.Items) > 0 {
		return nil, errors.New("cannot specify additional parameters (`terms`) to dictionary filters")
	}
	if f.MatchMissing.Missing == nil {
		return nil, errors.New("cannot specify additional parameters (`match_missing`) to dictionary filters")
	}
	if len(f.Field.Dict.Items) == 0 {
		return nil, errors.New("cannot create filter from an empty dictionary")
	}

	var filter Filter
	for _, kv := range f.Field.Dict.Items {
		if key, ok := kv.Key.(*primitive.ObjectString); !ok {
			panic("not string")
		} else {
			var items []string
			var err error
			switch value := kv.Value.(type) {
			case *primitive.ObjectString:
				items = []string{value.Value}
				err = nil
			case *primitive.ObjectList:
				items, err = f.resolveTerms(value.Items)
			case *primitive.ObjectTuple:
				items, err = f.resolveTerms(value.Items)
			default:
			}
			if err != nil {
				return nil, err
			} else {
				nextFilter := NewFilterKeyValue(prototypeFilter, key.Value, items, false)
				if filter != nil {
					filter = newFilterAnd(filter, nextFilter)
				} else {
					filter = nextFilter
				}
			}
		}
	}
	if filter == nil {
		panic("no")
	}
	return filter, nil
}

func (f ffiFilter) resolveTerms(items []itypes.Object) ([]string, error) {
	var terms []string
	for _, item := range items {
		if term, ok := item.(*primitive.ObjectString); !ok {
			return nil, errors.New("not string")
		} else {
			terms = append(terms, term.Value)
		}
	}
	return terms, nil
}

func (fkv *FilterKeyValue) Repr() string {
	return repr(fkv)
}
