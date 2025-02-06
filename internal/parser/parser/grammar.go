package parser

import (
	"fmt"

	"guppy/internal/parser/ast"
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

func (p *Parser) ParseProgram() (ast.StatementProgram, *ParseError) {
	/*
	   program
	     : ( NEWLINE | statement )* EOF
	     ;
	*/
	var statements []ast.Statement
	for !p.match(tokenizer.TokenTypeEOF) {
		if p.match(tokenizer.TokenTypeNewLine) {
			continue
		}
		if statement, err := p.parseStatement(); err != nil {
			return ast.StatementProgram{}, failErr(err)
		} else {
			statements = append(statements, statement)
		}
	}

	return ast.NewStatementProgram(ast.NewStatementList(statements)), nil
}

func (p *Parser) parseStatement() (ast.Statement, *ParseError) {
	/*
	   statement
	     : simple_statement
	     | compound_statement
	     ;
	*/
	if t, ok := p.capture(tokenizer.TokenTypeIf, tokenizer.TokenTypeDef, tokenizer.TokenTypeAt); ok {
		return wrap(p.parseCompoundStatement(t.Type))
	} else {
		return wrap(p.parseSimpleStatement())
	}
}

func (p *Parser) parseCompoundStatement(tt tokenizer.TokenType) (ast.Statement, *ParseError) {
	/*
		compound_statement
		  : if_statement
		  | function_definition
		  | decorated
		  ;
	*/
	switch tt {
	case tokenizer.TokenTypeIf:
		return wrap(p.parseIf())
	case tokenizer.TokenTypeDef:
		return wrap(p.parseDef())
	case tokenizer.TokenTypeAt:
		return wrap(p.parseDecorator())
	default:
		return nil, failMsgf("expected IF, DEF, or '@'")
	}
}

func (p *Parser) parseSimpleStatement() (ast.Statement, *ParseError) {
	/*
		simple_statement
		  :  small_statement ( ';' small_statement )* ';'? (NEWLINE | EOF)
		  ;
	*/
	var smallStatements []ast.Statement
	for {
		smallStatement, err := p.parseSmallStatement()
		if err != nil {
			return nil, failErr(err)
		}

		smallStatements = append(smallStatements, smallStatement)

		if p.match(tokenizer.TokenTypeNewLine) || p.match(tokenizer.TokenTypeEOF) {
			break
		}

		if !p.match(tokenizer.TokenTypeSemiColon) {
			return nil, failMsgf("expecting NEWLINE, EOF, or ';' following statement")
		}

		if p.match(tokenizer.TokenTypeNewLine) || p.match(tokenizer.TokenTypeEOF) {
			break
		}
	}

	if len(smallStatements) == 1 {
		return smallStatements[0], nil
	} else {
		return ast.NewStatementList(smallStatements), nil
	}
}

func (p *Parser) parseSmallStatement() (ast.Statement, *ParseError) {
	/*
		small_statement
		  : expr_statement
		  | flow_statement
		  | import_statement
		  | assert_statement
		  ;
	*/
	if p.match(tokenizer.TokenTypeReturn) {
		return wrap(p.parseReturn())
	} else if t, ok := p.capture(tokenizer.TokenTypeImport, tokenizer.TokenTypeFrom); ok {
		return wrap(p.parseImport(t.Type))
	} else if p.match(tokenizer.TokenTypeAssert) {
		return wrap(p.parseAssert())
	} else {
		return wrap(p.parseExpressionStatement())
	}
}

func (p *Parser) parseIf() (ast.Statement, *ParseError) {
	return nil, failMsgf("if not supported")
}

func (p *Parser) parseDef() (ast.Statement, *ParseError) {
	return nil, failMsgf("def not supported")
}

func (p *Parser) parseDecorator() (ast.Statement, *ParseError) {
	return nil, failMsgf("decorator not supported")
}

func (p *Parser) parseReturn() (ast.Statement, *ParseError) {
	return nil, failMsgf("return not supported")
}

func (p *Parser) parseImport(tt tokenizer.TokenType) (ast.Statement, *ParseError) {
	return nil, failMsgf("import not supported")
}

func (p *Parser) parseAssert() (ast.Statement, *ParseError) {
	return nil, failMsgf("assert not supported")
}

func (p *Parser) parseExpressionStatement() (ast.Statement, *ParseError) {
	/*
		expr_statement
		  : (id_list ASSIGN)? testlist
		  ;
	*/
	var assignList []string
	if p.peekMatch(0, tokenizer.TokenTypeIdentifier) && p.peekMatch(1, tokenizer.TokenTypeComma, tokenizer.TokenTypeEqual) {
		var err *ParseError
		if assignList, err = p.parseIdList(); err != nil {
			return nil, failErr(err)
		}
		if !p.match(tokenizer.TokenTypeEqual) {
			return nil, failMsgf("expecting '=' after identifier")
		}
	}

	if exprs, err := p.parseTestList(); err != nil {
		return nil, failErr(err)
	} else {
		return ast.NewStatementExpression(assignList, exprs), nil
	}
}

func (p *Parser) parseIdList() ([]string, *ParseError) {
	// TODO
	/*
		id_list
		  : ID (',' ID)* ','?
		  ;
	*/
	t, _ := p.capture(tokenizer.TokenTypeIdentifier)
	return []string{t.Lexeme}, nil
}

func (p *Parser) isAtomStart() bool {
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
	return p.peekMatch(0, atomTokens...)
}

func (p *Parser) parseTestList() (ast.ExpressionList, *ParseError) {
	/*
		testlist
		  : test (COMMA test)* COMMA?
		  ;
	*/

	var exprList []ast.Expression
	for {
		if expr, err := p.parseTest(); err != nil {
			return ast.ExpressionList{}, failErr(err)
		} else {
			exprList = append(exprList, expr)
			if !p.match(tokenizer.TokenTypeComma) {
				return ast.NewExpressionList(exprList, len(exprList) > 1), nil
			} else if !p.isAtomStart() {
				return ast.NewExpressionList(exprList, true), nil
			} else {
				// loop
			}
		}
	}
}

func (p *Parser) parseTest() (ast.Expression, *ParseError) {
	/*
		test
		  : or_test (IF or_test ELSE test)?
		  | lambdef
		  ;
	*/
	if p.match(tokenizer.TokenTypeLambda) {
		return wrap(p.parseLambda())
	} else if exprLeft, err := p.parseOrTest(); err != nil {
		return nil, failErr(err)
	} else if p.match(tokenizer.TokenTypeIf) {
		return nil, failMsgf("expr IF not supported")
	} else {
		return exprLeft, nil
	}
}

func (p *Parser) parseLambda() (ast.Expression, *ParseError) {
	return nil, failMsgf("lambda not supported")
}

func (p *Parser) parseBinary(
	next func() (ast.Expression, *ParseError),
	tokens ...tokenizer.TokenType,
) (ast.Expression, *ParseError) {
	leftExpression, err := next()
	if err != nil {
		return nil, failErrSkip(err, "", 1)
	}

	for op, ok := p.capture(tokens...); ok; op, ok = p.capture(tokens...) {
		rightExpression, err := next()
		if err != nil {
			return nil, failErrSkip(err, "", 1)
		}
		leftExpression = ast.NewExpressionBinary(leftExpression, op, rightExpression)
	}
	return leftExpression, nil
}

func (p *Parser) parseBinaryCall(
	next func() (ast.Expression, *ParseError),
	tokens ...tokenizer.TokenType,
) (ast.Expression, *ParseError) {
	leftExpression, err := next()
	if err != nil {
		return nil, failErrSkip(err, "", 1)
	}

	for op, ok := p.capture(tokens...); ok; op, ok = p.capture(tokens...) {
		member, ok := magicNames[op.Type]
		if !ok {
			return nil, failErrSkip(fmt.Errorf("unrecognized tokenType %s", op.Type), "", 1)
		}
		rightExpression, err := next()
		if err != nil {
			return nil, failErrSkip(err, "", 1)
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

func (p *Parser) parseOrTest() (ast.Expression, *ParseError) {
	/*
		or_test
		  : and_test (OR and_test)*
		  ;
	*/
	return p.parseBinaryCall(p.parseAndTest, tokenizer.TokenTypeOr)
}

func (p *Parser) parseAndTest() (ast.Expression, *ParseError) {
	/*
		and_test
		  : not_test (AND not_test)*
		  ;
	*/
	return p.parseBinaryCall(p.parseNotTest, tokenizer.TokenTypeAnd)
}

func (p *Parser) parseNotTest() (ast.Expression, *ParseError) {
	/*
		not_test
		  : NOT not_test
		  | comparison
		  ;
	*/
	if p.match(tokenizer.TokenTypeNot) {
		if e, err := p.parseNotTest(); err != nil {
			return nil, failErr(err)
		} else {
			return ast.NewExpressionUnary(tokenizer.TokenTypeNot, e), nil
		}
	} else {
		return wrap(p.parseComparison())
	}
}

func (p *Parser) parseComparison() (ast.Expression, *ParseError) {
	/*
		comparison
		  : expr ((LESS_THAN | LT_EQ | EQUALS | NOT_EQ_1 | NOT_EQ_2 | GREATER_THAN | GT_EQ | IS | IS NOT) expr)*
		  ;
	*/
	leftExpression, err := p.parseExpression()
	if err != nil {
		return nil, failErrSkip(err, "", 1)
	}

	for op, ok := p.capture(
		tokenizer.TokenTypeLess,
		tokenizer.TokenTypeLessEqual,
		tokenizer.TokenTypeEqualEqual,
		tokenizer.TokenTypeBangEqual,
		tokenizer.TokenTypeLessGreater,
		tokenizer.TokenTypeGreater,
		tokenizer.TokenTypeGreaterEqual,
		tokenizer.TokenTypeIs,
	); ok; op, ok = p.capture(
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
			if p.match(tokenizer.TokenTypeNot) {
				op.Type = tokenizer.TokenTypeIsNot
			}
		}
		rightExpression, err := p.parseExpression()
		if err != nil {
			return nil, failErrSkip(err, "", 1)
		}
		leftExpression = ast.NewExpressionBinary(leftExpression, op, rightExpression)
	}
	return leftExpression, nil
}

func (p *Parser) parseExpression() (ast.Expression, *ParseError) {
	/*
		expr
		  : xor_expr ( '|' xor_expr )*
		  ;
	*/
	return p.parseBinaryCall(p.parseXorExpr, tokenizer.TokenTypePipe)
}

func (p *Parser) parseXorExpr() (ast.Expression, *ParseError) {
	/*
		xor_expr
		  :  and_expr ( '^' and_expr )*
		  ;
	*/
	return p.parseBinaryCall(p.parseAndExpr, tokenizer.TokenTypeCaret)
}

func (p *Parser) parseAndExpr() (ast.Expression, *ParseError) {
	/*
		and_expr
		  : shift_expr ( '&' shift_expr )*
		  ;
	*/
	return p.parseBinaryCall(p.parseShiftExpr, tokenizer.TokenTypeAmper)
}

func (p *Parser) parseShiftExpr() (ast.Expression, *ParseError) {
	/*
		shift_expr
		  : arith_expr ( '<<' arith_expr
		               | '>>' arith_expr
		               )*
		  ;
	*/
	return p.parseBinaryCall(p.parseArithExpr, tokenizer.TokenTypeLessLess, tokenizer.TokenTypeGreaterGreater)
}

func (p *Parser) parseArithExpr() (ast.Expression, *ParseError) {
	/*
		arith_expr
		  : term ((ADD | MINUS) term)*
		  ;
	*/
	return p.parseBinaryCall(p.parseTerm, tokenizer.TokenTypePlus, tokenizer.TokenTypeMinus)
}

func (p *Parser) parseTerm() (ast.Expression, *ParseError) {
	/*
		term
		  : factor ((STAR | DIV) factor)*
		  ;
	*/
	return p.parseBinaryCall(p.parseFactor, tokenizer.TokenTypeStar, tokenizer.TokenTypeSlash)
}

func (p *Parser) parseFactor() (ast.Expression, *ParseError) {
	/*
		factor
		  : (ADD | MINUS | NOT_OP) factor
		  | power
		  ;
	*/
	if t, ok := p.capture(tokenizer.TokenTypePlus, tokenizer.TokenTypeMinus, tokenizer.TokenTypeTilde); ok {
		if expr, err := p.parseFactor(); err != nil {
			return nil, failErr(err)
		} else {
			return ast.NewExpressionUnary(t.Type, expr), nil
		}
	}
	return wrap(p.parsePower())
}

func (p *Parser) parsePower() (ast.Expression, *ParseError) {
	/*
		power
		  : atom_expr (POWER factor)?
		  ;
	*/
	if atom, err := p.parseAtomExpr(); err != nil {
		return nil, failErr(err)
	} else if t, ok := p.capture(tokenizer.TokenTypeCaret); !ok {
		return atom, nil
	} else if factor, err := p.parseFactor(); err != nil {
		return nil, failErr(err)
	} else {
		return ast.NewExpressionBinary(atom, t, factor), nil
	}
}

func (p *Parser) parseAtomExpr() (ast.Expression, *ParseError) {
	/*
			atom_expr
		  : atom trailer*
		  ;
	*/
	if expr, err := p.parseAtom(); err != nil {
		return nil, failErr(err)
	} else {
		for p.isTrailerStart() {
			if expr, err = p.parseTrailer(expr); err != nil {
				return nil, failErr(err)
			}
		}
		return expr, nil
	}
}

func (p *Parser) isTrailerStart() bool {
	trailerTokens := []tokenizer.TokenType{
		tokenizer.TokenTypeLeftParen,
		tokenizer.TokenTypeLeftSquare,
		tokenizer.TokenTypeDot,
	}
	return p.peekMatch(0, trailerTokens...)
}

func (p *Parser) parseTrailer(expr ast.Expression) (ast.Expression, *ParseError) {
	/*
		trailer
		  : OPEN_PAREN actual_args? CLOSE_PAREN
		  | OPEN_BRACK subscript CLOSE_BRACK
		  | DOT ID
		  ;
	*/

	var err *ParseError
	if p.match(tokenizer.TokenTypeLeftParen) {
		if p.isActualArgsStart() {
			if expr, err = p.parseActualArgs(expr); err != nil {
				return nil, failErr(err)
			}
		} else {
			expr = ast.NewExpressionCall(expr, nil, nil, nil)
		}
		if !p.match(tokenizer.TokenTypeRightParen) {
			return nil, failMsgf("expecting ')' after args found %s", p.tokens.Peek(0).Value().Type)
		} else {
			return expr, nil
		}
	} else if p.match(tokenizer.TokenTypeLeftSquare) {
		if expr, err = p.parseSubscript(expr); err != nil {
			return nil, failErr(err)
		} else if !p.match(tokenizer.TokenTypeRightSquare) {
			return nil, failMsgf("expecting ']' after subscript")
		} else {
			return expr, nil
		}
	} else if p.match(tokenizer.TokenTypeDot) {
		if ident, ok := p.capture(tokenizer.TokenTypeIdentifier); ok {
			return ast.NewExpressionMember(expr, ident.Lexeme), nil
		} else {
			return nil, failMsgf("expecting IDENT after '.'")
		}
	} else {
		// We should be checking this before reaching here, so this shouldn't occur
		return nil, failMsgf("expecting '(', '[', or '.' to start trailer")
	}
}

func (p *Parser) isActualArgsStart() bool {
	actualArgsTokens := []tokenizer.TokenType{
		tokenizer.TokenTypeStar,
		tokenizer.TokenTypeStarStar,
	}
	return p.peekMatch(0, actualArgsTokens...) || p.isAtomStart()
}

func (p *Parser) parseActualArgs(expr ast.Expression) (ast.Expression, *ParseError) {
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
		if p.isAtomStart() { // TODO: Make sure isAtomStart does what we want
			if arg, err := p.parseArgument(); err != nil {
				return nil, failErr(err)
			} else {
				if arg.Assign != "" {
					haveFirstNamedArg = true
				} else if haveFirstNamedArg {
					// No assign, but we've had a named arg
					return nil, failMsgf("positional argument follows keyword argument")
				}
				args = append(args, arg)
				if !p.match(tokenizer.TokenTypeComma) {
					return ast.NewExpressionCall(expr, args, nil, nil), nil
				}
			}
		} else if p.match(tokenizer.TokenTypeStar) {
			starArgs, err := p.parseActualStarArg()
			if err != nil {
				return nil, failErr(err)
			}
			for {
				// TODO: COMMA in the inner loop
				if p.isAtomStart() {
					if arg, err := p.parseArgument(); err != nil {
						return nil, failErr(err)
					} else {
						args = append(args, arg)
					}
				} else if p.match(tokenizer.TokenTypeStarStar) {
					if keywordArgs, err := p.parseActualKeywordArg(); err != nil {
						return nil, failErr(err)
					} else {
						return ast.NewExpressionCall(expr, args, starArgs, keywordArgs), nil
					}
				} else {
					return ast.NewExpressionCall(expr, args, starArgs, nil), nil
				}
			}
		} else if p.match(tokenizer.TokenTypeStarStar) {
			if keywordArgs, err := p.parseActualKeywordArg(); err != nil {
				return nil, failErr(err)
			} else {
				return ast.NewExpressionCall(expr, args, nil, keywordArgs), nil
			}
		} else {
			return ast.NewExpressionCall(expr, args, nil, nil), nil
		}
	}
}

func (p *Parser) parseActualStarArg() (ast.Expression, *ParseError) {
	/*
		actual_star_arg
		  :  STAR test
		  ;
	*/
	return wrap(p.parseTest())
}

func (p *Parser) parseActualKeywordArg() (ast.Expression, *ParseError) {
	/*
	   actual_kws_arg
	     : POWER test
	     ;
	*/
	return wrap(p.parseTest())
}

func (p *Parser) parseArgument() (ast.DataArgument, *ParseError) {
	/*
	   argument
	     : test (comp_for)? | (ID ASSIGN)? test
	     ;
	*/
	if p.peekMatch(0, tokenizer.TokenTypeIdentifier) && p.peekMatch(1, tokenizer.TokenTypeEqual) {
		tokenId, _ := p.capture(tokenizer.TokenTypeIdentifier) // Already validated
		p.match(tokenizer.TokenTypeEqual)                      // Already validated
		if test, err := p.parseTest(); err != nil {
			return ast.DataArgument{}, failErr(err)
		} else {
			return ast.NewDataArgument(tokenId.Lexeme, test), nil
		}
	} else {
		if expr, err := p.parseTest(); err != nil {
			return ast.DataArgument{}, failErr(err)
		} else if !p.match(tokenizer.TokenTypeFor) {
			return ast.NewDataArgument("", expr), nil
		} else if expr, err = p.parseCompFor(expr); err != nil {
			return ast.DataArgument{}, failErr(err)
		} else {
			return ast.NewDataArgument("", expr), nil
		}
	}
}

func (p *Parser) parseCompFor(expr ast.Expression) (ast.Expression, *ParseError) {
	return nil, failMsgf("compFor not implemented")
}

func (p *Parser) parseSubscript(expr ast.Expression) (ast.Expression, *ParseError) {
	/*
		subscript
		  : test
		  | test? COLON test?
		  ;
	*/
	return nil, failMsgf("subscript not implemented")
}

func (p *Parser) parseAtom() (ast.Expression, *ParseError) {
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
	if p.match(tokenizer.TokenTypeLeftSquare) {
		return wrap(p.parseListExpr())
	} else if p.match(tokenizer.TokenTypeLeftParen) {
		return wrap(p.parseTupleExpr())
	} else if p.match(tokenizer.TokenTypeLeftBrace) {
		return wrap(p.parseDictExpr())
	} else if t, ok := p.capture(tokenizer.TokenTypeIdentifier); ok {
		return ast.NewExpressionVariable(t.Lexeme), nil
	} else if t, ok := p.capture(tokenizer.TokenTypeString); ok {
		return ast.NewExpressionLiteral(t.LiteralString), nil
	} else if t, ok := p.capture(tokenizer.TokenTypeInt); ok {
		return ast.NewExpressionLiteral(t.LiteralInteger), nil
	} else if _, ok := p.capture(tokenizer.TokenTypeFloat); ok {
		return nil, failMsgf("float literals not supported")
	} else if p.match(tokenizer.TokenTypeNone) {
		return ast.NewExpressionLiteral(nil), nil
	} else if p.match(tokenizer.TokenTypeTrue) {
		return ast.NewExpressionLiteral(true), nil
	} else if p.match(tokenizer.TokenTypeFalse) {
		return ast.NewExpressionLiteral(false), nil
	} else {
		t := p.tokens.Peek(0).Value()
		return nil, failMsgf("atom not supported: %s", t.Type)
	}
}

func (p *Parser) parseListExpr() (ast.Expression, *ParseError) {
	/*
		list_expr
		  : OPEN_BRACK list_maker? CLOSE_BRACK
		  ;
	*/
	if !p.isAtomStart() {
		if p.match(tokenizer.TokenTypeRightSquare) {
			return ast.NewExpressionList(nil, false), nil
		}
		return nil, failMsgf("expecting atom in listexpr")
	} else if expr, err := p.parseListMaker(); err != nil {
		return nil, failErr(err)
	} else if !p.match(tokenizer.TokenTypeRightSquare) {
		return nil, failMsgf("expecting ']' after listMaker")
	} else {
		return expr, nil
	}
}

func (p *Parser) parseListMaker() (ast.Expression, *ParseError) {
	/*
		list_maker
		  : test ( list_for | (COMMA test)* (COMMA)? )
		  ;
	*/
	if expr, err := p.parseTest(); err != nil {
		return nil, failErr(err)
	} else if p.match(tokenizer.TokenTypeFor) {
		return nil, failMsgf("list_for not implemented")
	} else if p.match(tokenizer.TokenTypeComma) {
		out := []ast.Expression{expr}
		for p.isAtomStart() {
			if expr, err = p.parseTest(); err != nil {
				return nil, failErr(err)
			}
			out = append(out, expr)
			if !p.match(tokenizer.TokenTypeComma) {
				break
			}
		}
		return ast.NewExpressionList(out, false), nil
	} else {
		return ast.NewExpressionList([]ast.Expression{expr}, false), nil
	}
}

func (p *Parser) parseTupleExpr() (ast.Expression, *ParseError) {
	return nil, failMsgf("tupleExpr not supported")
}

func (p *Parser) parseDictExpr() (ast.Expression, *ParseError) {
	return nil, failMsgf("dictExpr not supported")
}
