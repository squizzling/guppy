package tokenizer

import (
	"errors"
	"fmt"

	"github.com/squizzling/types/pkg/result"
)

type TokenType int

var keywords = map[string]TokenType{
	"and":    TokenTypeAnd,
	"assert": TokenTypeAssert,
	"def":    TokenTypeFrom,
	"for":    TokenTypeFor,
	"from":   TokenTypeFrom,
	"if":     TokenTypeIf,
	"import": TokenTypeImport,
	"is":     TokenTypeIs,
	"lambda": TokenTypeLambda,
	"not":    TokenTypeNot,
	"or":     TokenTypeOr,
	"pass":   TokenTypePass,
	"return": TokenTypeReturn,

	"False": TokenTypeFalse,
	"None":  TokenTypeNone,
	"True":  TokenTypeTrue,
}

var singles = map[int]TokenType{
	'&': TokenTypeAmper,
	'@': TokenTypeAt,
	'^': TokenTypeCaret,
	':': TokenTypeColon,
	',': TokenTypeComma,
	'{': TokenTypeLeftBrace,
	'(': TokenTypeLeftParen,
	'[': TokenTypeLeftSquare,
	'-': TokenTypeMinus,
	'|': TokenTypePipe,
	'+': TokenTypePlus,
	'}': TokenTypeRightBrace,
	')': TokenTypeRightParen,
	']': TokenTypeRightSquare,
	';': TokenTypeSemiColon,
	'~': TokenTypeTilde,
}

var indent = map[int]bool{
	'(': true,
	'{': true,
	'[': true,
}

var dedent = map[int]bool{
	')': true,
	'}': true,
	']': true,
}

var doubles = map[int]TokenType{
	'!':          TokenTypeBang,
	'!'<<8 | '=': TokenTypeBangEqual,

	'=':          TokenTypeEqual,
	'='<<8 | '=': TokenTypeEqualEqual,

	'>':          TokenTypeGreater,
	'>'<<8 | '=': TokenTypeGreaterEqual,
	'>'<<8 | '>': TokenTypeGreaterGreater,

	'<':          TokenTypeLess,
	'<'<<8 | '=': TokenTypeLessEqual,
	'<'<<8 | '>': TokenTypeLessGreater,
	'<'<<8 | '<': TokenTypeLessLess,

	'/':          TokenTypeSlash,
	'/'<<8 | '/': TokenTypeSlashSlash,

	'*':          TokenTypeStar,
	'*'<<8 | '*': TokenTypeStarStar,
}

const (
	TokenTypeEOF = TokenType(iota + 1)

	TokenTypeDedent
	TokenTypeIndent
	TokenTypeNewLine

	TokenTypeAt
	TokenTypeAmper
	TokenTypeAnd
	TokenTypeBang
	TokenTypeBangEqual
	TokenTypeCaret
	TokenTypeColon
	TokenTypeComma
	TokenTypeDot
	TokenTypeEqual
	TokenTypeEqualEqual
	TokenTypeGreater
	TokenTypeGreaterEqual
	TokenTypeGreaterGreater
	TokenTypeLeftBrace
	TokenTypeLeftParen
	TokenTypeLeftSquare
	TokenTypeLess
	TokenTypeLessEqual
	TokenTypeLessGreater
	TokenTypeLessLess
	TokenTypeMinus
	TokenTypeNot
	TokenTypeOr
	TokenTypePipe
	TokenTypePlus
	TokenTypeRightBrace
	TokenTypeRightParen
	TokenTypeRightSquare
	TokenTypeSemiColon
	TokenTypeSlash
	TokenTypeSlashSlash
	TokenTypeStar
	TokenTypeStarStar
	TokenTypeTilde

	TokenTypeAssert
	TokenTypeDef
	TokenTypeFor
	TokenTypeFrom
	TokenTypeIf
	TokenTypeImport
	TokenTypeIs
	TokenTypeIsNot // This is not generated in the lexer, it's a synthetic token the parser uses
	TokenTypeLambda
	TokenTypePass
	TokenTypeReturn

	TokenTypeFalse
	TokenTypeNone
	TokenTypeTrue

	TokenTypeFloat
	TokenTypeIdentifier
	TokenTypeInt
	TokenTypeString
)

