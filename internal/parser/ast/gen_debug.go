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

func s(a any, err error) string {
	return a.(string)
}

func (dw DebugWriter) VisitDataArgument(da DataArgument) (any, error) {
	_s := "DataArgument(\n"
	dw.i()
	_s += dw.p() + "Assign: string(" + da.Assign + ")\n"
	if da.Expr != nil {
		_s += dw.p() + "Expr: " + s(da.Expr.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitDataParameterList(dpl DataParameterList) (any, error) {
	_s := "DataParameterList(\n"
	dw.i()
	if dpl.List == nil {
		_s += dw.p() + "List: nil\n"
	} else if len(dpl.List) == 0 {
		_s += dw.p() + "List: []\n"
	} else {
		_s += dw.p() + "List: [\n"
		dw.i()
		for _, _r := range dpl.List {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitDataParameter(dp DataParameter) (any, error) {
	_s := "DataParameter(\n"
	dw.i()
	_s += dw.p() + "Name: string(" + dp.Name + ")\n"
	_s += dw.p() + "Type: string(" + dp.Type + ")\n"
	if dp.Default != nil {
		_s += dw.p() + "Default: " + s(dp.Default.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Default: nil\n"
	}
	_s += dw.p() + "StarArg: bool(" + fmt.Sprintf("%t", dp.StarArg) + ")\n"
	_s += dw.p() + "KeywordArg: bool(" + fmt.Sprintf("%t", dp.KeywordArg) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementProgram(sp StatementProgram) (any, error) {
	_s := "StatementProgram(\n"
	dw.i()
	_s += dw.p() + "Statements: " + s(sp.Statements.Accept(dw)) // IsConcrete
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementExpression(se StatementExpression) (any, error) {
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
	_s += dw.p() + "Expr: " + s(se.Expr.Accept(dw)) // IsConcrete
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementReturn(sr StatementReturn) (any, error) {
	_s := "StatementReturn(\n"
	dw.i()
	if sr.Expr != nil {
		_s += dw.p() + "Expr: " + s(sr.Expr.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementImport(si StatementImport) (any, error) {
	_s := "StatementImport(\n"
	dw.i()
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementAssert(sa StatementAssert) (any, error) {
	_s := "StatementAssert(\n"
	dw.i()
	if sa.Test != nil {
		_s += dw.p() + "Test: " + s(sa.Test.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Test: nil\n"
	}
	if sa.Raise != nil {
		_s += dw.p() + "Raise: " + s(sa.Raise.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Raise: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementIf(si StatementIf) (any, error) {
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
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
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
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	_s += dw.p() + "StatementElse: " + s(si.StatementElse.Accept(dw)) // IsConcrete
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementFunction(sf StatementFunction) (any, error) {
	_s := "StatementFunction(\n"
	dw.i()
	_s += dw.p() + "Token: string(" + sf.Token + ")\n"
	_s += dw.p() + "Params: " + s(sf.Params.Accept(dw)) // IsConcrete
	if sf.Body != nil {
		_s += dw.p() + "Body: " + s(sf.Body.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Body: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementDecorated(sd StatementDecorated) (any, error) {
	_s := "StatementDecorated(\n"
	dw.i()
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitStatementList(sl StatementList) (any, error) {
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
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionList(el ExpressionList) (any, error) {
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
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	_s += dw.p() + "Tuple: bool(" + fmt.Sprintf("%t", el.Tuple) + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionBinary(eb ExpressionBinary) (any, error) {
	_s := "ExpressionBinary(\n"
	dw.i()
	if eb.Left != nil {
		_s += dw.p() + "Left: " + s(eb.Left.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Left: nil\n"
	}
	// TODO: 1 Op tokenizer.Token
	_s += dw.p() + fmt.Sprintf("Op: %T(%v)\n", eb.Op, eb.Op)
	if eb.Right != nil {
		_s += dw.p() + "Right: " + s(eb.Right.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionDict(ed ExpressionDict) (any, error) {
	_s := "ExpressionDict(\n"
	dw.i()
	if ed.Keys == nil {
		_s += dw.p() + "Keys: nil\n"
	} else if len(ed.Keys) == 0 {
		_s += dw.p() + "Keys: []\n"
	} else {
		_s += dw.p() + "Keys: [\n"
		dw.i()
		for _, _r := range ed.Keys {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	if ed.Values == nil {
		_s += dw.p() + "Values: nil\n"
	} else if len(ed.Values) == 0 {
		_s += dw.p() + "Values: []\n"
	} else {
		_s += dw.p() + "Values: [\n"
		dw.i()
		for _, _r := range ed.Values {
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionGrouping(eg ExpressionGrouping) (any, error) {
	_s := "ExpressionGrouping(\n"
	dw.i()
	if eg.Expr != nil {
		_s += dw.p() + "Expr: " + s(eg.Expr.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionLiteral(el ExpressionLiteral) (any, error) {
	_s := "ExpressionLiteral(\n"
	dw.i()
	// TODO: 0 Value any
	_s += dw.p() + fmt.Sprintf("Value: %T(%v)\n", el.Value, el.Value)
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionLogical(el ExpressionLogical) (any, error) {
	_s := "ExpressionLogical(\n"
	dw.i()
	if el.Left != nil {
		_s += dw.p() + "Left: " + s(el.Left.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Left: nil\n"
	}
	// TODO: 1 Op tokenizer.Token
	_s += dw.p() + fmt.Sprintf("Op: %T(%v)\n", el.Op, el.Op)
	if el.Right != nil {
		_s += dw.p() + "Right: " + s(el.Right.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionTernary(et ExpressionTernary) (any, error) {
	_s := "ExpressionTernary(\n"
	dw.i()
	if et.Left != nil {
		_s += dw.p() + "Left: " + s(et.Left.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Left: nil\n"
	}
	if et.Cond != nil {
		_s += dw.p() + "Cond: " + s(et.Cond.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Cond: nil\n"
	}
	if et.Right != nil {
		_s += dw.p() + "Right: " + s(et.Right.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionUnary(eu ExpressionUnary) (any, error) {
	_s := "ExpressionUnary(\n"
	dw.i()
	// TODO: 0 Op tokenizer.TokenType
	_s += dw.p() + fmt.Sprintf("Op: %T(%v)\n", eu.Op, eu.Op)
	if eu.Right != nil {
		_s += dw.p() + "Right: " + s(eu.Right.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Right: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionVariable(ev ExpressionVariable) (any, error) {
	_s := "ExpressionVariable(\n"
	dw.i()
	_s += dw.p() + "Identifier: string(" + ev.Identifier + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionMember(em ExpressionMember) (any, error) {
	_s := "ExpressionMember(\n"
	dw.i()
	if em.Expr != nil {
		_s += dw.p() + "Expr: " + s(em.Expr.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	_s += dw.p() + "Identifier: string(" + em.Identifier + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionSubscript(es ExpressionSubscript) (any, error) {
	_s := "ExpressionSubscript(\n"
	dw.i()
	if es.Expr != nil {
		_s += dw.p() + "Expr: " + s(es.Expr.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "Expr: nil\n"
	}
	_s += dw.p() + "Identifier: string(" + es.Identifier + ")\n"
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}

func (dw DebugWriter) VisitExpressionCall(ec ExpressionCall) (any, error) {
	_s := "ExpressionCall(\n"
	dw.i()
	if ec.Expr != nil {
		_s += dw.p() + "Expr: " + s(ec.Expr.Accept(dw)) // IsInterface
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
			_s += dw.p() + s(_r.Accept(dw)) // IsInterfaceArray
		}
		dw.o()
		_s += dw.p() + "]\n"
	}
	if ec.StarArgs != nil {
		_s += dw.p() + "StarArgs: " + s(ec.StarArgs.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "StarArgs: nil\n"
	}
	if ec.KeywordArgs != nil {
		_s += dw.p() + "KeywordArgs: " + s(ec.KeywordArgs.Accept(dw)) // IsInterface
	} else {
		_s += dw.p() + "KeywordArgs: nil\n"
	}
	dw.o()
	_s += dw.p() + ")\n"
	return _s, nil
}
