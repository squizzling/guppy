package interpreter

import (
	"errors"
	"fmt"

	"github.com/squizzling/types/pkg/result"

	"guppy/internal/parser/ast"
	"guppy/internal/parser/tokenizer"
)

func (i *Interpreter) VisitExpressionBinary(eb ast.ExpressionBinary) any {
	defer i.trace()()

	resultLeft := r(eb.Left.Accept(i))
	if !resultLeft.Ok() {
		return result.Err[Object](resultLeft.Err())
	}
	left := resultLeft.Value()

	resultRight := r(eb.Right.Accept(i))
	if !resultRight.Ok() {
		return result.Err[Object](resultRight.Err())
	}
	right := resultRight.Value()

	switch eb.Op.Type {
	case tokenizer.TokenTypeAnd:
		return i.doAnd(left, right)
	default:
		return result.Err[Object](fmt.Errorf("unhandled binary op: %s", eb.Op.Type))
	}
}

func (i *Interpreter) VisitExpressionCall(ec ast.ExpressionCall) any {
	defer i.trace()()

	exprResult := r(ec.Expr.Accept(i))
	if !exprResult.Ok() {
		return exprResult
	}
	expr := exprResult.Value()

	i.pushScope()
	defer i.popScope()

	var argData []ArgData
	if argDataResult := i.doArgs(expr); !argDataResult.Ok() {
		return result.Err[Object](argDataResult.Err())
	} else {
		argData = argDataResult.Value()
	}

	ecArgs := ec.Args

	if lv, ok := expr.(*ObjectLValue); ok {
		// We have an lvalue, push lv.left on to the start of ecArgs
		ecArgs = append([]ast.DataArgument{{Expr: ast.NewExpressionLiteral(lv.left)}}, ecArgs...)
	}

	if len(ecArgs) > len(argData) {
		return result.Err[Object](fmt.Errorf("too many args provided (provided=%d, expected=%d)", len(ecArgs), len(argData)))
	}

	providedArgs := make(map[string]bool)
	for idx, exprArg := range ecArgs {
		if exprArg.Assign == "" {
			if argResult := r(exprArg.Expr.Accept(i)); !argResult.Ok() {
				return result.Err[Object](argResult.Err())
			} else {
				i.Scope.DeclareSet(argData[idx].Name, argResult.Value())
				providedArgs[argData[idx].Name] = true
			}
		} else if _, ok := providedArgs[exprArg.Assign]; ok {
			// Everything after this will be provided, so we don't need to track the
			// index anymore.  Instead we just need to make sure there's no duplicates
			return result.Err[Object](fmt.Errorf("duplicate argument %s", exprArg.Assign))
		} else {
			// Set it
			// TODO: This is duplicated of the first arm
			if argResult := r(exprArg.Expr.Accept(i)); !argResult.Ok() {
				return result.Err[Object](argResult.Err())
			} else {
				i.Scope.DeclareSet(argData[idx].Name, argResult.Value())
				providedArgs[argData[idx].Name] = true
			}
		}
	}

	// Loop through argData.  If an argument is not provided, and it's got a default, then provide it

	for _, arg := range argData {
		if _, ok := providedArgs[arg.Name]; !ok {
			if arg.Default != nil {
				i.Scope.DeclareSet(arg.Name, arg.Default)
				providedArgs[arg.Name] = true
			}
		}
	}

	// If the length of providedArgs matches the expected, we can invoke the call

	if len(argData) != len(providedArgs) {
		// TODO: Make this better
		return result.Err[Object](fmt.Errorf("arg count wrong %d vs %d", len(argData), len(providedArgs)))
	}

	if ec.StarArgs != nil {
		return result.Err[Object](errors.New("star arguments are not supported"))
	}
	if ec.KeywordArgs != nil {
		return result.Err[Object](errors.New("keyword arguments are not supported"))
	}

	return i.doCall(expr)
}

func (i *Interpreter) VisitExpressionGrouping(eg ast.ExpressionGrouping) any {
	defer i.trace()()

	panic("ExpressionGrouping")
}

func (i *Interpreter) VisitExpressionList(el ast.ExpressionList) any {
	defer i.trace()()

	var o []Object
	for _, expr := range el.Expressions {
		exprResult := r(expr.Accept(i))
		if !exprResult.Ok() {
			return result.Err[Object](exprResult.Err())
		}
		o = append(o, exprResult.Value())
	}
	return result.Ok(NewObjectList(o...))
}

func (i *Interpreter) VisitExpressionLiteral(el ast.ExpressionLiteral) any {
	defer i.trace("Value: %v", el.Value)()

	switch v := el.Value.(type) {
	case string:
		return result.Ok(NewObjectString(v))
	case nil:
		return result.Ok(NewObjectNone())
	case Object:
		return result.Ok(v)
	default:
		return result.Err[Object](fmt.Errorf("unknown literal type: %T", v))
	}
}

func (i *Interpreter) VisitExpressionLogical(el ast.ExpressionLogical) any {
	defer i.trace()()

	panic("ExpressionLogical")
}

func (i *Interpreter) VisitExpressionMember(em ast.ExpressionMember) any {
	defer i.trace("Member: %s", em.Identifier)()

	if exprResult := r(em.Expr.Accept(i)); !exprResult.Ok() {
		return exprResult
	} else {
		return exprResult.Value().Member(i, exprResult.Value(), em.Identifier)
	}
}

func (i *Interpreter) VisitExpressionSubscript(es ast.ExpressionSubscript) any {
	defer i.trace()()

	panic("ExpressionSubscript")
}

func (i *Interpreter) VisitExpressionUnary(eu ast.ExpressionUnary) any {
	defer i.trace()()

	panic("ExpressionUnary")
}

func (i *Interpreter) VisitExpressionVariable(ev ast.ExpressionVariable) any {
	defer i.trace("Identifier: %s", ev.Identifier)()

	return i.Scope.Get(ev.Identifier)
}