func (tt TokenType) String() string {
	switch tt {
	case TokenTypeEOF:
		return "EOF"

	case TokenTypeDedent:
		return "DEDENT"
	case TokenTypeIndent:
		return "INDENT"
	case TokenTypeNewLine:
		return "NEW_LINE"

	case TokenTypeAmper:
		return "AMPER"
	case TokenTypeAt:
		return "AT"
	case TokenTypeBang:
		return "BANG"
	case TokenTypeBangEqual:
		return "BANG_EQUAL"
	case TokenTypeCaret:
		return "CARET"
	case TokenTypeColon:
		return "COLON"
	case TokenTypeComma:
		return "COMMA"
	case TokenTypeDot:
		return "DOT"
	case TokenTypeEqual:
		return "EQUAL"
	case TokenTypeEqualEqual:
		return "EQUAL_EQUAL"
	case TokenTypeFor:
		return "FOR"
	case TokenTypeGreater:
		return "GREATER"
	case TokenTypeGreaterEqual:
		return "GREATER_EQUAL"
	case TokenTypeLeftBrace:
		return "LEFT_BRACE"
	case TokenTypeLeftParen:
		return "LEFT_PAREN"
	case TokenTypeLeftSquare:
		return "LEFT_SQUARE"
	case TokenTypeLess:
		return "LESS"
	case TokenTypeLessEqual:
		return "LESS_EQUAL"
	case TokenTypeLessGreater:
		return "LESS_GREATER"
	case TokenTypeMinus:
		return "MINUS"
	case TokenTypePipe:
		return "PIPE"
	case TokenTypePlus:
		return "PLUS"
	case TokenTypeRightBrace:
		return "RIGHT_BRACE"
	case TokenTypeRightParen:
		return "RIGHT_PAREN"
	case TokenTypeRightSquare:
		return "RIGHT_SQUARE"
	case TokenTypeSemiColon:
		return "SEMI_COLON"
	case TokenTypeSlash:
		return "SLASH"
	case TokenTypeSlashSlash:
		return "SLASH_SLASH"
	case TokenTypeStar:
		return "STAR"
	case TokenTypeStarStar:
		return "STAR_STAR"
	case TokenTypeTilde:
		return "TILDE"

	case TokenTypeAnd:
		return "AND"
	case TokenTypeAssert:
		return "ASSERT"
	case TokenTypeDef:
		return "DEF"
	case TokenTypeFrom:
		return "FROM"
	case TokenTypeIf:
		return "IF"
	case TokenTypeImport:
		return "IMPORT"
	case TokenTypeIs:
		return "IS"
	case TokenTypeIsNot:
		return "IS_NOT"
	case TokenTypeLambda:
		return "LAMBDA"
	case TokenTypeNot:
		return "NOT"
	case TokenTypeOr:
		return "OR"
	case TokenTypePass:
		return "PASS"
	case TokenTypeReturn:
		return "RETURN"

	case TokenTypeFalse:
		return "FALSE"
	case TokenTypeNone:
		return "NONE"
	case TokenTypeTrue:
		return "TRUE"

	case TokenTypeFloat:
		return "FLOAT"
	case TokenTypeIdentifier:
		return "IDENTIFIER"
	case TokenTypeInt:
		return "INT"
	case TokenTypeString:
		return "STRING"

	default:
		panic(fmt.Sprintf("TokenType(%d)", tt))
	}
}

type Token struct {
	Lexeme         string
	LiteralString  string
	LiteralInteger int

	Type TokenType
}

type Tokenizer struct {
	data   string
	offset int
	start  int

	startOfLine bool
	inGroup     int

	indents       []int
	dedentPending int

	tokensPending []result.Result[Token]
}

func (t *Tokenizer) newToken(tt TokenType) Token {
	return Token{
		Lexeme: t.data[t.start:t.offset],
		Type:   tt,
	}
}

func (t *Tokenizer) newTokenIdentifier() Token {
	lexeme := t.data[t.start:t.offset]
	tt, ok := keywords[lexeme]
	if !ok {
		tt = TokenTypeIdentifier
	}
	return Token{
		Lexeme: lexeme,
		Type:   tt,
	}
}

