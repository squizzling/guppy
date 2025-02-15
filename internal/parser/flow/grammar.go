package flow

import (
	"fmt"

	"guppy/internal/parser/ast"
	"guppy/internal/parser/parser"
	"guppy/internal/parser/tokenizer"
)

var magicNames = map[tokenizer.TokenType]string{
	tokenizer.TokenTypeAmper:          "__and__",
	tokenizer.TokenTypeAnd:            "__binary_and__",
	tokenizer.TokenTypeCaret:          "__xor__",
	tokenizer.TokenTypeGreaterGreater: "__rshift__",
	tokenizer.TokenTypeLessLess:       "__lshift__",
	tokenizer.TokenTypeMinus:          "__sub__",
	tokenizer.TokenTypeOr:             "__binary_or__",
	tokenizer.TokenTypePipe:           "__or__",
	tokenizer.TokenTypePlus:           "__add__",
	tokenizer.TokenTypeSlash:          "__truediv__",
	tokenizer.TokenTypeStar:           "__mul__",
}

func ParseProgram(p *parser.Parser) (ast.StatementProgram, *parser.ParseError) {
	/*
	   program
	     : ( NEWLINE | statement )* EOF
	     ;
	*/
	var statements []ast.Statement
	for !p.Match(tokenizer.TokenTypeEOF) {
		if p.Match(tokenizer.TokenTypeNewLine) {
			continue
		}
		if statement, err := parseStatement(p); err != nil {
			return ast.StatementProgram{}, parser.FailErr(err)
		} else {
			statements = append(statements, statement)
		}
	}

	return ast.NewStatementProgram(ast.NewStatementList(statements)), nil
}

func parseStatement(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	/*
	   statement
	     : simple_statement
	     | compound_statement
	     ;
	*/
	if p.PeekMatch(0, tokenizer.TokenTypeIf, tokenizer.TokenTypeDef, tokenizer.TokenTypeAt) {
		return parser.Wrap(parseCompoundStatement(p))
	} else {
		return parser.Wrap(parseSimpleStatement(p))
	}
}

func parseCompoundStatement(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	/*
		compound_statement
		  : if_statement
		  | function_definition
		  | decorated
		  ;
	*/
	switch p.Peek(0).Type {
	case tokenizer.TokenTypeIf:
		return parser.Wrap(parseIf(p))
	case tokenizer.TokenTypeDef:
		return parser.Wrap(parseFunctionDefinition(p))
	case tokenizer.TokenTypeAt:
		return parser.Wrap(parseDecorator(p))
	default:
		return nil, parser.FailMsgf("expected IF, DEF, or '@'")
	}
}

func parseSimpleStatement(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	/*
		simple_statement
		  :  small_statement ( ';' small_statement )* ';'? (NEWLINE | EOF)
		  ;
	*/
	var smallStatements []ast.Statement
	for {
		smallStatement, err := parseSmallStatement(p)
		if err != nil {
			return nil, parser.FailErr(err)
		}

		smallStatements = append(smallStatements, smallStatement)

		if p.Match(tokenizer.TokenTypeNewLine) || p.Match(tokenizer.TokenTypeEOF) {
			break
		}

		if !p.Match(tokenizer.TokenTypeSemiColon) {
			return nil, parser.FailMsgf("expecting NEWLINE, EOF, or ';' following statement")
		}

		if p.Match(tokenizer.TokenTypeNewLine) || p.Match(tokenizer.TokenTypeEOF) {
			break
		}
	}

	if len(smallStatements) == 1 {
		return smallStatements[0], nil
	} else {
		return ast.NewStatementList(smallStatements), nil
	}
}

func parseSmallStatement(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	/*
		small_statement
		  : expr_statement
		  | flow_statement
		  | import_statement
		  | assert_statement
		  ;
	*/
	if p.Match(tokenizer.TokenTypeReturn) {
		return parser.Wrap(parseReturn(p))
	} else if t, ok := p.Capture(tokenizer.TokenTypeImport, tokenizer.TokenTypeFrom); ok {
		return parser.Wrap(parseImport(p, t.Type))
	} else if p.Match(tokenizer.TokenTypeAssert) {
		return parser.Wrap(parseAssert(p))
	} else {
		return parser.Wrap(parseExpressionStatement(p))
	}
}

func parseIf(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	return nil, parser.FailMsgf("if not supported")
}

func parseFunctionDefinition(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	/*
		function_definition
		  : DEF ID parameters ':' suite
		  ;
	*/
	return nil, parser.FailMsgf("parseFunctionDefinition not supported")
}

