package interpreter

import (
	"fmt"

	"guppy/internal/parser/ast"
)

func (i *Interpreter) VisitStatementAssert(sa ast.StatementAssert) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementAssert")
}

func (i *Interpreter) VisitStatementDecorated(sd ast.StatementDecorated) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementDecorated")
}

func (i *Interpreter) VisitStatementExpression(se ast.StatementExpression) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	valuesAny, err := se.Expr.Accept(i)
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

	values := valuesAny.(*ObjectList)
	if len(se.Assign) == 0 { // Do nothing
		return nil, nil
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
	if len(values.Items) != len(se.Assign) {
		if len(se.Assign) != 1 {
			return nil, fmt.Errorf("assigning invalid count (assign %d, return %d)", len(se.Assign), len(values.Items))
		}

		if err := i.Scope.Set(se.Assign[0], values); err != nil {
			return nil, err
		}
	} else {
		for idx, value := range values.Items {
			if err := i.Scope.Set(se.Assign[idx], value); err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (i *Interpreter) VisitStatementFunction(sf ast.StatementFunction) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementFunction")
}

func (i *Interpreter) VisitStatementIf(si ast.StatementIf) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementIf")
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
		_, err := stmt.Accept(i)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (i *Interpreter) VisitStatementProgram(sp ast.StatementProgram) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	i.pushScope()
	defer i.popScope()
	if _, err := sp.Statements.Accept(i); err != nil {
		return nil, err
	} else if err := i.resolveDeferred(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *Interpreter) VisitStatementReturn(sr ast.StatementReturn) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("StatementReturn")
}
