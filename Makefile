.PHONY: gen
gen:
	go run ./cmd/gen-ast   >internal/parser/ast/gen_ast.go
	go run ./cmd/gen-debug >internal/parser/ast/gen_debug.go

.PHONY: imports
imports:
	goimports -d -local guppy .

.PHONY: test
test:
	go test -coverprofile profile.out ./...
	go tool cover -html profile.out -o profile.html

.PHONY: rebuild-tests
rebuild-tests:
	go test -run TestRebuild -tags rebuild ./internal/parser/flow -v