func parseParameters(p *parser.Parser) (*ast.DataParameterList, *parser.ParseError) {
	/*
		parameters
		  : OPEN_PAREN var_args_list? CLOSE_PAREN
		  ;
	*/
	return nil, parser.FailMsgf("parameters not supported")
}

func parseVarArgsList(p *parser.Parser) (*ast.DataParameterList, *parser.ParseError) {
	/*
		var_args_list
		  : var_args_list_param_def ( ',' var_args_list_param_def )* ( ','  ( (var_args_star_param (',' var_args_list_param_def)* (',' var_args_kws_param)?)
		                                                                    | var_args_kws_param
		                                                                    )?
		                                                             )?
		  | var_args_star_param (',' var_args_list_param_def)* (',' var_args_kws_param)?
		  | var_args_kws_param
		  ;
	*/
	params := []*ast.DataParameter{}

	if p.PeekMatch(0, tokenizer.TokenTypeIdentifier) {
		if param, err := parseVarArgsListParamDef(p); err != nil {
			return nil, parser.FailErr(err)
		} else {
			params = append(params, param)
			for {
				if p.Match(tokenizer.TokenTypeComma) {
					if p.PeekMatch(0, tokenizer.TokenTypeIdentifier) {
						if param, err = parseVarArgsListParamDef(p); err != nil {
							return nil, parser.FailErr(err)
						} else {
							params = append(params, param)
						}
					} else if p.PeekMatch(0, tokenizer.TokenTypeStar, tokenizer.TokenTypeStarStar) {
						break
					} else {
						return nil, parser.FailMsgf("expecting IDENTIFIER, '*', or STARSTAR are ',' in parsevarArgsList, found: %s", p.Peek(0).Type)
					}
				} else {
					return &ast.DataParameterList{List: params}, nil
				}
			}
		}
	}

	if p.PeekMatch(0, tokenizer.TokenTypeStar) {
		if param, err := parseVarArgsStarParam(p); err != nil {
			return nil, parser.FailErr(err)
		} else {
			params = append(params, param)
			for {
				if p.Match(tokenizer.TokenTypeComma) {
					if p.PeekMatch(0, tokenizer.TokenTypeIdentifier) {
						if param, err = parseVarArgsListParamDef(p); err != nil {
							return nil, parser.FailErr(err)
						} else {
							params = append(params, param)
						}
					} else if p.PeekMatch(0, tokenizer.TokenTypeStarStar) {
						break
					} else {
						return nil, parser.FailMsgf("expecting IDENTIFIER or STARSTAR after ',' in parseVarArgsList, found: %s", p.Peek(0).Type)
					}
				} else {
					return &ast.DataParameterList{List: params}, nil
				}
			}
		}
	}

	if p.PeekMatch(0, tokenizer.TokenTypeStarStar) {
		if param, err := parseVarArgsKwsParam(p); err != nil {
			return nil, parser.FailErr(err)
		} else {
			params = append(params, param)
		}
	}

	if len(params) == 0 { // If we have nothing, then no arms matched.
		return nil, parser.FailMsgf("expecting IDENTIFIER, '*', or STARSTAR in parseVarArgsList, found: %s", p.Peek(0).Type)
	}

	return &ast.DataParameterList{List: params}, nil
}

func parseVarArgsListParamDef(p *parser.Parser) (*ast.DataParameter, *parser.ParseError) {
	/*
		var_args_list_param_def
		  : var_args_list_param_name ( '=' test)?
		  ;
	*/
	if param, err := parseVarArgsListParamName(p); err != nil {
		return nil, parser.FailErr(err)
	} else if !p.Match(tokenizer.TokenTypeEqual) {
		return param, nil
	} else if expr, err := parseTest(p); err != nil {
		return nil, parser.FailErr(err)
	} else {
		param.Default = expr
		return param, nil
	}
}

func parseVarArgsListParamName(p *parser.Parser) (*ast.DataParameter, *parser.ParseError) {
	/*
	   var_args_list_param_name
	     : ID param_type?
	     ;
	*/
	if t, ok := p.Capture(tokenizer.TokenTypeIdentifier); !ok {
		return nil, parser.FailMsgf("expecting IDENTIFIER in parseVarArgsListParamName, found: %s", p.Peek(0).Type)
	} else if !p.PeekMatch(0, tokenizer.TokenTypeColon) {
		return &ast.DataParameter{
			Name: t.Lexeme,
		}, nil
	} else if param, err := parseParamType(p); err != nil {
		return nil, parser.FailErr(err)
	} else {
		param.Name = t.Lexeme
		return param, nil
	}
}

