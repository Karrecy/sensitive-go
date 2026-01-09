.PHONY: all build test test-coverage bench lint clean help

# Build variables
BINARY_NAME=gosensitive
GO=go
GOTEST=$(GO) test
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean

all: test build

## build: Build the project
build:
	$(GOBUILD) -v ./...

## test: Run all tests
test:
	$(GOTEST) -v ./...

## test-coverage: Run tests with coverage
test-coverage:
	$(GOTEST) -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	$(GO) tool cover -html=coverage.txt -o coverage.html

## bench: Run benchmark tests
bench:
	$(GOTEST) -bench=. -benchmem -run=^$$ ./benchmarks

## bench-compare: Run comparison benchmarks
bench-compare:
	$(GOTEST) -bench=Compare -benchmem -run=^$$ ./benchmarks

## lint: Run linter (requires golangci-lint)
lint:
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed" && exit 1)
	golangci-lint run ./...

## fmt: Format code
fmt:
	$(GO) fmt ./...

## vet: Run go vet
vet:
	$(GO) vet ./...

## mod-tidy: Tidy go modules
mod-tidy:
	$(GO) mod tidy

## clean: Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f coverage.txt coverage.html

## examples: Run example code
examples:
	$(GO) run examples/basic/main.go
	$(GO) run examples/advanced/main.go

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'


