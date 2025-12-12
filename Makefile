.PHONY: gen
gen:
	go run ./cmd/gen-flow-ast     >pkg/parser/ast/gen_flow_ast.go
	go run ./cmd/gen-flow-debug   >pkg/parser/ast/gen_flow_debug.go
	go run ./cmd/gen-stream-ast   >pkg/flow/stream/gen_stream_ast.go
	go run ./cmd/gen-stream-debug >pkg/flow/stream/gen_stream_debug.go
	go run ./cmd/gen-filter-ast   >pkg/flow/filter/gen_filter_ast.go


.PHONY: imports
imports:
	goimports -d -local guppy .

.PHONY: fmt
fmt:
	gofmt -l $$(find -name '*.go')

.PHONY: test
test:
	go test -coverprofile profile.out ./...
	go tool cover -html profile.out -o profile.html

.PHONY: rebuild-tests
rebuild-tests:
	go test -run TestRebuild -tags rebuild ./pkg/parser/flow -v
