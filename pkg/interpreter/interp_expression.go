package interpreter

import (
	"fmt"

	"guppy/pkg/parser/ast"
	"guppy/pkg/parser/tokenizer"
)

func (i *Interpreter) VisitExpressionBinary(eb ast.ExpressionBinary) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

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

func (i *Interpreter) VisitExpressionCall(ec ast.ExpressionCall) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	i.debug("Entering %#v", ec.Expr)

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

	// Convert expressions and splat expression in to objects
	objFunc, unnamedArgs, err := i.resolveUnnamedArgs(ec.Expr, ec.Args, ec.StarArg)
	if err != nil {
		return nil, err
	}

	i.debug("%d unnamed args", len(unnamedArgs))
	for idx, a := range unnamedArgs {
		i.debug("%d: %#v", idx, a)
	}

	paramData, err := i.doParams(objFunc)
	if err != nil {
		return nil, err
	}

	paramData.Dump(i)

	/*
	  If there are more positional arguments than there are formal parameter slots,
	    a TypeError exception is raised,
	  unless a formal parameter using the syntax *identifier is present; in this case,
	    that formal parameter receives a tuple containing the excess positional arguments (or an empty tuple if there were no excess positional arguments).
	*/

	// If keyword arguments are present, they are first converted to positional arguments, as follows.
	//        First, a list of unfilled slots is created for the formal parameters.
	formalParams := make([]Object, len(paramData.Params)+len(paramData.KWParams))
	formalNames := make([]string, len(paramData.Params)+len(paramData.KWParams))
	occupiedParam := make([]bool, len(paramData.Params)+len(paramData.KWParams))
	i.debug("formalParams: %d", len(formalParams))

	//        If there are N positional arguments,
	//          they are placed in the first N slots.
	for idx, arg := range unnamedArgs[:min(len(unnamedArgs), len(paramData.Params))] {
		// Note: we only want to fill up to len(paramData.Params), not up to the full slice.  The second half of
		// the slice is for KWParams.
		formalParams[idx] = arg
		formalNames[idx] = paramData.Params[idx].Name
		occupiedParam[idx] = true
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

	var starArgs []Object
	if len(unnamedArgs)-len(formalParams) > 0 {
		if !(paramData.StarParam != "") {
			return nil, fmt.Errorf("passing extra arguments to function not expecting it")
		} else {
			starArgs = unnamedArgs[len(formalParams):]
		}
	}

	// If any of our arguments are deferred, then we're deferred
	desired := []string{}
	for _, p := range formalParams {
		if od, ok := p.(*ObjectDeferred); ok {
			desired = append(desired, od.desired...)
		}
	}
	for _, p := range starArgs {
		if od, ok := p.(*ObjectDeferred); ok {
			desired = append(desired, od.desired...)
		}
	}
	for _, p := range kwArguments {
		if od, ok := p.(*ObjectDeferred); ok {
			desired = append(desired, od.desired...)
		}
	}
	if len(desired) > 0 {
		return NewObjectDeferred(ec, desired...), nil
	}

	// Perform all argument resolution above here, so we don't pollute the scope as we evaluate things.
	if s, ok := objFunc.(*ObjectFunction); ok {
		i.pushClosure(s.scope)
	} else {
		i.pushClosure(i.Globals)
	}
	defer i.popScope()

	if err := i.assignArgs(
		formalNames, formalParams,
		paramData.StarParam, starArgs,
		paramData.KWParam, kwArguments,
	); err != nil {
		return nil, err
	}

	return i.doCall(objFunc)
}

func (i *Interpreter) resolveUnnamedArgs(exprFunction ast.Expression, unnamedArgExpressions []ast.Expression, starArg ast.Expression) (Object, []Object, error) {
	var unnamedArgs []Object

	// This effectively resolves "self", or the x in `x.y(...)` (which is y(x, ...))
	objFunction, err := r(exprFunction.Accept(i))
	if err != nil {
		return nil, nil, err
	} else if lv, ok := objFunction.(*ObjectLValue); ok {
		unnamedArgs = append(unnamedArgs, lv.left)
	}

	for _, expr := range unnamedArgExpressions {
		if o, err := r(expr.Accept(i)); err != nil {
			return nil, nil, err
		} else {
			unnamedArgs = append(unnamedArgs, o)
		}
	}

	if starArg != nil {
		if starArgs, err := r(starArg.Accept(i)); err != nil {
			return nil, nil, err
		} else {
			switch starArgs := starArgs.(type) {
			case *ObjectList:
				unnamedArgs = append(unnamedArgs, starArgs.Items...)
			case *ObjectTuple:
				unnamedArgs = append(unnamedArgs, starArgs.Items...)
			case *ObjectDeferred:
				unnamedArgs = append(unnamedArgs, starArgs)
			default:
				return nil, nil, fmt.Errorf("[resolveUnnamedArgs] expecting *interpreter.ObjectList or *interpreter.ObjectTuple got %T", starArgs)
			}
		}
	}
	return objFunction, unnamedArgs, nil
}

