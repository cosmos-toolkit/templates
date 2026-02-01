.PHONY: build run test lint clean zip

BINARY_NAME ?= bootstrap
CMD_PATH    := ./cmd/handler

# Lambda runs the binary as "bootstrap" in the runtime environment.
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(BINARY_NAME) $(CMD_PATH)

# Local run (requires AWS Lambda Go support or a small test harness).
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
	rm -rf bin/ coverage.out coverage.html *.zip

# Package for Lambda deployment (binary + zip). Upload bin/bootstrap or the zip to Lambda.
zip: build
	cd bin && zip -q ../lambda.zip bootstrap
