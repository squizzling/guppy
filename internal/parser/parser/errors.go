package parser

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type ParseError struct {
	Location string
	Message  string
	err      error
}

func (p *ParseError) Error() string {
	if p == nil {
		panic("nil p")
	}
	if p.err == nil {
		panic("nil p.err")
	}
	return p.err.Error()
}

func (p *ParseError) Stack() []*ParseError {
	if pse, ok := p.err.(*ParseError); ok && pse != nil {
		return append(
			[]*ParseError{p},
			pse.Stack()...,
		)
	}
	return []*ParseError{p}
}

func caller(skip int) (string, string, int) {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+3, rpc)
	if n < 1 {
		return "unknown", "unknown", -1
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	return frame.Function, frame.File, frame.Line
}

func callerName(skip int) string {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+3, rpc)
	if n < 1 {
		return "unknown"
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	parts := strings.Split(frame.Function, ".")
	return parts[len(parts)-1]
}

func FailErrSkip(err error, msg string, skip int) *ParseError {
	fn, fl, l := caller(skip)
	return &ParseError{
		Location: fmt.Sprintf("%s:%d [%s]", fl, l, fn),
		Message:  msg,
		err:      err,
	}
}

func FailErr(err *ParseError) *ParseError {
	return FailErrSkip(err, "", 1)
}

func FailMsgf(f string, args ...any) *ParseError {
	msg := fmt.Sprintf(f, args...)
	return FailErrSkip(errors.New(msg), msg, 1)
}

func Wrap[T any](t T, err *ParseError) (T, *ParseError) {
	if err != nil {
		var zero T
		return zero, FailErrSkip(err, "", 1)
	} else {
		return t, nil
	}
}
