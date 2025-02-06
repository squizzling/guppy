package ast

import (
	"guppy/internal/parser/tokenizer"
)

type VisitorData interface {
	VisitDataArgument(da DataArgument) any
}

type Data interface {
	Accept(vd VisitorData) any
}

type DataArgument struct {
	Assign string
	Expr   Expression
}

func NewDataArgument(
	Assign string,
	Expr Expression,
) DataArgument {
	return DataArgument{
		Assign: Assign,
		Expr:   Expr,
	}
}

func (da DataArgument) Accept(vd VisitorData) any {
	return vd.VisitDataArgument(da)
}

type VisitorStatement interface {
	VisitStatementProgram(sp StatementProgram) any
	VisitStatementExpression(se StatementExpression) any
	VisitStatementReturn(sr StatementReturn) any
	VisitStatementImport(si StatementImport) any
	VisitStatementAssert(sa StatementAssert) any
	VisitStatementIf(si StatementIf) any
	VisitStatementFunction(sf StatementFunction) any
	VisitStatementDecorated(sd StatementDecorated) any
	VisitStatementList(sl StatementList) any
}

type Statement interface {
	Accept(vs VisitorStatement) any
}

type StatementProgram struct {
	Statements StatementList
}

func NewStatementProgram(
	Statements StatementList,
) StatementProgram {
	return StatementProgram{
		Statements: Statements,
	}
}

func (sp StatementProgram) Accept(vs VisitorStatement) any {
	return vs.VisitStatementProgram(sp)
}

type StatementExpression struct {
	Assign []string
	Expr   ExpressionList
}

func NewStatementExpression(
	Assign []string,
	Expr ExpressionList,
) Statement {
	return StatementExpression{
		Assign: Assign,
		Expr:   Expr,
	}
}

func (se StatementExpression) Accept(vs VisitorStatement) any {
	return vs.VisitStatementExpression(se)
}

type StatementReturn struct {
	Expr Expression
}

func NewStatementReturn(
	Expr Expression,
) Statement {
	return StatementReturn{
		Expr: Expr,
	}
}

func (sr StatementReturn) Accept(vs VisitorStatement) any {
	return vs.VisitStatementReturn(sr)
}

type StatementImport struct {
}

func NewStatementImport() Statement {
	return StatementImport{}
}

func (si StatementImport) Accept(vs VisitorStatement) any {
	return vs.VisitStatementImport(si)
}

type StatementAssert struct {
	Test  Expression
	Raise Expression
}

func NewStatementAssert(
	Test Expression,
	Raise Expression,
) Statement {
	return StatementAssert{
		Test:  Test,
		Raise: Raise,
	}
}

func (sa StatementAssert) Accept(vs VisitorStatement) any {
	return vs.VisitStatementAssert(sa)
}

type StatementIf struct {
	Condition     []Expression
	Statement     []StatementList
	StatementElse StatementList
}

func NewStatementIf(
	Condition []Expression,
	Statement []StatementList,
	StatementElse StatementList,
) Statement {
	return StatementIf{
		Condition:     Condition,
		Statement:     Statement,
		StatementElse: StatementElse,
	}
}

func (si StatementIf) Accept(vs VisitorStatement) any {
	return vs.VisitStatementIf(si)
}

type StatementFunction struct {
	Token string
	Body  StatementList
}

func NewStatementFunction(
	Token string,
	Body StatementList,
) Statement {
	return StatementFunction{
		Token: Token,
		Body:  Body,
	}
}

func (sf StatementFunction) Accept(vs VisitorStatement) any {
	return vs.VisitStatementFunction(sf)
}

type StatementDecorated struct {
}

func NewStatementDecorated() Statement {
	return StatementDecorated{}
}

func (sd StatementDecorated) Accept(vs VisitorStatement) any {
	return vs.VisitStatementDecorated(sd)
}

type StatementList struct {
	Statements []Statement
}

func NewStatementList(
	Statements []Statement,
) StatementList {
	return StatementList{
		Statements: Statements,
	}
}

func (sl StatementList) Accept(vs VisitorStatement) any {
	return vs.VisitStatementList(sl)
}

