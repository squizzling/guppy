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
	// TODO: I think there's some weirdness happening with calls and self, but
	//       because flow doesn't support classes, it might not be a problem?
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

	/*
	  If keyword arguments are present, they are first converted to positional arguments, as follows.
	  First, a list of unfilled slots is created for the formal parameters.
	  If there are N positional arguments,
	    they are placed in the first N slots.
	  Next, for each keyword argument, the identifier is used to determine the corresponding slot (if the identifier is the same as the first formal parameter name, the first slot is used, and so on).
	  If the slot is already filled,
	    a TypeError exception is raised.
	  Otherwise,
	    the value of the argument is placed in the slot, filling it (even if the expression is None, it fills the slot).
	  When all arguments have been processed, the slots that are still unfilled are filled with the corresponding default value from the function definition.
	  (Default values are calculated, once, when the function is defined; thus, a mutable object such as a list or dictionary used as default value will be shared by all calls that don’t specify an argument value for the corresponding slot; this should usually be avoided.)
	  If there are any unfilled slots for which no default value is specified,
	    a TypeError exception is raised.
	  Otherwise,
	   the list of filled slots is used as the argument list for the call.
	*/
	/*
	  If there are more positional arguments than there are formal parameter slots,
	    a TypeError exception is raised,
	  unless a formal parameter using the syntax *identifier is present; in this case,
	    that formal parameter receives a tuple containing the excess positional arguments (or an empty tuple if there were no excess positional arguments).
	*/
	/*
	  If any keyword argument does not correspond to a formal parameter name,
	    a TypeError exception is raised,
	  unless a formal parameter using the syntax **identifier is present; in this case,
	    that formal parameter receives a dictionary containing the excess keyword arguments (using the keywords as keys and the argument values as corresponding values), or a (new) empty dictionary if there were no excess keyword arguments.
	*/

	if lv, ok := expr.(*ObjectLValue); ok {
		// We have an lvalue, push lv.left on to the start of ecArgs
		ecArgs = append([]ast.DataArgument{{Expr: ast.NewExpressionLiteral(lv.left)}}, ecArgs...)
	}

	var sProvidedArgs []string
	var expectedArgs []string

	for _, a := range ecArgs {
		sProvidedArgs = append(sProvidedArgs, a.Assign)
	}
	for _, a := range argData {
		expectedArgs = append(expectedArgs, a.Name)
	}

	if len(ecArgs) > len(argData) {
		// calculate ecArgs - argData
		var extraArgs []string
		for _, a1 := range ecArgs {
			if a1.Assign == "" {
				continue
			}
			found := false
			for _, a2 := range argData {
				if a1.Assign == a2.Name {
					found = true
					break
				}
			}
			if !found {
				extraArgs = append(extraArgs, a1.Assign)
			}
		}
		return nil, fmt.Errorf("too many args provided (provided=%s, expected=%s, extra=%s)", sProvidedArgs, expectedArgs, extraArgs)
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

func (i *Interpreter) VisitExpressionTernary(et ast.ExpressionTernary) (any, error) {
	defer i.trace()()

	panic("ExpressionTernary")
}

func (i *Interpreter) VisitExpressionUnary(eu ast.ExpressionUnary) (any, error) {
	defer i.trace()()

	panic("ExpressionUnary")
}

func (i *Interpreter) VisitExpressionVariable(ev ast.ExpressionVariable) (any, error) {
	defer i.trace("Identifier: %s", ev.Identifier)()

	return i.Scope.Get(ev.Identifier)
}
