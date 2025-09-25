package interpreter

import (
	"fmt"

	"guppy/internal/parser/ast"
)

func (i *Interpreter) VisitStatementAssert(sa ast.StatementAssert) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	if test, err := r(sa.Test.Accept(i)); err != nil {
		return nil, err
	} else if od, isDeferred := test.(*ObjectDeferred); isDeferred {
		i.Scope.DeferAnonymous(NewObjectDeferred(sa.Test, od.desired...).(*ObjectDeferred))
		return nil, nil
	} else if t, err := isTruthy(test); err != nil {
		return nil, err
	} else if !t {
		panic("Assertion failed")
	} else {
		return nil, nil
	}
}

func (i *Interpreter) VisitStatementDecorated(sd ast.StatementDecorated) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementDecorated")
}

func (i *Interpreter) VisitStatementExpression(se ast.StatementExpression) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	valuesAny, err := r(se.Expr.Accept(i))
	if err != nil {
		return nil, err
	}

	// If the result is deferred, save it for later
	if od, ok := valuesAny.(*ObjectDeferred); ok {
		if len(se.Assign) == 0 {
			// TODO: Figure out the type we want here.
			i.Scope.DeferAnonymous(NewObjectDeferred(se.Expr, od.desired...).(*ObjectDeferred))
			return nil, nil
		}

		err := i.Scope.SetDefers(se.Assign, NewObjectDeferred(se.Expr, od.desired...).(*ObjectDeferred))
		return nil, err
	}

	// TODO: Clean up this mess.  Originally we returned a list of length 1 from functions,
	//       but now we don't.  This turns it back in to a list of length 1, but that's also
	//       kinda dumb.
	if len(se.Assign) == 0 { // Do nothing
		return nil, nil
	}

	var values []Object
	switch valuesAny := valuesAny.(type) {
	case *ObjectList:
		values = valuesAny.Items
	case *ObjectTuple:
		values = valuesAny.Items
	case Object:
		values = []Object{valuesAny}
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

func (i *Interpreter) VisitStatementFunction(sf ast.StatementFunction) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	params := &Params{}
	for _, param := range sf.Params.List {
		if param.StarArg {
			params.StarParam = param.Name
		} else if param.KeywordArg {
			params.KWParam = param.Name
		} else {
			var def Object
			if param.Default != nil {
				var ok bool
				if defRaw, err := param.Default.Accept(i); err != nil {
					return nil, err
				} else if def, ok = defRaw.(Object); !ok {
					return nil, fmt.Errorf("default is a %T, not an Object", defRaw)
				}
			}
			if params.StarParam == "" {
				params.Params = append(params.Params, ParamDef{
					Name:    param.Name,
					Default: def,
				})
			} else {
				params.KWParams = append(params.KWParams, ParamDef{
					Name:    param.Name,
					Default: def,
				})
			}
		}
	}

	err := i.Scope.Set(sf.Token, NewObjectFunction(sf.Token, params, i.Scope, sf.Body))
	return nil, err
}

func (i *Interpreter) VisitStatementIf(si ast.StatementIf) (returnValue any, errOut error) {
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

func (i *Interpreter) VisitStatementImportFrom(sif ast.StatementImportFrom) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementImportFrom")
}

func (i *Interpreter) VisitStatementImportFromStar(sifs ast.StatementImportFromStar) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementImportFromStar")
}

func (i *Interpreter) VisitStatementImportNames(sif ast.StatementImportNames) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementImportNames")
}

func (i *Interpreter) VisitStatementList(sl ast.StatementList) (returnValue any, errOut error) {
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
	return NewObjectNone(), nil
}

func (i *Interpreter) VisitStatementProgram(sp ast.StatementProgram) (returnValue any, errOut error) {
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

func (i *Interpreter) VisitStatementReturn(sr ast.StatementReturn) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	if sr.Expr == nil {
		return NewObjectNone(), nil
	}

	return sr.Expr.Accept(i)
}