func parseParamType(p *parser.Parser) (*ast.DataParameter, *parser.ParseError) {
	/*
	   param_type
	     : ':' ID
	     ;
	*/
	if !p.Match(tokenizer.TokenTypeColon) {
		return nil, parser.FailMsgf("expecting ':' in parseParamType, found: %s", p.Peek(0).Type)
	} else if t, ok := p.Capture(tokenizer.TokenTypeIdentifier); !ok {
		return nil, parser.FailMsgf("expecting IDENTIFIER in parseParamType, found: %s", p.Peek(0).Type)
	} else {
		return &ast.DataParameter{
			Type: t.Lexeme,
		}, nil
	}
}

func parseVarArgsStarParam(p *parser.Parser) (*ast.DataParameter, *parser.ParseError) {
	/*
		var_args_star_param
		  :  STAR var_args_list_param_name?
		  ;
	*/
	if !p.Match(tokenizer.TokenTypeStar) {
		return nil, parser.FailMsgf("expecting '*' in parseVarArgsStarParam, found: %s", p.Peek(0).Type)
	} else if p.PeekMatch(0, tokenizer.TokenTypeIdentifier) {
		if param, err := parseVarArgsListParamName(p); err != nil {
			return nil, parser.FailErr(err)
		} else {
			param.StarArg = true
			return param, nil
		}
	} else {
		return &ast.DataParameter{StarArg: true}, nil
	}
}

func parseVarArgsKwsParam(p *parser.Parser) (*ast.DataParameter, *parser.ParseError) {
	/*
		var_args_kws_param
		  : POWER var_args_list_param_name
		  ;
	*/
	if !p.Match(tokenizer.TokenTypeStarStar) {
		return nil, parser.FailMsgf("expecting STARSTAR in parseVarArgsKwsParam, found: %s", p.Peek(0).Type)
	} else if param, err := parseVarArgsListParamName(p); err != nil {
		return nil, parser.FailErr(err)
	} else {
		param.KeywordArg = true
		return param, nil
	}
}

func parseSuite(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	/*
		suite
		  : simple_statement
		  | NEWLINE INDENT statement+ DEDENT
		  ;
	*/
	return nil, parser.FailMsgf("suite not supported")
}

func parseDecorator(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	return nil, parser.FailMsgf("decorator not supported")
}

func parseReturn(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	return nil, parser.FailMsgf("return not supported")
}

func parseImport(p *parser.Parser, tt tokenizer.TokenType) (ast.Statement, *parser.ParseError) {
	return nil, parser.FailMsgf("import not supported")
}

func parseAssert(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	return nil, parser.FailMsgf("assert not supported")
}

func parseExpressionStatement(p *parser.Parser) (ast.Statement, *parser.ParseError) {
	/*
		expr_statement
		  : (id_list ASSIGN)? testlist
		  ;
	*/
	var assignList []string
	if p.PeekMatch(1, tokenizer.TokenTypeComma, tokenizer.TokenTypeEqual) {
		var err *parser.ParseError
		if assignList, err = parseIdList(p); err != nil {
			return nil, parser.FailErr(err)
		}
		if !p.Match(tokenizer.TokenTypeEqual) {
			return nil, parser.FailMsgf("expecting '=' after identifier")
		}
	}

	if exprs, err := parseTestList(p); err != nil {
		return nil, parser.FailErr(err)
	} else {
		return ast.NewStatementExpression(assignList, exprs), nil
	}
}

func parseIdList(p *parser.Parser) ([]string, *parser.ParseError) {
	/*
		id_list
		  : ID (',' ID)* ','?
		  ;
	*/
	t, ok := p.Capture(tokenizer.TokenTypeIdentifier)
	if !ok {
		return nil, parser.FailMsgf("expecting ID, found %s", p.Peek(0).Type)
	}
	idList := []string{t.Lexeme}
	for p.Match(tokenizer.TokenTypeComma) { // Read a comma
		if t, ok := p.Capture(tokenizer.TokenTypeIdentifier); ok { // If there's an identifier after it,
			idList = append(idList, t.Lexeme) // Keep it
		} else { // If there's no identifier...
			break // It was a dangling comma and that's ok
		}
	}
	return idList, nil
}

