# Makefile for Calculator TDD Project

.PHONY: test coverage lint all clean help

# Default target
all: test lint

# Run test suite
test:
	@echo "Running test suite..."
	go test -v ./test/...

# Generate HTML coverage report
coverage:
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run golangci-lint
lint:
	@echo "Running golangci-lint..."
	golangci-lint run

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f coverage.out coverage.html

# Show help
help:
	@echo "Available targets:"
	@echo "  test      - Run the test suite"
	@echo "  coverage  - Generate HTML coverage report"
	@echo "  lint      - Run golangci-lint static analysis"
	@echo "  all       - Run all checks (test + lint)"
	@echo "  clean     - Remove generated files"
	@echo "  help      - Show this help message"
