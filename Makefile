.PHONY: gen
gen:
	go run ./cmd/gen-ast   >internal/parser/ast/gen_ast.go
	go run ./cmd/gen-debug >internal/parser/ast/gen_debug.go

.PHONY: imports
imports:
	goimports -d -local guppy .
