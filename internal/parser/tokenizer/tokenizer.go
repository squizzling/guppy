package tokenizer

import (
	"errors"
	"fmt"
	"strconv"
)

type TokenType int

var keywords = map[string]TokenType{
	"and":    TokenTypeAnd,
	"assert": TokenTypeAssert,
	"def":    TokenTypeDef,
	"else":   TokenTypeElse,
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
	TokenTypeError

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
	TokenTypeElse
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
	case TokenTypeError:
		return "ERROR"

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
	case TokenTypeElse:
		return "ELSE"
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
	LiteralFloat   float64
	Err            error

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

	tokensPending []Token
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

func (t *Tokenizer) newTokenError(err error) Token {
	return Token{
		Type:   TokenTypeError,
		Lexeme: err.Error(),
		Err:    err,
	}
}

func (t *Tokenizer) newTokenInt(n int) Token {
	return Token{
		Lexeme:         t.data[t.start:t.offset],
		Type:           TokenTypeInt,
		LiteralInteger: n,
	}
}

func (t *Tokenizer) newTokenNumber(preDot string, postDot string, exponentNegative bool, exponent string) Token {
	if postDot == "" && exponent == "" {
		if n, err := strconv.ParseInt(preDot, 10, 64); err != nil {
			return t.newTokenError(err)
		} else {
			return t.newTokenInt(int(n))
		}
	}

	lexeme := t.data[t.start:t.offset]

	parseInput := lexeme
	if false {
		// For fuzz testing and to see if it behaviors are close
		// enough to pass the raw lexeme to parseFloat
		parseInput = preDot
		if postDot != "" {
			parseInput += "." + postDot
		}
		if exponent != "" {
			parseInput += "e"
			if exponentNegative {
				parseInput += "-"
			}
			parseInput += exponent
		}
	}

	if f, err := strconv.ParseFloat(parseInput, 64); err != nil {
		return t.newTokenError(err)
	} else {
		return Token{
			Lexeme:       lexeme,
			Type:         TokenTypeFloat,
			LiteralFloat: f,
		}
	}
}

func (t *Tokenizer) newTokenIndent(indent int) Token {
	t.indents = append(t.indents, indent)
	return Token{
		Type: TokenTypeIndent,
	}
}

func (t *Tokenizer) calculateIndent(indent int) error {
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

func (t *Tokenizer) readDigits() string {
	start := t.offset
	for t.more() {
		if isNum(t.peek(0)) {
			t.offset++
		} else {
			break
		}
	}
	return t.data[start:t.offset]
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

func (t *Tokenizer) RemainingTokens() int {
	// Should only be used in tests, it peeks to the end of the buffer and returns the pending size.
	_ = t.Peek(len(t.data))
	return len(t.tokensPending)
}

func (t *Tokenizer) Peek(n int) Token {
	for len(t.tokensPending) <= n { // We need to read more.
		// If we have anything in the buffer, and the last thing we read was an EOF or error, return it
		if len(t.tokensPending) > 0 {
			if lastToken := t.tokensPending[len(t.tokensPending)-1]; lastToken.Type == TokenTypeEOF || lastToken.Type == TokenTypeError {
				return lastToken
			}
		}
		// Read the next thing, if it's an EOF or error, we'll catch it next iteration
		tok := t.getNext()
		if tok.Type == TokenTypeString {
			// Eat all the strings we can find
			for {
				tokNext := t.getNext()
				if tokNext.Type == TokenTypeString {
					tok.LiteralString += tokNext.LiteralString
					tok.Lexeme = tok.Lexeme + " " + tokNext.Lexeme
				} else {
					t.tokensPending = append(t.tokensPending, tok)     // Add the string
					t.tokensPending = append(t.tokensPending, tokNext) // Add the next token
					return t.tokensPending[0]                          // Return the string
				}
			}
		}
		t.tokensPending = append(t.tokensPending, tok)
	}
	return t.tokensPending[n]
}

func (t *Tokenizer) Get() Token {
	firstToken := t.Peek(0)
	if firstToken.Type != TokenTypeEOF && firstToken.Type != TokenTypeError {
		t.tokensPending = t.tokensPending[1:]
	}
	return firstToken
}

func (t *Tokenizer) Advance() {
	firstToken := t.Peek(0)
	if firstToken.Type != TokenTypeEOF && firstToken.Type != TokenTypeError {
		t.tokensPending = t.tokensPending[1:]
	}
}

func (t *Tokenizer) getNext() Token {
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
					return t.newTokenError(errors.New("'\\' in indentation is not supported"))
				} else if t.match('#') {
					// Ignore the line entirely
					for t.more() && t.next() != '\n' { // Scan to EOL
					}
					indent = 0 // Reset the indent
					continue   // Re-enter the SOL loop
				} else if t.peek(0) == '\r' && t.peek(1) == '\n' {
					t.offset += 2
					indent = 0
					continue
				} else if t.match('\n') {
					indent = 0
					continue
				} else {
					break
				}
			}

			t.startOfLine = false
			if indent > t.lastIndent() {
				return t.newTokenIndent(indent)
			} else if indent < t.lastIndent() {
				if err := t.calculateIndent(indent); err != nil {
					return t.newTokenError(err)
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
				return t.newTokenError(errors.New("'\\' in indentation is not supported"))
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
		return Token{Type: TokenTypeDedent}
	}

	if !t.more() { // We're at EOF.  Deal with any dedents
		if len(t.indents) > 0 {
			t.indents = t.indents[:len(t.indents)-1]
			return Token{Type: TokenTypeDedent}
		}
		return t.newEOF()
	}

	// Skip any whitespace
	for t.more() && (t.match(' ') || t.match('\t')) {
	}

	t.start = t.offset

	// We are at the start of a fresh token.

	if !t.more() {
		return t.newTokenError(errors.New("EOF"))
	}

	ch := t.next()
	if tt, ok := singles[ch]; ok {
		if indent[ch] {
			t.inGroup++
		} else if dedent[ch] {
			t.inGroup--
		}
		return t.newToken(tt)
	} else if ttSingle, ok := doubles[ch]; ok {
		chNext := t.peek(0)
		chDouble := ch<<8 | chNext
		if ttDouble, ok := doubles[chDouble]; ok {
			t.offset++
			return t.newToken(ttDouble)
		} else {
			return t.newToken(ttSingle)
		}
	}

	for ch == '\\' {
		if t.peek(0) == '\n' {
			t.offset++
		} else if t.peek(0) == '\r' && t.peek(1) == '\n' {
			t.offset += 2
		} else {
			return t.newTokenError(errors.New("'\\' without new line"))
		}

		// Eat any whitespace
		for t.more() && (t.match(' ') || t.match('\t')) {
		}

		t.start = t.offset
		ch = t.next()
	}

	switch ch {
	case '\r':
		if !t.match('\n') {
			return t.newTokenError(errors.New("CR without LF"))
		}
		fallthrough
	case '\n': // Empty lines are already ignored and filtered out
		t.startOfLine = true
		return t.newToken(TokenTypeNewLine)
	case '#': // This is a comment after a token, so we want to emit a new line after processing it
		for t.more() && t.next() != '\n' {
		}
		t.startOfLine = true
		return t.newToken(TokenTypeNewLine)
	case '"', '\'':
		if t.peek(0) == ch && t.peek(1) == ch {
			return t.newTokenError(errors.New("deal with line docstrings"))
		}
		for t.more() {
			if t.next() == ch {
				return t.newTokenString(t.data[t.start+1 : t.offset-1])
			}
		}

		return t.newTokenError(errors.New("unterminated string"))
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		// Easy path: check for '.' followed by a non-number
		if ch == '.' && !isNum(t.peek(0)) {
			return t.newToken(TokenTypeDot)
		}
		t.reset()

		preDot := ""
		postDot := ""
		if ch != '.' { // If we didn't start with a dot
			// Read digits (preDot)
			preDot = t.readDigits()
			ch = t.peek(0)
		}

		if ch == '.' { // If we hit a dot
			// Read digits (postDot)
			t.offset++
			postDot = t.readDigits()
			if postDot == "" {
				// The lexer doesn't allow "1.", so this is parsed as INT DOT.
				// We need to unread the dot.
				t.offset--
				return t.newTokenNumber(preDot, "", false, "")
			}
			ch = t.peek(0)
		}

		// If we didn't hit an e or E
		if ch != 'e' && ch != 'E' {
			// Done
			return t.newTokenNumber(preDot, postDot, false, "")
		}

		// Remember where we are
		curOffset := t.offset
		t.offset++ // skip the e

		// If we hit a + or -, record it
		ch = t.peek(0)
		exponentNegative := false
		if ch == '-' || ch == '+' {
			if ch == '-' {
				exponentNegative = true
			}
			t.offset++
			ch = t.peek(0)
		}

		if !isNum(ch) {
			// If we hit a non-digit, rollback
			t.offset = curOffset
			return t.newTokenNumber(preDot, postDot, false, "")
		}

		exponent := t.readDigits()
		if exponent == "" { // rewind
			t.offset = curOffset
			return t.newTokenNumber(preDot, postDot, false, "")
		}
		return t.newTokenNumber(preDot, postDot, exponentNegative, exponent)
	default:
		if isAlpha(ch) || ch == '_' {
			for {
				ch := t.peek(0)
				if isAlpha(ch) || isNum(ch) || ch == '_' {
					t.offset++
					continue
				}
				return t.newTokenIdentifier()
			}
		} else {
			start := max(0, t.start)
			end := min(t.offset+5, len(t.data))
			return t.newTokenError(fmt.Errorf("lex error: [%s]", t.data[start:end]))
		}
	}
}

func (t *Tokenizer) Debug(s string, n int) {
	fmt.Printf("%s [", s)
	defer fmt.Printf("]\n")
	for i := 0; i < n; i++ {
		tok := t.Peek(i)
		if tok.Type == TokenTypeError {
			fmt.Printf("%s(%s) ", tok.Type, tok.Err)
		} else {
			fmt.Printf("%s(%s) ", tok.Type, tok.Lexeme)
		}

		if tok.Type == TokenTypeEOF || tok.Type == TokenTypeError {
			return
		}
	}
}

func isAlpha(ch int) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func isNum(ch int) bool {
	return ch >= '0' && ch <= '9'
}
