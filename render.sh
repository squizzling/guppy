dot -Tsvg -o $(basename $1 .flow).svg <(go run ./cmd/graph $1)