func isAtomStart(p *parser.Parser) bool {
	atomTokens := []tokenizer.TokenType{
		tokenizer.TokenTypePlus,
		tokenizer.TokenTypeMinus,
		tokenizer.TokenTypeTilde,
		tokenizer.TokenTypeLeftParen,
		tokenizer.TokenTypeLeftBrace,
		tokenizer.TokenTypeLeftSquare,
		tokenizer.TokenTypeIdentifier,
		tokenizer.TokenTypeLambda,
		tokenizer.TokenTypeInt,
		tokenizer.TokenTypeFloat,
		tokenizer.TokenTypeString,
		tokenizer.TokenTypeNone,
		tokenizer.TokenTypeTrue,
		tokenizer.TokenTypeFalse,
	}
	return p.PeekMatch(0, atomTokens...)
}

func parseTestList(p *parser.Parser) (ast.ExpressionList, *parser.ParseError) {
	/*
		testlist
		  : test (COMMA test)* COMMA?
		  ;
	*/

	var exprList []ast.Expression
	for {
		if expr, err := parseTest(p); err != nil {
			return ast.ExpressionList{}, parser.FailErr(err)
		} else {
			exprList = append(exprList, expr)
			if !p.Match(tokenizer.TokenTypeComma) {
				return ast.NewExpressionList(exprList, len(exprList) > 1), nil
			} else if !isAtomStart(p) {
				return ast.NewExpressionList(exprList, true), nil
			} else {
				// loop
			}
		}
	}
}

func parseTest(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		test
		  : or_test (IF or_test ELSE test)?
		  | lambdef
		  ;
	*/
	if p.PeekMatch(0, tokenizer.TokenTypeLambda) {
		return parser.Wrap(parseLambda(p))
	} else if exprLeft, err := parseOrTest(p); err != nil {
		return nil, parser.FailErr(err)
	} else if p.Match(tokenizer.TokenTypeIf) {

		if exprCond, err := parseOrTest(p); err != nil {
			return nil, parser.FailErr(err)
		} else if !p.Match(tokenizer.TokenTypeElse) {
			return nil, parser.FailMsgf("expecting 'ELSE' in test clause, found %s", p.Peek(0).Type)
		} else if exprRight, err := parseTest(p); err != nil {
			return nil, parser.FailErr(err)
		} else {
			return ast.NewExpressionTernary(exprLeft, exprCond, exprRight), nil
		}

	} else {
		return exprLeft, nil
	}
}

func parseLambda(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	return nil, parser.FailMsgf("lambda not supported")
}

func parseBinary(
	p *parser.Parser,
	next func() (ast.Expression, *parser.ParseError),
	tokens ...tokenizer.TokenType,
) (ast.Expression, *parser.ParseError) {
	leftExpression, err := next()
	if err != nil {
		return nil, parser.FailErrSkip(err, "", 1)
	}

	for op, ok := p.Capture(tokens...); ok; op, ok = p.Capture(tokens...) {
		rightExpression, err := next()
		if err != nil {
			return nil, parser.FailErrSkip(err, "", 1)
		}
		leftExpression = ast.NewExpressionBinary(leftExpression, op, rightExpression)
	}
	return leftExpression, nil
}

func parseBinaryCall(
	p *parser.Parser,
	next func(p *parser.Parser) (ast.Expression, *parser.ParseError),
	tokens ...tokenizer.TokenType,
) (ast.Expression, *parser.ParseError) {
	leftExpression, err := next(p)
	if err != nil {
		return nil, parser.FailErrSkip(err, "", 1)
	}

	for op, ok := p.Capture(tokens...); ok; op, ok = p.Capture(tokens...) {
		member, ok := magicNames[op.Type]
		if !ok {
			return nil, parser.FailErrSkip(fmt.Errorf("unrecognized tokenType %s", op.Type), "", 1)
		}
		rightExpression, err := next(p)
		if err != nil {
			return nil, parser.FailErrSkip(err, "", 1)
		}
		leftExpression = ast.NewExpressionCall(
			ast.NewExpressionMember(leftExpression, member),
			[]ast.DataArgument{
				{
					Assign: "right",
					Expr:   rightExpression,
				},
			},
			nil,
			nil,
		)
	}
	return leftExpression, nil
}

func parseOrTest(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		or_test
		  : and_test (OR and_test)*
		  ;
	*/
	return parseBinaryCall(p, parseAndTest, tokenizer.TokenTypeOr)
}

func parseAndTest(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		and_test
		  : not_test (AND not_test)*
		  ;
	*/
	return parseBinaryCall(p, parseNotTest, tokenizer.TokenTypeAnd)
}

