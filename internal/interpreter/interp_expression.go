package interpreter

import (
	"errors"
	"fmt"

	"guppy/internal/parser/ast"
	"guppy/internal/parser/tokenizer"
)

func (i *Interpreter) VisitExpressionBinary(eb ast.ExpressionBinary) (any, error) {
	defer i.trace()()

	left, err := r(eb.Left.Accept(i))
	if err != nil {
		return nil, err
	}

	right, err := r(eb.Right.Accept(i))
	if err != nil {
		return nil, err
	}

	switch eb.Op.Type {
	case tokenizer.TokenTypeAnd:
		return i.doAnd(left, right)
	default:
		return nil, fmt.Errorf("unhandled binary op: %s", eb.Op.Type)
	}
}

func (i *Interpreter) VisitExpressionCall(ec ast.ExpressionCall) (any, error) {
	defer i.trace()()

	expr, err := r(ec.Expr.Accept(i))
	if err != nil {
		return nil, err
	}

	i.pushScope()
	defer i.popScope()

	argData, err := i.doArgs(expr)
	if err != nil {
		return nil, err
	}

	ecArgs := ec.Args

	if lv, ok := expr.(*ObjectLValue); ok {
		// We have an lvalue, push lv.left on to the start of ecArgs
		ecArgs = append([]ast.DataArgument{{Expr: ast.NewExpressionLiteral(lv.left)}}, ecArgs...)
	}

	if len(ecArgs) > len(argData) {
		return nil, fmt.Errorf("too many args provided (provided=%d, expected=%d)", len(ecArgs), len(argData))
	}

	providedArgs := make(map[string]bool)
	for idx, exprArg := range ecArgs {
		if exprArg.Assign == "" {
			if arg, err := r(exprArg.Expr.Accept(i)); err != nil {
				return nil, err
			} else {
				i.Scope.DeclareSet(argData[idx].Name, arg)
				providedArgs[argData[idx].Name] = true
			}
		} else if _, ok := providedArgs[exprArg.Assign]; ok {
			// Everything after this will be provided, so we don't need to track the
			// index anymore.  Instead we just need to make sure there's no duplicates
			return nil, fmt.Errorf("duplicate argument %s", exprArg.Assign)
		} else {
			// Set it
			// TODO: This is duplicated of the first arm
			if arg, err := r(exprArg.Expr.Accept(i)); err != nil {
				return nil, err
			} else {
				i.Scope.DeclareSet(argData[idx].Name, arg)
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
		return nil, fmt.Errorf("arg count wrong %d vs %d", len(argData), len(providedArgs))
	}

	if ec.StarArgs != nil {
		return nil, errors.New("star arguments are not supported")
	}
	if ec.KeywordArgs != nil {
		return nil, errors.New("keyword arguments are not supported")
	}

	return i.doCall(expr)
}

func (i *Interpreter) VisitExpressionGrouping(eg ast.ExpressionGrouping) (any, error) {
	defer i.trace()()

	panic("ExpressionGrouping")
}

func (i *Interpreter) VisitExpressionList(el ast.ExpressionList) (any, error) {
	defer i.trace()()

	var o []Object
	for _, expr := range el.Expressions {
		exprResult, err := expr.Accept(i)
		if err != nil {
			return nil, err
		}
		o = append(o, exprResult.(Object))
	}
	return NewObjectList(o...), nil
}

func (i *Interpreter) VisitExpressionLiteral(el ast.ExpressionLiteral) (any, error) {
	defer i.trace("Value: %v", el.Value)()

	switch v := el.Value.(type) {
	case string:
		return NewObjectString(v), nil
	case nil:
		return NewObjectNone(), nil
	case Object:
		return v, nil
	default:
		return nil, fmt.Errorf("unknown literal type: %T", v)
	}
}

func (i *Interpreter) VisitExpressionLogical(el ast.ExpressionLogical) (any, error) {
	defer i.trace()()

	panic("ExpressionLogical")
}

func (i *Interpreter) VisitExpressionMember(em ast.ExpressionMember) (any, error) {
	defer i.trace("Member: %s", em.Identifier)()

	if expr, err := r(em.Expr.Accept(i)); err != nil {
		return nil, err
	} else {
		return expr.Member(i, expr, em.Identifier)
	}
}

func (i *Interpreter) VisitExpressionSubscript(es ast.ExpressionSubscript) (any, error) {
	defer i.trace()()

	panic("ExpressionSubscript")
}

func (i *Interpreter) VisitExpressionUnary(eu ast.ExpressionUnary) (any, error) {
	defer i.trace()()

	panic("ExpressionUnary")
}

func (i *Interpreter) VisitExpressionVariable(ev ast.ExpressionVariable) (any, error) {
	defer i.trace("Identifier: %s", ev.Identifier)()

	return i.Scope.Get(ev.Identifier)
}
