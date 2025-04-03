.PHONY: gen
gen:
	go run ./cmd/gen-flow-ast     >internal/parser/ast/gen_flow_ast.go
	go run ./cmd/gen-flow-debug   >internal/parser/ast/gen_flow_debug.go
	go run ./cmd/gen-stream-ast   >internal/flow/stream/gen_stream_ast.go
	go run ./cmd/gen-stream-debug >internal/flow/stream/gen_stream_debug.go


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
