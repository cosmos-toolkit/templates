.PHONY: build run test lint clean generate

BINARY_NAME ?= server
CMD_PATH    := ./cmd/server

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

# Generate Go code from proto (requires buf: https://buf.build/docs/installation)
generate:
	buf generate

# Docker (optional)
docker-build:
	docker build -t $(BINARY_NAME):latest -f build/Dockerfile .

docker-run:
	docker run --rm -p 9090:9090 $(BINARY_NAME):latest
