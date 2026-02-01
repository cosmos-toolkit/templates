.PHONY: test lint clean example

test:
	go test ./...

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...

clean:
	rm -rf coverage.out coverage.html

# Run the basic example (requires network if API is real).
example:
	go run ./examples/basic/main.go
