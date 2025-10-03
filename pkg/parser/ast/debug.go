package ast

func DebugExpression(e Expression) string {
	a, _ := e.Accept(DebugWriter{})
	return a.(string)
}

func DebugStatement(s Statement) string {
	a, _ := s.Accept(DebugWriter{})
	return a.(string)
}
