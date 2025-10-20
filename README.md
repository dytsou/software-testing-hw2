# Calculator TDD Project

A simple calculator implementation in Go developed using Test-Driven Development (TDD) methodology. This project demonstrates the complete TDD cycle with comprehensive test coverage and CI/CD pipeline setup.

## Features

- **Addition**: Add two integers
- **Subtraction**: Subtract two integers  
- **Multiplication**: Multiply two integers
- **Division**: Divide two integers with proper error handling for division by zero

## Project Structure

```
calc/
├── internal/
│   └── calc/
│       └── calc.go          # Calculator implementation
├── test/
│   └── calc_test.go         # Test suite
├── .github/
│   └── workflows/
│       └── ci.yml           # CI/CD pipeline
├── Makefile                 # Build automation
├── go.mod                   # Go module definition
├── NARRATIVE.md             # Detailed TDD process documentation
└── README.md                # This file
```

## Requirements

- Go 1.21 or later
- Git
- Make (optional, for using Makefile targets)

## Usage

### Import and Use

```go
import "github.com/dytsou/calc/internal/calc"

// Addition
result := calc.Add(2, 3)  // Returns 5

// Subtraction  
result := calc.Subtract(5, 3)  // Returns 2

// Multiplication
result := calc.Multiply(2, 3)  // Returns 6

// Division (returns float64 and error)
result, err := calc.Divide(6, 3)  // Returns 2.0, nil
result, err := calc.Divide(5, 0)  // Returns 0.0, error
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with verbose output
go test -v ./test/...

# Run tests with coverage
make coverage

# Run all checks (test + lint)
make all
```

### Available Make Targets

- `make test` - Run the test suite
- `make coverage` - Generate HTML coverage report
- `make lint` - Run golangci-lint static analysis
- `make all` - Run all checks (test + lint)

## Development Process

This project was developed using Test-Driven Development (TDD) with the following approach:

1. **Write Failing Test** - Create test cases for new functionality
2. **Implement Function** - Write minimal code to make tests pass
3. **Refactor** - Improve code quality while keeping tests green
4. **Commit** - Each TDD cycle committed separately

### TDD Cycles Completed

- **Cycle 1**: Add function with comprehensive test cases
- **Cycle 2**: Subtract function with edge case testing
- **Cycle 3**: Multiply function with large number testing
- **Cycle 4**: Divide function with error handling for division by zero

See [NARRATIVE.md](NARRATIVE.md) for detailed documentation of each TDD cycle.

## Continuous Integration

This project includes a GitHub Actions CI/CD pipeline that:

- **Tests**: Runs test suite on multiple Go versions (1.21, 1.22)
- **Coverage**: Generates and displays test coverage reports
- **Linting**: Performs static analysis using golangci-lint

### CI/CD Features

- Automated testing on push and pull requests
- Multi-version Go testing (1.21, 1.22)
- Code coverage reporting
- Static analysis with golangci-lint
- Build status badges

## Test Coverage

The test suite includes:

- **Normal Cases**: Standard arithmetic operations with positive and negative numbers
- **Edge Cases**: Operations with zero values, large numbers
- **Error Cases**: Division by zero error handling
- **Table-Driven Tests**: Comprehensive test coverage using Go's idiomatic testing pattern

## Error Handling

The Divide function implements proper error handling:

```go
result, err := calc.Divide(a, b)
if err != nil {
    // Handle division by zero error
    log.Printf("Error: %v", err)
    return
}
// Use result safely
```

## Git History

The project maintains a clean git history showing the TDD progression:

- Each TDD cycle is committed separately
- Clear commit messages describing the TDD step
- Easy to follow the red-green-refactor cycle

## Contributing

1. Fork the repository
2. Create a feature branch
3. Follow TDD methodology for new features
4. Ensure all tests pass
5. Submit a pull request

## License

This project is part of a software testing course assignment and is for educational purposes.
