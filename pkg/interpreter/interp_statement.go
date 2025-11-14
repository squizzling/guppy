package interpreter

import (
	"fmt"

	"guppy/pkg/interpreter/deferred"
	"guppy/pkg/interpreter/itypes"
	"guppy/pkg/interpreter/primitive"
	"guppy/pkg/parser/ast"
)

func (i *interpreter) VisitStatementAssert(sa ast.StatementAssert) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	if test, err := r(sa.Test.Accept(i)); err != nil {
		return nil, err
	} else if od, isDeferred := test.(*deferred.ObjectDeferred); isDeferred {
		i.Scope.DeferAnonymous(deferred.NewObjectDeferred(sa.Test, od.Desired...).(*deferred.ObjectDeferred))
		return nil, nil
	} else if t, err := isTruthy(test); err != nil {
		return nil, err
	} else if !t {
		panic("Assertion failed")
	} else {
		return nil, nil
	}
}

func (i *interpreter) VisitStatementDecorated(sd ast.StatementDecorated) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementDecorated")
}

func (i *interpreter) VisitStatementExpression(se ast.StatementExpression) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	raw, err := se.Expr.Accept(i)
	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, nil
	}

	valuesAny, isObject := raw.(itypes.Object)
	if !isObject {
		return nil, fmt.Errorf("assignment received %T not Object", raw)
	}

	// If the result is deferred, save it for later
	if od, ok := valuesAny.(*deferred.ObjectDeferred); ok {
		if len(se.Assign) == 0 {
			// TODO: Figure out the type we want here.
			i.Scope.DeferAnonymous(deferred.NewObjectDeferred(se.Expr, od.Desired...).(*deferred.ObjectDeferred))
			return nil, nil
		}

		err := i.Scope.SetDefers(se.Assign, deferred.NewObjectDeferred(se.Expr, od.Desired...).(*deferred.ObjectDeferred))
		return nil, err
	}

	if len(se.Assign) == 0 { // Do nothing
		return nil, nil
	}

	if len(se.Assign) == 1 { // Always assign it, never unpack it
		if err := i.Scope.Set(se.Assign[0], valuesAny); err != nil {
			return nil, err
		}
		return nil, nil
	}

	var values []itypes.Object
	switch valuesAny := valuesAny.(type) {
	case *primitive.ObjectList:
		values = valuesAny.Items
	case *primitive.ObjectTuple:
		values = valuesAny.Items
	case itypes.Object:
		values = []itypes.Object{valuesAny}
	default:
		return nil, fmt.Errorf("assigning from %T not *ObjectList, *ObjectTuple, or Object", valuesAny)
	}

	// Signalflow grammar doesn't do arbitrary tuple unpacking.  ie, it can handle:
	//
	// > a, b = 1, 2
	// > a, b = x()
	// > a, b = c  # where c = (1, 2)
	//
	// But it can't handle:
	//
	// > (a, b), c = (1, 2), 3
	// > (a, b) = 1, 2
	//
	// TODO: The grammar can handle it, but it may not be supported in reality.
	if len(values) != len(se.Assign) {
		if len(se.Assign) != 1 {
			return nil, fmt.Errorf("assigning invalid count (assign %d, return %d)", len(se.Assign), len(values))
		}

		if err := i.Scope.Set(se.Assign[0], valuesAny); err != nil {
			return nil, err
		}
	} else {
		for idx, value := range values {
			if err := i.Scope.Set(se.Assign[idx], value); err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (i *interpreter) VisitStatementFunction(sf ast.StatementFunction) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	params := &itypes.Params{}
	for _, param := range sf.Params.List {
		if param.StarArg {
			params.StarParam = param.Name
		} else if param.KeywordArg {
			params.KWParam = param.Name
		} else {
			var def itypes.Object
			if param.Default != nil {
				var ok bool
				if defRaw, err := param.Default.Accept(i); err != nil {
					return nil, err
				} else if def, ok = defRaw.(itypes.Object); !ok {
					return nil, fmt.Errorf("default is a %T, not an Object", defRaw)
				}
			}
			if params.StarParam == "" {
				params.Params = append(params.Params, itypes.ParamDef{
					Name:    param.Name,
					Default: def,
				})
			} else {
				params.KWParams = append(params.KWParams, itypes.ParamDef{
					Name:    param.Name,
					Default: def,
				})
			}
		}
	}

	i.Debug("defining function: %s", sf.Token)
	params.Dump(i)

	err := i.Scope.Set(sf.Token, NewObjectFunction(sf.Token, params, i.Scope, sf.Body))
	return nil, err
}

func (i *interpreter) VisitStatementIf(si ast.StatementIf) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	for idx, exprCond := range si.Condition {
		if cond, err := r(exprCond.Accept(i)); err != nil {
			return nil, err
		} else if cond, err := isTruthy(cond); err != nil {
			return nil, err
		} else if cond {
			return si.Statement[idx].Accept(i)
		}
	}
	if si.StatementElse != nil {
		return si.StatementElse.Accept(i)
	} else {
		return nil, nil
	}
}

func (i *interpreter) VisitStatementImportFrom(sif ast.StatementImportFrom) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementImportFrom")
}

func (i *interpreter) VisitStatementImportFromStar(sifs ast.StatementImportFromStar) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementImportFromStar")
}

func (i *interpreter) VisitStatementImportNames(sif ast.StatementImportNames) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementImportNames")
}

func (i *interpreter) VisitStatementList(sl ast.StatementList) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	for _, stmt := range sl.Statements {
		ret, err := stmt.Accept(i)
		if err != nil {
			return nil, err
		}
		if ret != nil {
			return ret, nil
		}
	}
	return nil, nil
}

func (i *interpreter) VisitStatementProgram(sp ast.StatementProgram) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	i.pushScope()
	defer i.popScope()
	if _, err := sp.Statements.Accept(i); err != nil {
		return nil, err
	} else if err := i.Scope.resolveDeferred(i); err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *interpreter) VisitStatementReturn(sr ast.StatementReturn) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	if sr.Expr == nil {
		return primitive.NewObjectNone(), nil
	}

	return sr.Expr.Accept(i)
}
