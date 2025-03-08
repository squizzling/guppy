package interpreter

import (
	"fmt"

	"guppy/internal/parser/ast"
)

func (i *Interpreter) VisitStatementAssert(sa ast.StatementAssert) (_ any, errOut error) {
	defer i.trace()(&errOut)

	panic("StatementAssert")
}

func (i *Interpreter) VisitStatementDecorated(sd ast.StatementDecorated) (_ any, errOut error) {
	defer i.trace()(&errOut)

	panic("StatementDecorated")
}

func (i *Interpreter) VisitStatementExpression(se ast.StatementExpression) (_ any, errOut error) {
	defer i.trace()(&errOut)

	valuesAny, err := se.Expr.Accept(i)
	if err != nil {
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
		for idx := 0; idx < len(values.Items); idx++ {
			if err := i.Scope.Set(se.Assign[idx], values.Items[idx]); err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (i *Interpreter) VisitStatementFunction(sf ast.StatementFunction) (_ any, errOut error) {
	defer i.trace()(&errOut)

	panic("StatementFunction")
}

func (i *Interpreter) VisitStatementIf(si ast.StatementIf) (_ any, errOut error) {
	defer i.trace()(&errOut)

	panic("StatementIf")
}

func (i *Interpreter) VisitStatementImportFrom(sif ast.StatementImportFrom) (_ any, errOut error) {
	defer i.trace()(&errOut)

	panic("StatementImportFrom")
}

func (i *Interpreter) VisitStatementImportFromStar(sifs ast.StatementImportFromStar) (_ any, errOut error) {
	defer i.trace()(&errOut)

	panic("StatementImportFromStar")
}

func (i *Interpreter) VisitStatementImportNames(sif ast.StatementImportNames) (_ any, errOut error) {
	defer i.trace()(&errOut)

	panic("StatementImportNames")
}

func (i *Interpreter) VisitStatementList(sl ast.StatementList) (_ any, errOut error) {
	defer i.trace()(&errOut)

	for _, stmt := range sl.Statements {
		_, err := stmt.Accept(i)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (i *Interpreter) VisitStatementProgram(sp ast.StatementProgram) (_ any, errOut error) {
	defer i.trace()(&errOut)

	return sp.Statements.Accept(i)
}

func (i *Interpreter) VisitStatementReturn(sr ast.StatementReturn) (_ any, errOut error) {
	defer i.trace()(&errOut)

	panic("StatementReturn")
}
