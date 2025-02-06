package ast

import (
	"fmt"
	"strings"
)

type DebugWriter struct {
	depth int
}

func (dw DebugWriter) p() string {
	return strings.Repeat(" ", 2*dw.depth)
}

func (dw *DebugWriter) i() {
	dw.depth++
}

func (dw *DebugWriter) o() {
	dw.depth--
}

func (dw DebugWriter) VisitDataArgument(da DataArgument) any {
	_s := "DataArgument(\n"
	dw.i()
	_s += dw.p() + "Assign: string(" + da.Assign + ")\n"
	if da.Expr != nil {
		_s += dw.p() + "Expr: " + da.Expr.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementProgram(sp StatementProgram) any {
	_s := "StatementProgram(\n"
	dw.i()
	_s += dw.p() + "Statements: " + sp.Statements.Accept(dw).(string) // IsConcrete
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementExpression(se StatementExpression) any {
	_s := "StatementExpression(\n"
	dw.i()
	if se.Assign == nil {
		_s += dw.p() + "Assign: nil\n"
	} else if len(se.Assign) == 0 {
		_s += dw.p() + "Assign: []\n"
	} else {
		_s += dw.p() + "Assign: [\n"
		dw.i()
		for _, _r := range se.Assign {
			_s += dw.p() + _r + "\n" // []string
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	_s += dw.p() + "Expr: " + se.Expr.Accept(dw).(string) // IsConcrete
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementReturn(sr StatementReturn) any {
	_s := "StatementReturn(\n"
	dw.i()
	if sr.Expr != nil {
		_s += dw.p() + "Expr: " + sr.Expr.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementImport(si StatementImport) any {
	_s := "StatementImport(\n"
	dw.i()
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementAssert(sa StatementAssert) any {
	_s := "StatementAssert(\n"
	dw.i()
	if sa.Test != nil {
		_s += dw.p() + "Test: " + sa.Test.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Test: nil\n"
	}
	if sa.Raise != nil {
		_s += dw.p() + "Raise: " + sa.Raise.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Raise: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementIf(si StatementIf) any {
	_s := "StatementIf(\n"
	dw.i()
	if si.Condition == nil {
		_s += dw.p() + "Condition: nil\n"
	} else if len(si.Condition) == 0 {
		_s += dw.p() + "Condition: []\n"
	} else {
		_s += dw.p() + "Condition: [\n"
		dw.i()
		for _, _r := range si.Condition {
			_s += dw.p() + _r.Accept(dw).(string) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	if si.Statement == nil {
		_s += dw.p() + "Statement: nil\n"
	} else if len(si.Statement) == 0 {
		_s += dw.p() + "Statement: []\n"
	} else {
		_s += dw.p() + "Statement: [\n"
		dw.i()
		for _, _r := range si.Statement {
			_s += dw.p() + _r.Accept(dw).(string) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	_s += dw.p() + "StatementElse: " + si.StatementElse.Accept(dw).(string) // IsConcrete
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementFunction(sf StatementFunction) any {
	_s := "StatementFunction(\n"
	dw.i()
	_s += dw.p() + "Token: string(" + sf.Token + ")\n"
	_s += dw.p() + "Body: " + sf.Body.Accept(dw).(string) // IsConcrete
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementDecorated(sd StatementDecorated) any {
	_s := "StatementDecorated(\n"
	dw.i()
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitStatementList(sl StatementList) any {
	_s := "StatementList(\n"
	dw.i()
	if sl.Statements == nil {
		_s += dw.p() + "Statements: nil\n"
	} else if len(sl.Statements) == 0 {
		_s += dw.p() + "Statements: []\n"
	} else {
		_s += dw.p() + "Statements: [\n"
		dw.i()
		for _, _r := range sl.Statements {
			_s += dw.p() + _r.Accept(dw).(string) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionList(el ExpressionList) any {
	_s := "ExpressionList(\n"
	dw.i()
	if el.Expressions == nil {
		_s += dw.p() + "Expressions: nil\n"
	} else if len(el.Expressions) == 0 {
		_s += dw.p() + "Expressions: []\n"
	} else {
		_s += dw.p() + "Expressions: [\n"
		dw.i()
		for _, _r := range el.Expressions {
			_s += dw.p() + _r.Accept(dw).(string) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	_s += dw.p() + "Tuple: bool(" + fmt.Sprintf("%t", el.Tuple) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionBinary(eb ExpressionBinary) any {
	_s := "ExpressionBinary(\n"
	dw.i()
	if eb.Left != nil {
		_s += dw.p() + "Left: " + eb.Left.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Left: nil\n"
	}
	// TODO: 1 Op tokenizer.Token
	_s += dw.p() + fmt.Sprintf("Op: %T(%v)\n", eb.Op, eb.Op)
	if eb.Right != nil {
		_s += dw.p() + "Right: " + eb.Right.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionGrouping(eg ExpressionGrouping) any {
	_s := "ExpressionGrouping(\n"
	dw.i()
	if eg.Expr != nil {
		_s += dw.p() + "Expr: " + eg.Expr.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionLiteral(el ExpressionLiteral) any {
	_s := "ExpressionLiteral(\n"
	dw.i()
	// TODO: 0 Value any
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", el.Value, el.Value)
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionLogical(el ExpressionLogical) any {
	_s := "ExpressionLogical(\n"
	dw.i()
	if el.Left != nil {
		_s += dw.p() + "Left: " + el.Left.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Left: nil\n"
	}
	// TODO: 1 Op tokenizer.Token
	_s += dw.p() + fmt.Sprintf("Op: %T(%v)\n", el.Op, el.Op)
	if el.Right != nil {
		_s += dw.p() + "Right: " + el.Right.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionUnary(eu ExpressionUnary) any {
	_s := "ExpressionUnary(\n"
	dw.i()
	// TODO: 0 Op tokenizer.TokenType
	_s += dw.p() + fmt.Sprintf("Op: %T(%v)\n", eu.Op, eu.Op)
	if eu.Right != nil {
		_s += dw.p() + "Right: " + eu.Right.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionVariable(ev ExpressionVariable) any {
	_s := "ExpressionVariable(\n"
	dw.i()
	_s += dw.p() + "Identifier: string(" + ev.Identifier + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionMember(em ExpressionMember) any {
	_s := "ExpressionMember(\n"
	dw.i()
	if em.Expr != nil {
		_s += dw.p() + "Expr: " + em.Expr.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	_s += dw.p() + "Identifier: string(" + em.Identifier + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionSubscript(es ExpressionSubscript) any {
	_s := "ExpressionSubscript(\n"
	dw.i()
	if es.Expr != nil {
		_s += dw.p() + "Expr: " + es.Expr.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	_s += dw.p() + "Identifier: string(" + es.Identifier + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}

func (dw DebugWriter) VisitExpressionCall(ec ExpressionCall) any {
	_s := "ExpressionCall(\n"
	dw.i()
	if ec.Expr != nil {
		_s += dw.p() + "Expr: " + ec.Expr.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	if ec.Args == nil {
		_s += dw.p() + "Args: nil\n"
	} else if len(ec.Args) == 0 {
		_s += dw.p() + "Args: []\n"
	} else {
		_s += dw.p() + "Args: [\n"
		dw.i()
		for _, _r := range ec.Args {
			_s += dw.p() + _r.Accept(dw).(string) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	if ec.StarArgs != nil {
		_s += dw.p() + "StarArgs: " + ec.StarArgs.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "StarArgs: nil\n"
	}
	if ec.KeywordArgs != nil {
		_s += dw.p() + "KeywordArgs: " + ec.KeywordArgs.Accept(dw).(string) // IsInterface
	} else {
		_s += dw.p() + "KeywordArgs: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s
}