func parseNotTest(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		not_test
		  : NOT not_test
		  | comparison
		  ;
	*/
	if p.Match(tokenizer.TokenTypeNot) {
		if e, err := parseNotTest(p); err != nil {
			return nil, parser.FailErr(err)
		} else {
			return ast.NewExpressionUnary(tokenizer.TokenTypeNot, e), nil
		}
	} else {
		return parser.Wrap(parseComparison(p))
	}
}

func parseComparison(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		comparison
		  : expr ((LESS_THAN | LT_EQ | EQUALS | NOT_EQ_1 | NOT_EQ_2 | GREATER_THAN | GT_EQ | IS | IS NOT) expr)*
		  ;
	*/
	leftExpression, err := parseExpression(p)
	if err != nil {
		return nil, parser.FailErrSkip(err, "", 1)
	}

	for op, ok := p.Capture(
		tokenizer.TokenTypeLess,
		tokenizer.TokenTypeLessEqual,
		tokenizer.TokenTypeEqualEqual,
		tokenizer.TokenTypeBangEqual,
		tokenizer.TokenTypeLessGreater,
		tokenizer.TokenTypeGreater,
		tokenizer.TokenTypeGreaterEqual,
		tokenizer.TokenTypeIs,
	); ok; op, ok = p.Capture(
		tokenizer.TokenTypeLess,
		tokenizer.TokenTypeLessEqual,
		tokenizer.TokenTypeEqualEqual,
		tokenizer.TokenTypeBangEqual,
		tokenizer.TokenTypeLessGreater,
		tokenizer.TokenTypeGreater,
		tokenizer.TokenTypeGreaterEqual,
		tokenizer.TokenTypeIs,
	) {
		if op.Type == tokenizer.TokenTypeIs {
			if p.Match(tokenizer.TokenTypeNot) {
				op.Type = tokenizer.TokenTypeIsNot
			}
		}
		rightExpression, err := parseExpression(p)
		if err != nil {
			return nil, parser.FailErrSkip(err, "", 1)
		}
		leftExpression = ast.NewExpressionBinary(leftExpression, op, rightExpression)
	}
	return leftExpression, nil
}

func parseExpression(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		expr
		  : xor_expr ( '|' xor_expr )*
		  ;
	*/
	return parseBinaryCall(p, parseXorExpr, tokenizer.TokenTypePipe)
}

func parseXorExpr(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		xor_expr
		  :  and_expr ( '^' and_expr )*
		  ;
	*/
	return parseBinaryCall(p, parseAndExpr, tokenizer.TokenTypeCaret)
}

func parseAndExpr(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		and_expr
		  : shift_expr ( '&' shift_expr )*
		  ;
	*/
	return parseBinaryCall(p, parseShiftExpr, tokenizer.TokenTypeAmper)
}

func parseShiftExpr(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		shift_expr
		  : arith_expr ( '<<' arith_expr
		               | '>>' arith_expr
		               )*
		  ;
	*/
	return parseBinaryCall(p, parseArithExpr, tokenizer.TokenTypeLessLess, tokenizer.TokenTypeGreaterGreater)
}

func parseArithExpr(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		arith_expr
		  : term ((ADD | MINUS) term)*
		  ;
	*/
	return parseBinaryCall(p, parseTerm, tokenizer.TokenTypePlus, tokenizer.TokenTypeMinus)
}

func parseTerm(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		term
		  : factor ((STAR | DIV) factor)*
		  ;
	*/
	return parseBinaryCall(p, parseFactor, tokenizer.TokenTypeStar, tokenizer.TokenTypeSlash)
}

func parseFactor(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		factor
		  : (ADD | MINUS | NOT_OP) factor
		  | power
		  ;
	*/
	if t, ok := p.Capture(tokenizer.TokenTypePlus, tokenizer.TokenTypeMinus, tokenizer.TokenTypeTilde); ok {
		if expr, err := parseFactor(p); err != nil {
			return nil, parser.FailErr(err)
		} else {
			return ast.NewExpressionUnary(t.Type, expr), nil
		}
	}
	return parser.Wrap(parsePower(p))
}

func parsePower(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		power
		  : atom_expr (POWER factor)?
		  ;
	*/
	if atom, err := parseAtomExpr(p); err != nil {
		return nil, parser.FailErr(err)
	} else if t, ok := p.Capture(tokenizer.TokenTypeCaret); !ok {
		return atom, nil
	} else if factor, err := parseFactor(p); err != nil {
		return nil, parser.FailErr(err)
	} else {
		return ast.NewExpressionBinary(atom, t, factor), nil
	}
}