func (i *Interpreter) assignArgs(
	formalNames []string,
	formalParams []Object,

	starParamName string,
	starArgs []Object,

	kwParamName string,
	kwArgs map[string]Object,
) error {
	if starParamName != "" {
		if err := i.Scope.Set(starParamName, NewObjectTuple(starArgs...)); err != nil {
			return err
		}
	}

	for idx, _ := range formalParams {
		if err := i.Scope.Set(formalNames[idx], formalParams[idx]); err != nil {
			return err
		}
	}
	if kwParamName != "" {
		if err := i.Scope.Set(kwParamName, NewObjectDictFromMap(kwArgs)); err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) VisitExpressionDict(ed ast.ExpressionDict) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	var items []DictItem
	for idx, keyExpr := range ed.Keys {
		valExpr := ed.Values[idx]

		if key, err := r(keyExpr.Accept(i)); err != nil {
			return nil, err
		} else if val, err := r(valExpr.Accept(i)); err != nil {
			return nil, err
		} else {
			items = append(items, DictItem{Key: key, Value: val})
		}
	}

	return NewObjectDict(items), nil
}

func (i *Interpreter) VisitExpressionGrouping(eg ast.ExpressionGrouping) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("ExpressionGrouping")
}

func (i *Interpreter) VisitExpressionLambda(el ast.ExpressionLambda) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("ExpressionLambda")
}

func (i *Interpreter) VisitExpressionList(el ast.ExpressionList) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	var o []Object
	var desired []string
	for _, expr := range el.Expressions {
		exprResult, err := r(expr.Accept(i))
		if err != nil {
			return nil, err
		} else if od, ok := exprResult.(*ObjectDeferred); ok {
			desired = append(desired, od.desired...)
		}
		o = append(o, exprResult)
	}

	if len(desired) > 0 {
		return NewObjectDeferred(el, desired...), nil
	}

	return NewObjectList(o...), nil
}

func (i *Interpreter) VisitExpressionListMaker(elm ast.ExpressionListMaker) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)
	/*
		[elm.Expr for elm.For.Idents[] in elm.For.Expr]
		[elm.Expr for elm.For.Idents[] in elm.For.Expr for Iter.For]
		[elm.Expr for elm.For.Idents[] in elm.For.Expr if Iter.If]
	*/

	return i.evalDataListFor(elm.For, elm.Expr)
}

func (i *Interpreter) evalDataListFor(dlf *ast.DataListFor, expr ast.Expression) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	o, err := dlf.Expr.Accept(i)
	if err != nil {
		return nil, err
	}

	var values []Object
	switch o := o.(type) {
	case *ObjectList:
		values = o.Items
	case *ObjectTuple:
		values = o.Items
	case *ObjectDeferred:
		return o, nil
	default:
		return nil, fmt.Errorf("for over a %T, expecting *ObjectList, or *ObjectTuple", o)
	}

	if dlf.Iter != nil {
		panic("dlf.Iter")
	}

	if len(dlf.Idents) != 1 {
		// I think we just need to unpack each element
		panic("dlf.Idents != 1")
	}

	var newValues []Object
	for _, value := range values {
		err := i.withScope(func() error {
			if err := i.Scope.Set(dlf.Idents[0], value); err != nil {
				return err
			} else if newValue, err := r(expr.Accept(i)); err != nil {
				return err
			} else {
				newValues = append(newValues, newValue)
			}
			return nil
		})
		if err != nil {
			return
		}
	}
	return NewObjectList(newValues...), nil

}

func (i *Interpreter) VisitExpressionLiteral(el ast.ExpressionLiteral) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	switch v := el.Value.(type) {
	case int:
		return NewObjectInt(v), nil
	case float64:
		return NewObjectDouble(v), nil
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

func (i *Interpreter) VisitExpressionLogical(el ast.ExpressionLogical) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("ExpressionLogical")
}

func (i *Interpreter) VisitExpressionMember(em ast.ExpressionMember) (returnValue any, errOut error) {
	defer i.trace("Member: %s", em.Identifier)(&returnValue, &errOut)

	if obj, err := r(em.Expr.Accept(i)); err != nil {
		return nil, err
	} else {
		return obj.Member(i, obj, em.Identifier)
	}
}

func (i *Interpreter) VisitExpressionSubscript(es ast.ExpressionSubscript) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("ExpressionSubscript")
}

func (i *Interpreter) VisitExpressionTernary(et ast.ExpressionTernary) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	if cond, err := r(et.Cond.Accept(i)); err != nil {
		return nil, err
	} else if ternary, ok := cond.(FlowTernary); ok {
		// We explicitly pass cond, because the interface may be (is) being satisfied
		// by an embedded field, and still needs to know the containing object.  That
		// is, cond doesn't necessarily equal ternary.
		return ternary.VisitExpressionTernary(i, et.Left, cond, et.Right)
	} else {
		return nil, fmt.Errorf("condition is %T not FlowTernary", cond)
	}
}

func (i *Interpreter) VisitExpressionTuple(et ast.ExpressionTuple) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	var o []Object
	var desired []string
	for _, expr := range et.Expressions {
		exprResult, err := r(expr.Accept(i))
		if err != nil {
			return nil, err
		} else if od, ok := exprResult.(*ObjectDeferred); ok {
			desired = append(desired, od.desired...)
		}
		o = append(o, exprResult)
	}

	if len(desired) > 0 {
		return NewObjectDeferred(et, desired...), nil
	}

	return NewObjectTuple(o...), nil
}

func (i *Interpreter) VisitExpressionUnary(eu ast.ExpressionUnary) (returnValue any, errOut error) {
	defer i.trace()(&returnValue, &errOut)

	panic("ExpressionUnary")
}

func (i *Interpreter) VisitExpressionVariable(ev ast.ExpressionVariable) (returnValue any, errOut error) {
	defer i.trace("Identifier: %s", ev.Identifier)(&returnValue, &errOut)

	return i.Scope.Get(ev.Identifier)
}
