.PHONY: run build test lint

run:
	go run ./cmd/main/...

build:
	go build -o ./bin/service ./cmd/main/...

test:
	go test ./... -v -cover

lint:
	golangci-lint run ./...