func parseAtomExpr(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
			atom_expr
		  : atom trailer*
		  ;
	*/
	if expr, err := parseAtom(p); err != nil {
		return nil, parser.FailErr(err)
	} else {
		for isTrailerStart(p) {
			if expr, err = parseTrailer(p, expr); err != nil {
				return nil, parser.FailErr(err)
			}
		}
		return expr, nil
	}
}

func isTrailerStart(p *parser.Parser) bool {
	trailerTokens := []tokenizer.TokenType{
		tokenizer.TokenTypeLeftParen,
		tokenizer.TokenTypeLeftSquare,
		tokenizer.TokenTypeDot,
	}
	return p.PeekMatch(0, trailerTokens...)
}

func parseTrailer(p *parser.Parser, expr ast.Expression) (ast.Expression, *parser.ParseError) {
	/*
		trailer
		  : OPEN_PAREN actual_args? CLOSE_PAREN
		  | OPEN_BRACK subscript CLOSE_BRACK
		  | DOT ID
		  ;
	*/

	var err *parser.ParseError
	if p.Match(tokenizer.TokenTypeLeftParen) {
		if isActualArgsStart(p) {
			if expr, err = parseActualArgs(p, expr); err != nil {
				return nil, parser.FailErr(err)
			}
		} else {
			expr = ast.NewExpressionCall(expr, nil, nil, nil)
		}
		if t, ok := p.Capture(tokenizer.TokenTypeRightParen); !ok {
			return nil, parser.FailMsgf("expecting ')' after args found %s", t.Type)
		} else {
			return expr, nil
		}
	} else if p.Match(tokenizer.TokenTypeLeftSquare) {
		if expr, err = parseSubscript(p, expr); err != nil {
			return nil, parser.FailErr(err)
		} else if !p.Match(tokenizer.TokenTypeRightSquare) {
			return nil, parser.FailMsgf("expecting ']' after subscript")
		} else {
			return expr, nil
		}
	} else if p.Match(tokenizer.TokenTypeDot) {
		if ident, ok := p.Capture(tokenizer.TokenTypeIdentifier); ok {
			return ast.NewExpressionMember(expr, ident.Lexeme), nil
		} else {
			return nil, parser.FailMsgf("expecting IDENT after '.'")
		}
	} else {
		// We should be checking this before reaching here, so this shouldn't occur
		return nil, parser.FailMsgf("expecting '(', '[', or '.' to start trailer")
	}
}

func isActualArgsStart(p *parser.Parser) bool {
	actualArgsTokens := []tokenizer.TokenType{
		tokenizer.TokenTypeStar,
		tokenizer.TokenTypeStarStar,
	}
	return p.PeekMatch(0, actualArgsTokens...) || isAtomStart(p)
}

func parseActualArgs(p *parser.Parser, expr ast.Expression) (ast.Expression, *parser.ParseError) {
	/*
		actual_args
		  : (argument COMMA)* ( argument COMMA?
		                      | actual_star_arg (COMMA argument)* (COMMA actual_kws_arg)?
		                      | actual_kws_arg
		                      )
		  ;
	*/
	var args []ast.DataArgument
	haveFirstNamedArg := false
	for {
		if isAtomStart(p) { // TODO: Make sure isAtomStart does what we want
			if arg, err := parseArgument(p); err != nil {
				return nil, parser.FailErr(err)
			} else {
				if arg.Assign != "" {
					haveFirstNamedArg = true
				} else if haveFirstNamedArg {
					// No assign, but we've had a named arg
					return nil, parser.FailMsgf("positional argument follows keyword argument")
				}
				args = append(args, arg)
				if !p.Match(tokenizer.TokenTypeComma) {
					return ast.NewExpressionCall(expr, args, nil, nil), nil
				}
			}
		} else if p.Match(tokenizer.TokenTypeStar) {
			starArgs, err := parseActualStarArg(p)
			if err != nil {
				return nil, parser.FailErr(err)
			}
			for {
				// TODO: COMMA in the inner loop
				if isAtomStart(p) {
					if arg, err := parseArgument(p); err != nil {
						return nil, parser.FailErr(err)
					} else {
						args = append(args, arg)
					}
				} else if p.Match(tokenizer.TokenTypeStarStar) {
					if keywordArgs, err := parseActualKeywordArg(p); err != nil {
						return nil, parser.FailErr(err)
					} else {
						return ast.NewExpressionCall(expr, args, starArgs, keywordArgs), nil
					}
				} else {
					return ast.NewExpressionCall(expr, args, starArgs, nil), nil
				}
			}
		} else if p.Match(tokenizer.TokenTypeStarStar) {
			if keywordArgs, err := parseActualKeywordArg(p); err != nil {
				return nil, parser.FailErr(err)
			} else {
				return ast.NewExpressionCall(expr, args, nil, keywordArgs), nil
			}
		} else {
			return ast.NewExpressionCall(expr, args, nil, nil), nil
		}
	}
}

