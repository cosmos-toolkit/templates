.PHONY: build run test lint clean api worker cli up down

BINARY_DIR := bin
CMD_API    := ./cmd/api
CMD_WORKER := ./cmd/worker
CMD_CLI    := ./cmd/cli

# Build all binaries
build: build-api build-worker build-cli

build-api:
	go build -o $(BINARY_DIR)/api $(CMD_API)

build-worker:
	go build -o $(BINARY_DIR)/worker $(CMD_WORKER)

build-cli:
	go build -o $(BINARY_DIR)/cli $(CMD_CLI)

# Run locally (separate terminals or background)
run-api:
	go run $(CMD_API)/main.go

run-worker:
	go run $(CMD_WORKER)/main.go

run-cli:
	go run $(CMD_CLI)/main.go

# Docker Compose
up:
	docker compose up -d --build

down:
	docker compose down

logs:
	docker compose logs -f

# Tests and lint
test:
	go test ./...

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...

clean:
	rm -rf $(BINARY_DIR) coverage.out coverage.html

# Examples
version:
	go run $(CMD_CLI)/main.go version

hello:
	go run $(CMD_CLI)/main.go run --name "Cosmos"
