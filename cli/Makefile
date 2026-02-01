.PHONY: build run test lint clean install

BINARY_NAME ?= cli
CMD_PATH    := ./cmd/cli

build:
	go build -o bin/$(BINARY_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)/main.go

test:
	go test ./...

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/ coverage.out coverage.html

# Install CLI to $GOBIN (default $GOPATH/bin)
install:
	go install $(CMD_PATH)

# Example: run version subcommand
version:
	go run $(CMD_PATH)/main.go version

# Example: run sample command with flag
hello:
	go run $(CMD_PATH)/main.go run --name "Cosmos"
