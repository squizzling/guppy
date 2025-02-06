package interpreter

import (
	"fmt"

	"guppy/internal/parser/ast"
)

func (i *Interpreter) VisitStatementAssert(sa ast.StatementAssert) any {
	defer i.trace()()

	panic("StatementAssert")
}

func (i *Interpreter) VisitStatementDecorated(sd ast.StatementDecorated) any {
	defer i.trace()()

	panic("StatementDecorated")
}

func (i *Interpreter) VisitStatementExpression(se ast.StatementExpression) any {
	defer i.trace()()

	resultValues := r(se.Expr.Accept(i))
	if !resultValues.Ok() {
		return resultValues.Err()
	}
	values := resultValues.Value().(*ObjectList)

	if len(se.Assign) == 0 { // Do nothing
		return nil
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
			return fmt.Errorf("assigning invalid count (assign %d, return %d)", len(se.Assign), len(values.Items))
		}

		i.Scope.DeclareSet(se.Assign[0], values)
	} else {
		for idx := 0; idx < len(values.Items); idx++ {
			i.Scope.DeclareSet(se.Assign[idx], values.Items[idx])
		}
	}
	return nil
}

func (i *Interpreter) VisitStatementFunction(sf ast.StatementFunction) any {
	defer i.trace()()

	panic("StatementFunction")
}

func (i *Interpreter) VisitStatementIf(si ast.StatementIf) any {
	defer i.trace()()

	panic("StatementIf")
}

func (i *Interpreter) VisitStatementImport(si ast.StatementImport) any {
	defer i.trace()()

	panic("StatementImport")
}

func (i *Interpreter) VisitStatementList(sl ast.StatementList) any {
	defer i.trace()()

	for _, stmt := range sl.Statements {
		err, ok := stmt.Accept(i).(error)
		if ok && err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) VisitStatementProgram(sp ast.StatementProgram) any {
	defer i.trace()()

	return sp.Statements.Accept(i)
}

func (i *Interpreter) VisitStatementReturn(sr ast.StatementReturn) any {
	defer i.trace()()

	panic("StatementReturn")
}
