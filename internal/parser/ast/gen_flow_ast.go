package ast

import (
	"guppy/internal/parser/tokenizer"
)

type VisitorData interface {
	VisitDataArgument(da DataArgument) (any, error)
	VisitDataArgumentList(dal DataArgumentList) (any, error)
	VisitDataImportAs(dia DataImportAs) (any, error)
	VisitDataListIter(dli DataListIter) (any, error)
	VisitDataListFor(dlf DataListFor) (any, error)
	VisitDataListIf(dli DataListIf) (any, error)
	VisitDataParameter(dp DataParameter) (any, error)
	VisitDataParameterList(dpl DataParameterList) (any, error)
	VisitDataSubscript(ds DataSubscript) (any, error)
}

type Data interface {
	Accept(vd VisitorData) (any, error)
}

type DataArgument struct {
	Assign string
	Expr   Expression
}

func NewDataArgument(
	Assign string,
	Expr Expression,
) *DataArgument {
	return &DataArgument{
		Assign: Assign,
		Expr:   Expr,
	}
}

func (da DataArgument) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataArgument(da)
}

type DataArgumentList struct {
	Args       []Expression
	NamedArgs  []*DataArgument
	StarArg    Expression
	KeywordArg Expression
}

func NewDataArgumentList(
	Args []Expression,
	NamedArgs []*DataArgument,
	StarArg Expression,
	KeywordArg Expression,
) *DataArgumentList {
	return &DataArgumentList{
		Args:       Args,
		NamedArgs:  NamedArgs,
		StarArg:    StarArg,
		KeywordArg: KeywordArg,
	}
}

func (dal DataArgumentList) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataArgumentList(dal)
}

type DataImportAs struct {
	Name []string
	As   string
}

func NewDataImportAs(
	Name []string,
	As string,
) *DataImportAs {
	return &DataImportAs{
		Name: Name,
		As:   As,
	}
}

func (dia DataImportAs) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataImportAs(dia)
}

type DataListIter struct {
	For *DataListFor
	If  *DataListIf
}

func NewDataListIter(
	For *DataListFor,
	If *DataListIf,
) *DataListIter {
	return &DataListIter{
		For: For,
		If:  If,
	}
}

func (dli DataListIter) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataListIter(dli)
}

type DataListFor struct {
	Idents []string
	Expr   Expression
	Iter   *DataListIter
}

func NewDataListFor(
	Idents []string,
	Expr Expression,
	Iter *DataListIter,
) *DataListFor {
	return &DataListFor{
		Idents: Idents,
		Expr:   Expr,
		Iter:   Iter,
	}
}

func (dlf DataListFor) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataListFor(dlf)
}

type DataListIf struct {
	Expr Expression
	Iter *DataListIter
}

func NewDataListIf(
	Expr Expression,
	Iter *DataListIter,
) *DataListIf {
	return &DataListIf{
		Expr: Expr,
		Iter: Iter,
	}
}

func (dli DataListIf) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataListIf(dli)
}

type DataParameter struct {
	Name       string
	Type       string
	Default    Expression
	StarArg    bool
	KeywordArg bool
}

func NewDataParameter(
	Name string,
	Type string,
	Default Expression,
	StarArg bool,
	KeywordArg bool,
) *DataParameter {
	return &DataParameter{
		Name:       Name,
		Type:       Type,
		Default:    Default,
		StarArg:    StarArg,
		KeywordArg: KeywordArg,
	}
}

func (dp DataParameter) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataParameter(dp)
}

type DataParameterList struct {
	List []*DataParameter
}

func NewDataParameterList(
	List []*DataParameter,
) *DataParameterList {
	return &DataParameterList{
		List: List,
	}
}

func (dpl DataParameterList) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataParameterList(dpl)
}

type DataSubscript struct {
	Start Expression
	End   Expression
	Range bool
}

func NewDataSubscript(
	Start Expression,
	End Expression,
	Range bool,
) *DataSubscript {
	return &DataSubscript{
		Start: Start,
		End:   End,
		Range: Range,
	}
}

func (ds DataSubscript) Accept(vd VisitorData) (any, error) {
	return vd.VisitDataSubscript(ds)
}