func parseActualStarArg(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		actual_star_arg
		  :  STAR test
		  ;
	*/
	return parser.Wrap(parseTest(p))
}

func parseActualKeywordArg(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
	   actual_kws_arg
	     : POWER test
	     ;
	*/
	return parser.Wrap(parseTest(p))
}

func parseArgument(p *parser.Parser) (ast.DataArgument, *parser.ParseError) {
	/*
	   argument
	     : test (comp_for)? | (ID ASSIGN)? test
	     ;
	*/
	if p.PeekMatch(0, tokenizer.TokenTypeIdentifier) && p.PeekMatch(1, tokenizer.TokenTypeEqual) {
		tokenId, _ := p.Capture(tokenizer.TokenTypeIdentifier) // Already validated
		p.Match(tokenizer.TokenTypeEqual)                      // Already validated
		if test, err := parseTest(p); err != nil {
			return ast.DataArgument{}, parser.FailErr(err)
		} else {
			return ast.NewDataArgument(tokenId.Lexeme, test), nil
		}
	} else {
		if expr, err := parseTest(p); err != nil {
			return ast.DataArgument{}, parser.FailErr(err)
		} else if !p.Match(tokenizer.TokenTypeFor) {
			return ast.NewDataArgument("", expr), nil
		} else if expr, err = parseCompFor(p, expr); err != nil {
			return ast.DataArgument{}, parser.FailErr(err)
		} else {
			return ast.NewDataArgument("", expr), nil
		}
	}
}

func parseCompFor(p *parser.Parser, expr ast.Expression) (ast.Expression, *parser.ParseError) {
	/*
		comp_for
		  : FOR id_list IN or_test (comp_iter)?
		  ;
	*/
	return nil, parser.FailMsgf("compFor not implemented")
}

func parseSubscript(p *parser.Parser, expr ast.Expression) (ast.Expression, *parser.ParseError) {
	/*
		subscript
		  : test
		  | test? COLON test?
		  ;
	*/
	return nil, parser.FailMsgf("subscript not implemented")
}

func parseAtom(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		atom
		  : list_expr
		  | tuple_expr
		  | dict_expr
		  | ID
		  | INT
		  | FLOAT
		  | STRING+
		  | NONE
		  | TRUE
		  | FALSE
		  ;
	*/
	if p.Match(tokenizer.TokenTypeLeftSquare) {
		return parser.Wrap(parseListExpr(p))
	} else if p.PeekMatch(0, tokenizer.TokenTypeLeftParen) {
		return parser.Wrap(parseTupleExpr(p))
	} else if p.PeekMatch(0, tokenizer.TokenTypeLeftBrace) {
		return parser.Wrap(parseDictExpr(p))
	} else if t, ok := p.Capture(tokenizer.TokenTypeIdentifier); ok {
		return ast.NewExpressionVariable(t.Lexeme), nil
	} else if t, ok := p.Capture(tokenizer.TokenTypeString); ok {
		return ast.NewExpressionLiteral(t.LiteralString), nil
	} else if t, ok := p.Capture(tokenizer.TokenTypeInt); ok {
		return ast.NewExpressionLiteral(t.LiteralInteger), nil
	} else if _, ok := p.Capture(tokenizer.TokenTypeFloat); ok {
		return ast.NewExpressionLiteral(t.LiteralFloat), nil
	} else if p.Match(tokenizer.TokenTypeNone) {
		return ast.NewExpressionLiteral(nil), nil
	} else if p.Match(tokenizer.TokenTypeTrue) {
		return ast.NewExpressionLiteral(true), nil
	} else if p.Match(tokenizer.TokenTypeFalse) {
		return ast.NewExpressionLiteral(false), nil
	} else {
		return nil, parser.FailMsgf("atom not supported: %s", p.Peek(0).Type)
	}
}

