package tokenizer

import (
	"fmt"
	"strconv"
	"testing"
)

func FuzzFloat(f *testing.F) {
	// Technically this is fuzzing the entire lexer, but we're only looking at int and float outputs
	f.Add("12")
	f.Add("12e10")
	f.Add("12e-10")
	f.Add("12e+10")
	f.Add("12E10")
	f.Add("12E-10")
	f.Add("12E+10")
	f.Add("12.456")
	f.Add("12.456e10")
	f.Add("12.456e-10")
	f.Add("12.456e+10")
	f.Add("12.456E10")
	f.Add("12.456E-10")
	f.Add("12.456E+10")
	f.Add(".456")
	f.Add(".456e10")
	f.Add(".456e-10")
	f.Add(".456e+10")
	f.Add(".456E10")
	f.Add(".456E-10")
	f.Add(".456E+10")
	f.Fuzz(func(t *testing.T, s string) {
		tok := NewTokenizer(s).getNext()
		if tok.Type == TokenTypeFloat {
			parsedFloat, err := strconv.ParseFloat(tok.Lexeme, 64)
			if err != nil {
				t.Fatal(err.Error())
			}
			if parsedFloat != tok.LiteralFloat {
				t.Fatal(fmt.Sprintf("input: [%s], go parsed float: [%v], token float: [%v]", s, parsedFloat, tok.LiteralFloat))
			}
		} else if tok.Type == TokenTypeInt {
			parsedInt, err := strconv.ParseInt(tok.Lexeme, 10, 64)
			if err != nil {
				t.Fatal(err.Error())
			}
			if int(parsedInt) != tok.LiteralInteger {
				t.Fatal(fmt.Sprintf("input: [%s], go parsed int: [%v], token int: [%v]", s, parsedInt, tok.LiteralInteger))
			}
		}
	})
}