type VisitorStatement interface {
	VisitStatementProgram(sp StatementProgram) (any, error)
	VisitStatementExpression(se StatementExpression) (any, error)
	VisitStatementReturn(sr StatementReturn) (any, error)
	VisitStatementImportFrom(sif StatementImportFrom) (any, error)
	VisitStatementImportFromStar(sifs StatementImportFromStar) (any, error)
	VisitStatementImportNames(sin StatementImportNames) (any, error)
	VisitStatementAssert(sa StatementAssert) (any, error)
	VisitStatementIf(si StatementIf) (any, error)
	VisitStatementFunction(sf StatementFunction) (any, error)
	VisitStatementDecorated(sd StatementDecorated) (any, error)
	VisitStatementList(sl StatementList) (any, error)
}

type Statement interface {
	Accept(vs VisitorStatement) (any, error)
}

type StatementProgram struct {
	Statements *StatementList
}

func NewStatementProgram(
	Statements *StatementList,
) *StatementProgram {
	return &StatementProgram{
		Statements: Statements,
	}
}

func (sp StatementProgram) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementProgram(sp)
}

type StatementExpression struct {
	Assign []string
	Expr   Expression
}

func NewStatementExpression(
	Assign []string,
	Expr Expression,
) Statement {
	return StatementExpression{
		Assign: Assign,
		Expr:   Expr,
	}
}

func (se StatementExpression) Accept(vs VisitorStatement) (any, error) {
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

func (sr StatementReturn) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementReturn(sr)
}

type StatementImportFrom struct {
	From    []string
	Imports []*DataImportAs
}

func NewStatementImportFrom(
	From []string,
	Imports []*DataImportAs,
) Statement {
	return StatementImportFrom{
		From:    From,
		Imports: Imports,
	}
}

func (sif StatementImportFrom) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementImportFrom(sif)
}

type StatementImportFromStar struct {
	From []string
}

func NewStatementImportFromStar(
	From []string,
) Statement {
	return StatementImportFromStar{
		From: From,
	}
}

func (sifs StatementImportFromStar) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementImportFromStar(sifs)
}

type StatementImportNames struct {
	Imports []*DataImportAs
}

func NewStatementImportNames(
	Imports []*DataImportAs,
) Statement {
	return StatementImportNames{
		Imports: Imports,
	}
}

func (sin StatementImportNames) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementImportNames(sin)
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

func (sa StatementAssert) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementAssert(sa)
}

type StatementIf struct {
	Condition     []Expression
	Statement     []Statement
	StatementElse Statement
}

func NewStatementIf(
	Condition []Expression,
	Statement []Statement,
	StatementElse Statement,
) Statement {
	return StatementIf{
		Condition:     Condition,
		Statement:     Statement,
		StatementElse: StatementElse,
	}
}

func (si StatementIf) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementIf(si)
}

type StatementFunction struct {
	Token  string
	Params *DataParameterList
	Body   Statement
}

func NewStatementFunction(
	Token string,
	Params *DataParameterList,
	Body Statement,
) Statement {
	return StatementFunction{
		Token:  Token,
		Params: Params,
		Body:   Body,
	}
}

func (sf StatementFunction) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementFunction(sf)
}

type StatementDecorated struct {
}

func NewStatementDecorated() Statement {
	return StatementDecorated{}
}

func (sd StatementDecorated) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementDecorated(sd)
}

type StatementList struct {
	Statements []Statement
}

func NewStatementList(
	Statements []Statement,
) *StatementList {
	return &StatementList{
		Statements: Statements,
	}
}

func (sl StatementList) Accept(vs VisitorStatement) (any, error) {
	return vs.VisitStatementList(sl)
}

