package interpreter

import (
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

func findParamSlot(params *Params, name string) int {
	for idx, param := range params.Params {
		if param.Name == name {
			return idx
		}
	}
	for idx, param := range params.KWParams {
		if param.Name == name {
			return idx + len(params.Params)
		}
	}
	return -1
}

func (i *Interpreter) VisitExpressionCall(ec ast.ExpressionCall) (any, error) {
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

	// We don't perfectly implement this, because this is 3 distinct sections, but the last 2 are done inline with the first.
	// We're trying to do it exactly like that, however flow is somewhere between py2 and py3, so it's a bit of a mess.
	// Once this is stable, as verified by hammering existing flow in the wild, we'll add some tests to verify the interface,
	// and clean it all up.

	// TODO: I think there's some weirdness happening with calls and self, but
	//       because flow doesn't support classes, it might not be a problem?
	defer i.trace()()

	expr, err := r(ec.Expr.Accept(i))
	if err != nil {
		return nil, err
	}

	paramData, err := i.doParams(expr)
	if err != nil {
		return nil, err
	}

	unnamedArgs := ec.Args
	if lv, ok := expr.(*ObjectLValue); ok {
		// TODO: Can we push this up to *ObjectLValue.Call?
		unnamedArgs = append([]ast.Expression{ast.NewExpressionLiteral(lv.left)}, unnamedArgs...)
	}

	/*
	  If there are more positional arguments than there are formal parameter slots,
	    a TypeError exception is raised,
	  unless a formal parameter using the syntax *identifier is present; in this case,
	    that formal parameter receives a tuple containing the excess positional arguments (or an empty tuple if there were no excess positional arguments).
	*/
	if len(unnamedArgs) > len(paramData.Params) {
		if paramData.StarParam == "" { // No star param
			return nil, fmt.Errorf("too many params (1)")
		} else if ec.StarArg != nil { // Star param is already occupied
			return nil, fmt.Errorf("too many params (2)")
		}
	}

	// If keyword arguments are present, they are first converted to positional arguments, as follows.
	//        First, a list of unfilled slots is created for the formal parameters.
	formalParams := make([]Object, len(paramData.Params)+len(paramData.KWParams))
	formalNames := make([]string, len(paramData.Params)+len(paramData.KWParams))
	occupiedParam := make([]bool, len(paramData.Params)+len(paramData.KWParams))

	//        If there are N positional arguments,
	//          they are placed in the first N slots.
	for idx, arg := range unnamedArgs[:min(len(unnamedArgs), len(paramData.Params))] {
		// Note: we only want to fill up to len(paramData.Params), not up to the full slice.  The second half of
		// the slice is for KWParams.
		if o, err := r(arg.Accept(i)); err != nil {
			return nil, err
		} else {
			formalParams[idx] = o
			formalNames[idx] = paramData.Params[idx].Name
			occupiedParam[idx] = true
		}
	}

	var kwArguments map[string]Object
	if paramData.KWParam != "" {
		kwArguments = make(map[string]Object)
	}

	// Next, for each keyword argument,
	for _, arg := range ec.NamedArgs {
		// the identifier is used to determine the corresponding slot (if the identifier is the same as the first formal parameter name, the first slot is used, and so on).
		slotIndex := findParamSlot(paramData, arg.Assign)

		//        If the slot is already filled,
		//          a TypeError exception is raised.
		if slotIndex == -1 {
			if paramData.KWParam == "" {
				return nil, fmt.Errorf("got an unexpected keyword argument: '%s'", arg.Assign)
			} else if o, err := r(arg.Expr.Accept(i)); err != nil {
				return nil, err
			} else {
				// TODO: Check if it's already present
				kwArguments[arg.Assign] = o
				continue
			}
		}
		if occupiedParam[slotIndex] {
			return nil, fmt.Errorf("got multiple values for keyword argument '%s'", arg.Assign)
		}

		//        Otherwise,
		//          the value of the argument is placed in the slot, filling it (even if the expression is None, it fills the slot).
		if o, err := r(arg.Expr.Accept(i)); err != nil {
			return nil, err
		} else {
			//fmt.Printf("Filling slot %d\n", slotIndex)
			formalParams[slotIndex] = o
			formalNames[slotIndex] = arg.Assign
			occupiedParam[slotIndex] = true
		}
	}

	for idx, param := range paramData.Params {
		if !occupiedParam[idx] && param.Default != nil {
			formalParams[idx] = param.Default
			formalNames[idx] = param.Name
			occupiedParam[idx] = true
		}
	}
	for idx, param := range paramData.KWParams {
		if !occupiedParam[idx+len(paramData.Params)] && param.Default != nil {
			formalParams[idx+len(paramData.Params)] = param.Default
			formalNames[idx+len(paramData.Params)] = param.Name
			occupiedParam[idx+len(paramData.Params)] = true
		}
	}

	// 	  If there are any unfilled slots for which no default value is specified,
	//	    a TypeError exception is raised.
	for idx, occupied := range occupiedParam {
		if !occupied {
			return nil, fmt.Errorf("parameter %d is not occupied", idx)
		}
	}

	//	  Otherwise,
	//	   the list of filled slots is used as the argument list for the call.

	starArgs, err := i.resolveStarArgs(unnamedArgs, len(formalParams), paramData)
	if err != nil {
		return nil, err
	}

	// Perform all argument resolution above here, so we don't pollute the scope as we evaluate things.
	i.pushScope()
	defer i.popScope()

	if err := i.assignArgs(
		formalNames, formalParams,
		paramData.StarParam, starArgs,
		paramData.KWParam, kwArguments,
	); err != nil {
		return nil, err
	}

	return i.doCall(expr)
}

func (i *Interpreter) assignArgs(
	formalNames []string,
	formalParams []Object,

	starParamName string,
	starArgs []Object,

	kwParamName string,
	kwArgs map[string]Object, // TODO: I don't like this map key, so punting it.
) error {
	if starParamName != "" {
		i.Scope.DeclareSet(starParamName, NewObjectTuple(starArgs...))
	}

	for idx, _ := range formalParams {
		i.Scope.DeclareSet(formalNames[idx], formalParams[idx])
	}
	if kwParamName != "" {
		return fmt.Errorf("can't set kwparam")
	}
	return nil
}

func (i *Interpreter) resolveStarArgs(unnamedArgs []ast.Expression, formalParamCount int, paramData *Params) ([]Object, error) {
	// TODO: I don't love this signature
	var starArgs []Object
	if len(unnamedArgs)-formalParamCount > 0 {
		if paramData.StarParam == "" {
			return nil, fmt.Errorf("too many params (3)")
		}
		for _, arg := range unnamedArgs[formalParamCount:] {
			if o, err := r(arg.Accept(i)); err != nil {
				return nil, err
			} else {
				starArgs = append(starArgs, o)
			}
		}
	}
	return starArgs, nil
}

func (i *Interpreter) VisitExpressionDict(ed ast.ExpressionDict) (any, error) {
	defer i.trace()()

	panic("ExpressionDict")
}

func (i *Interpreter) VisitExpressionGrouping(eg ast.ExpressionGrouping) (any, error) {
	defer i.trace()()

	panic("ExpressionGrouping")
}

func (i *Interpreter) VisitExpressionLambda(el ast.ExpressionLambda) (any, error) {
	defer i.trace()()

	panic("ExpressionLambda")
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

func (i *Interpreter) VisitExpressionListMaker(elm ast.ExpressionListMaker) (any, error) {
	defer i.trace()()

	panic("ExpressionListMaker")
}

func (i *Interpreter) VisitExpressionLiteral(el ast.ExpressionLiteral) (any, error) {
	defer i.trace("Value: %v", el.Value)()

	switch v := el.Value.(type) {
	case int:
		return NewObjectInt(v), nil
	case string:
		return NewObjectString(v), nil
	case nil:
		return NewObjectNone(), nil
	case Object:
		return v, nil
	case bool:
		return NewObjectBool(v), nil
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