func (t *Tokenizer) newTokenInt(n int) Token {
	lexeme := t.data[t.start:t.offset]
	tt, ok := keywords[lexeme]
	if !ok {
		tt = TokenTypeInt
	}
	return Token{
		Lexeme:         lexeme,
		Type:           tt,
		LiteralInteger: n,
	}
}

func (t *Tokenizer) newTokenFloat() Token {
	lexeme := t.data[t.start:t.offset]
	tt, ok := keywords[lexeme]
	if !ok {
		tt = TokenTypeFloat
	}
	return Token{
		Lexeme: lexeme,
		Type:   tt,
	}
}

func (t *Tokenizer) newTokenIndent(indent int) Token {
	t.indents = append(t.indents, indent)
	return Token{
		Type: TokenTypeIndent,
	}
}

func (t *Tokenizer) calculateIndent(indent int) error {
	fmt.Printf("doing dedent to %d\n", indent)
	for len(t.indents) > 0 {
		amt := t.indents[len(t.indents)-1]
		if amt > indent {
			t.indents = t.indents[:len(t.indents)-1]
			t.dedentPending++
		} else if amt == indent {
			return nil
		} else if amt < indent {
			return errors.New("invalid indent level")
		}
	}

	// Nothing left to pop, we hit the end, which is a valid dedent.
	return nil
}

func (t *Tokenizer) newTokenString(literal string) Token {
	return Token{
		Lexeme:        t.data[t.start:t.offset],
		LiteralString: literal,
		Type:          TokenTypeString,
	}
}

func (t *Tokenizer) newEOF() Token {
	return Token{
		Type: TokenTypeEOF,
	}
}

func (t *Tokenizer) more() bool {
	return t.offset < len(t.data)
}

func (t *Tokenizer) next() int {
	ch := t.data[t.offset]
	t.offset++
	return int(ch)
}

func (t *Tokenizer) reset() {
	t.offset = t.start
}

func (t *Tokenizer) match(ch byte) bool {
	if !t.more() {
		return false
	}
	if t.data[t.offset] == ch {
		t.offset++
		return true
	}
	return false
}

func (t *Tokenizer) peek(n int) int {
	if t.offset+n >= len(t.data) {
		return -1
	}
	return int(t.data[t.offset+n])
}

func NewTokenizer(data string) *Tokenizer {
	return &Tokenizer{
		data:        data,
		startOfLine: true,
	}
}

func (t *Tokenizer) lastIndent() int {
	if len(t.indents) == 0 {
		return 0
	} else {
		return t.indents[len(t.indents)-1]
	}
}

func (t *Tokenizer) Peek(n int) result.Result[Token] {
	for len(t.tokensPending) <= n {
		if len(t.tokensPending) > 0 {
			lastToken := t.tokensPending[len(t.tokensPending)-1]
			if !lastToken.Ok() || lastToken.Value().Type == TokenTypeEOF {
				return lastToken
			}
		}
		t.tokensPending = append(t.tokensPending, t.getNext())
	}
	return t.tokensPending[n]
}

func (t *Tokenizer) Get() result.Result[Token] {
	t.Peek(0)
	firstToken := t.tokensPending[0]
	if !firstToken.Ok() || firstToken.Value().Type == TokenTypeEOF {
		return firstToken
	}
	t.tokensPending = t.tokensPending[1:]
	return firstToken
}