type VisitorExpression interface {
	VisitExpressionList(el ExpressionList) (any, error)
	VisitExpressionListMaker(elm ExpressionListMaker) (any, error)
	VisitExpressionBinary(eb ExpressionBinary) (any, error)
	VisitExpressionDict(ed ExpressionDict) (any, error)
	VisitExpressionGrouping(eg ExpressionGrouping) (any, error)
	VisitExpressionLambda(el ExpressionLambda) (any, error)
	VisitExpressionLiteral(el ExpressionLiteral) (any, error)
	VisitExpressionLogical(el ExpressionLogical) (any, error)
	VisitExpressionTernary(et ExpressionTernary) (any, error)
	VisitExpressionTuple(et ExpressionTuple) (any, error)
	VisitExpressionUnary(eu ExpressionUnary) (any, error)
	VisitExpressionVariable(ev ExpressionVariable) (any, error)
	VisitExpressionMember(em ExpressionMember) (any, error)
	VisitExpressionSubscript(es ExpressionSubscript) (any, error)
	VisitExpressionCall(ec ExpressionCall) (any, error)
}

type Expression interface {
	Accept(ve VisitorExpression) (any, error)
}

type ExpressionList struct {
	Expressions []Expression
}

func NewExpressionList(
	Expressions []Expression,
) Expression {
	return ExpressionList{
		Expressions: Expressions,
	}
}

func (el ExpressionList) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionList(el)
}

type ExpressionListMaker struct {
	Expr Expression
	For  *DataListFor
}

func NewExpressionListMaker(
	Expr Expression,
	For *DataListFor,
) Expression {
	return ExpressionListMaker{
		Expr: Expr,
		For:  For,
	}
}

func (elm ExpressionListMaker) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionListMaker(elm)
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

func (eb ExpressionBinary) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionBinary(eb)
}

type ExpressionDict struct {
	Keys   []Expression
	Values []Expression
}

func NewExpressionDict(
	Keys []Expression,
	Values []Expression,
) Expression {
	return ExpressionDict{
		Keys:   Keys,
		Values: Values,
	}
}

func (ed ExpressionDict) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionDict(ed)
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

func (eg ExpressionGrouping) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionGrouping(eg)
}

type ExpressionLambda struct {
	Identifier string
	Expr       Expression
}

func NewExpressionLambda(
	Identifier string,
	Expr Expression,
) Expression {
	return ExpressionLambda{
		Identifier: Identifier,
		Expr:       Expr,
	}
}

func (el ExpressionLambda) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionLambda(el)
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

func (el ExpressionLiteral) Accept(ve VisitorExpression) (any, error) {
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

func (el ExpressionLogical) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionLogical(el)
}

type ExpressionTernary struct {
	Left  Expression
	Cond  Expression
	Right Expression
}

func NewExpressionTernary(
	Left Expression,
	Cond Expression,
	Right Expression,
) Expression {
	return ExpressionTernary{
		Left:  Left,
		Cond:  Cond,
		Right: Right,
	}
}

func (et ExpressionTernary) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionTernary(et)
}

type ExpressionTuple struct {
	Expressions []Expression
}

func NewExpressionTuple(
	Expressions []Expression,
) Expression {
	return ExpressionTuple{
		Expressions: Expressions,
	}
}

func (et ExpressionTuple) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionTuple(et)
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

func (eu ExpressionUnary) Accept(ve VisitorExpression) (any, error) {
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

func (ev ExpressionVariable) Accept(ve VisitorExpression) (any, error) {
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

func (em ExpressionMember) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionMember(em)
}

type ExpressionSubscript struct {
	Expr  Expression
	Start Expression
	End   Expression
	Range bool
}

func NewExpressionSubscript(
	Expr Expression,
	Start Expression,
	End Expression,
	Range bool,
) Expression {
	return ExpressionSubscript{
		Expr:  Expr,
		Start: Start,
		End:   End,
		Range: Range,
	}
}

func (es ExpressionSubscript) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionSubscript(es)
}

type ExpressionCall struct {
	Expr       Expression
	Args       []Expression
	NamedArgs  []*DataArgument
	StarArg    Expression
	KeywordArg Expression
}

func NewExpressionCall(
	Expr Expression,
	Args []Expression,
	NamedArgs []*DataArgument,
	StarArg Expression,
	KeywordArg Expression,
) Expression {
	return ExpressionCall{
		Expr:       Expr,
		Args:       Args,
		NamedArgs:  NamedArgs,
		StarArg:    StarArg,
		KeywordArg: KeywordArg,
	}
}

func (ec ExpressionCall) Accept(ve VisitorExpression) (any, error) {
	return ve.VisitExpressionCall(ec)
}