func parseListExpr(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		list_expr
		  : OPEN_BRACK list_maker? CLOSE_BRACK
		  ;
	*/
	if !isAtomStart(p) {
		if p.Match(tokenizer.TokenTypeRightSquare) {
			return ast.NewExpressionList(nil, false), nil
		}
		return nil, parser.FailMsgf("expecting atom in listexpr")
	} else if expr, err := parseListMaker(p); err != nil {
		return nil, parser.FailErr(err)
	} else if !p.Match(tokenizer.TokenTypeRightSquare) {
		return nil, parser.FailMsgf("expecting ']' after listMaker")
	} else {
		return expr, nil
	}
}

func parseListMaker(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		list_maker
		  : test ( list_for | (COMMA test)* (COMMA)? )
		  ;
	*/
	if expr, err := parseTest(p); err != nil {
		return nil, parser.FailErr(err)
	} else if p.Match(tokenizer.TokenTypeFor) {
		return nil, parser.FailMsgf("list_for not implemented")
	} else if p.Match(tokenizer.TokenTypeComma) {
		out := []ast.Expression{expr}
		for isAtomStart(p) {
			if expr, err = parseTest(p); err != nil {
				return nil, parser.FailErr(err)
			}
			out = append(out, expr)
			if !p.Match(tokenizer.TokenTypeComma) {
				break
			}
		}
		return ast.NewExpressionList(out, false), nil
	} else {
		return ast.NewExpressionList([]ast.Expression{expr}, false), nil
	}
}

func parseTupleExpr(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		tuple_expr
		  : OPEN_PAREN testlist_comp? CLOSE_PAREN
		  ;
	*/
	if !p.Match(tokenizer.TokenTypeLeftParen) {
		return nil, parser.FailMsgf("expecting '(' in tupleExpr")
	}

	if p.Match(tokenizer.TokenTypeRightParen) {
		return ast.NewExpressionList([]ast.Expression{}, true), nil
	}

	if expr, err := parseTestListComp(p); err != nil {
		return nil, parser.FailErr(err)
	} else if !p.Match(tokenizer.TokenTypeRightParen) {
		return nil, parser.FailMsgf("expecting ')' after tupleExpr")
	} else {
		return expr, nil
	}
}

func parseTestListComp(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		testlist_comp
		  : test (comp_for | (COMMA test)* COMMA?)
		  ;
	*/

	// () -> tuple of 0 (handled elsewhere)
	// (a) -> grouping, not a tuple
	// (a, ) -> tuple of 1
	// (a, b, ...) -> tuple of 2+

	expr, err := parseTest(p)
	if err != nil {
		return nil, parser.FailErr(err)
	}

	if p.PeekMatch(0, tokenizer.TokenTypeFor) {
		return parser.Wrap(parseCompFor(p, expr))
	}

	exprList := []ast.Expression{expr}
	isTuple := false
	for p.Match(tokenizer.TokenTypeComma) { // Read a comma
		isTuple = true
		if isAtomStart(p) { // If there's an expression after it...
			if expr, err = parseExpression(p); err == nil {
				exprList = append(exprList, expr) // Keep it
			} else {
				return nil, parser.FailErr(err) // Or fail
			}
		} else { // There was no expression, so it was a dangling comma
			break
		}
	}
	if len(exprList) == 1 && !isTuple {
		// This is a grouping, not a tuple
		return exprList[0], nil
	}
	return ast.NewExpressionList(exprList, isTuple), nil
}

func parseDictExpr(p *parser.Parser) (ast.Expression, *parser.ParseError) {
	/*
		dict_expr
		  : OPEN_BRACE (test ':' test ( ',' test ':' test )* ','?)? CLOSE_BRACE
		  ;
	*/

	// {}
	// {a:a}
	// {a:a,}
	// {a:a,a:a}
	// {a:a,a:a,}

	if !p.Match(tokenizer.TokenTypeLeftBrace) {
		return nil, parser.FailMsgf("expecting '{' in dictExpr, found: %s", p.Peek(0).Type)
	} else {
		var exprKeys []ast.Expression
		var exprValues []ast.Expression
		for !p.Match(tokenizer.TokenTypeRightBrace) {
			if exprKey, err := parseExpression(p); err != nil {
				return nil, parser.FailErr(err)
			} else if !p.Match(tokenizer.TokenTypeColon) {
				return nil, parser.FailMsgf("expecting ':' in dictExpr, found: %s", p.Peek(0).Type)
			} else if exprValue, err := parseExpression(p); err != nil {
				return nil, parser.FailErr(err)
			} else {
				exprKeys = append(exprKeys, exprKey)
				exprValues = append(exprValues, exprValue)
				if p.Match(tokenizer.TokenTypeComma) {
					continue
				}
			}
		}
		return ast.NewExpressionDict(exprKeys, exprValues), nil
	}
}