func (t *Tokenizer) getNext() result.Result[Token] {
	// If it's start of the line, measure our spaces
	if t.inGroup == 0 {
		if t.startOfLine {
			indent := 0

			for {
				if t.match(' ') {
					indent++
				} else if t.match('\t') {
					indent = ((indent / 8) + 1) * 8
				} else if t.match('\\') {
					panic("'\\' in indentation is not supported")
				} else if t.match('#') {
					// Ignore the line entirely
					for t.more() && t.next() != '\n' { // Scan to EOL
					}
					indent = 0 // Reset the indent
					continue   // Re-enter the SOL loop
				} else if t.match('\n') {
					indent = 0
					continue
				} else {
					break
				}
			}

			t.startOfLine = false
			if indent > t.lastIndent() {
				return result.Ok(t.newTokenIndent(indent))
			} else if indent < t.lastIndent() {
				if err := t.calculateIndent(indent); err != nil {
					return result.Err[Token](err)
				}

			}
		}
	} else {
		// We're in a group
		// Eat whitespace, including comments and shit
		for {
			if t.match(' ') {
			} else if t.match('\t') {
			} else if t.match('\\') {
				panic("'\\' in indentation is not supported")
			} else if t.match('#') {
				// Ignore the line entirely
				for t.more() && t.next() != '\n' { // Scan to EOL
				}
				continue // Re-enter the SOL loop
			} else if t.match('\n') {
				continue
			} else {
				break
			}
		}
		t.startOfLine = false
	}

	// If there was an indent, we already returned
	// If there was a dedent, we already returned
	// We must be at the start of the line, after indents, but there may be pending dedents, so handle them if necessary.
	if t.dedentPending > 0 {
		t.dedentPending--
		return result.Ok(Token{Type: TokenTypeDedent})
	}

	if !t.more() { // We're at EOF.  Deal with any dedents
		t.dedentPending = len(t.indents)
		if t.dedentPending == 0 {
			return result.Ok(t.newEOF())
		} else {
			return result.Ok(Token{Type: TokenTypeDedent})
		}
	}

	// Skip any whitespace
	for t.more() && (t.match(' ') || t.match('\t')) {
	}

	t.start = t.offset

	// We are at the start of a fresh token.
	ch := t.next()
	if tt, ok := singles[ch]; ok {
		if indent[ch] {
			t.inGroup++
		} else if dedent[ch] {
			t.inGroup--
		}
		return result.Ok(t.newToken(tt))
	} else if ttSingle, ok := doubles[ch]; ok {
		chNext := t.peek(0)
		chDouble := ch<<8 | chNext
		if ttDouble, ok := doubles[chDouble]; ok {
			return result.Ok(t.newToken(ttDouble))
		} else {
			return result.Ok(t.newToken(ttSingle))
		}
	}

	switch ch {
	case '\n': // Empty lines are already ignored and filtered out
		t.startOfLine = true
		return result.Ok(t.newToken(TokenTypeNewLine))
	case '#': // This is a comment after a token, so we want to emit a new line after processing it
		for t.more() && t.next() != '\n' {
		}
		t.startOfLine = true
		return result.Ok(t.newToken(TokenTypeNewLine))
	case '\\':
		panic("deal with line wrapping")
	case '"', '\'':
		if t.peek(0) == ch && t.peek(1) == ch {
			panic("deal with docstrings")
		}
		for t.more() {
			if t.next() == ch {
				return result.Ok(t.newTokenString(t.data[t.start+1 : t.offset-1]))
			}
		}

		return result.Err[Token](errors.New("unterminated string"))
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		// Easy path: check for '.' followed by a non-number
		if ch == '.' && !isNum(t.peek(0)) {
			return result.Ok(t.newToken(TokenTypeDot))
		}
		t.reset()

		// consume 0-9 until we hit something else
		val := 0
		for isNum(t.peek(0)) {
			n := t.next()
			val = val*10 + (n - '0')
		}
		// if it's not a dot, then we have an int
		return result.Ok(t.newTokenInt(val))

		/*
			INT
			  : [0-9]+
			  ;

			FLOAT
			  : [0-9]* '.'? [0-9]+ ([eE][+-]?[0-9]+)?
			  ;
		*/
		// TODO: This can probably be done more consistently with the spec
		panic("deal with floats")
	default:
		if isAlpha(ch) || ch == '_' {
			for {
				ch := t.peek(0)
				if isAlpha(ch) || isNum(ch) || ch == '_' {
					t.offset++
					continue
				}
				return result.Ok(t.newTokenIdentifier())
			}
		} else {
			start := max(0, t.start)
			end := min(t.offset+5, len(t.data))
			return result.Err[Token](fmt.Errorf("lex error: [%s]", t.data[start:end]))
		}
	}
}

func (t *Tokenizer) Debug(s string, n int) {
	fmt.Printf("%s [", s)
	for i := 0; i < n; i++ {
		tok := t.Peek(i)
		if tok.Ok() {
			fmt.Printf("%s(%s) ", tok.Value().Type, tok.Value().Lexeme)
		} else {
			fmt.Printf("(%s) ", tok.Err())
		}
	}
	fmt.Printf("]\n")
}

func isAlpha(ch int) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func isNum(ch int) bool {
	return ch >= '0' && ch <= '9'
}