type VisitorExpression interface {
	VisitExpressionList(el ExpressionList) any
	VisitExpressionBinary(eb ExpressionBinary) any
	VisitExpressionGrouping(eg ExpressionGrouping) any
	VisitExpressionLiteral(el ExpressionLiteral) any
	VisitExpressionLogical(el ExpressionLogical) any
	VisitExpressionUnary(eu ExpressionUnary) any
	VisitExpressionVariable(ev ExpressionVariable) any
	VisitExpressionMember(em ExpressionMember) any
	VisitExpressionSubscript(es ExpressionSubscript) any
	VisitExpressionCall(ec ExpressionCall) any
}

type Expression interface {
	Accept(ve VisitorExpression) any
}

type ExpressionList struct {
	Expressions []Expression
	Tuple       bool
}

func NewExpressionList(
	Expressions []Expression,
	Tuple bool,
) ExpressionList {
	return ExpressionList{
		Expressions: Expressions,
		Tuple:       Tuple,
	}
}

func (el ExpressionList) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionList(el)
}

type ExpressionBinary struct {
	Left  Expression
	Op    tokenizer.Token
	Right Expression
}

func NewExpressionBinary(
	Left Expression,
	Op tokenizer.Token,
	Right Expression,
) Expression {
	return ExpressionBinary{
		Left:  Left,
		Op:    Op,
		Right: Right,
	}
}

func (eb ExpressionBinary) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionBinary(eb)
}

type ExpressionGrouping struct {
	Expr Expression
}

func NewExpressionGrouping(
	Expr Expression,
) Expression {
	return ExpressionGrouping{
		Expr: Expr,
	}
}

func (eg ExpressionGrouping) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionGrouping(eg)
}

type ExpressionLiteral struct {
	Value any
}

func NewExpressionLiteral(
	Value any,
) Expression {
	return ExpressionLiteral{
		Value: Value,
	}
}

func (el ExpressionLiteral) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionLiteral(el)
}

type ExpressionLogical struct {
	Left  Expression
	Op    tokenizer.Token
	Right Expression
}

func NewExpressionLogical(
	Left Expression,
	Op tokenizer.Token,
	Right Expression,
) Expression {
	return ExpressionLogical{
		Left:  Left,
		Op:    Op,
		Right: Right,
	}
}

func (el ExpressionLogical) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionLogical(el)
}

type ExpressionUnary struct {
	Op    tokenizer.TokenType
	Right Expression
}

func NewExpressionUnary(
	Op tokenizer.TokenType,
	Right Expression,
) Expression {
	return ExpressionUnary{
		Op:    Op,
		Right: Right,
	}
}

func (eu ExpressionUnary) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionUnary(eu)
}

type ExpressionVariable struct {
	Identifier string
}

func NewExpressionVariable(
	Identifier string,
) Expression {
	return ExpressionVariable{
		Identifier: Identifier,
	}
}

func (ev ExpressionVariable) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionVariable(ev)
}

type ExpressionMember struct {
	Expr       Expression
	Identifier string
}

func NewExpressionMember(
	Expr Expression,
	Identifier string,
) Expression {
	return ExpressionMember{
		Expr:       Expr,
		Identifier: Identifier,
	}
}

func (em ExpressionMember) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionMember(em)
}

type ExpressionSubscript struct {
	Expr       Expression
	Identifier string
}

func NewExpressionSubscript(
	Expr Expression,
	Identifier string,
) Expression {
	return ExpressionSubscript{
		Expr:       Expr,
		Identifier: Identifier,
	}
}

func (es ExpressionSubscript) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionSubscript(es)
}

type ExpressionCall struct {
	Expr        Expression
	Args        []DataArgument
	StarArgs    Expression
	KeywordArgs Expression
}

func NewExpressionCall(
	Expr Expression,
	Args []DataArgument,
	StarArgs Expression,
	KeywordArgs Expression,
) Expression {
	return ExpressionCall{
		Expr:        Expr,
		Args:        Args,
		StarArgs:    StarArgs,
		KeywordArgs: KeywordArgs,
	}
}

func (ec ExpressionCall) Accept(ve VisitorExpression) any {
	return ve.VisitExpressionCall(ec)
}
