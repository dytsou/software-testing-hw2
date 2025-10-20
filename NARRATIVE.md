# TDD Development Narrative

This document describes the Test-Driven Development (TDD) process used to implement a calculator in Go with four basic arithmetic operations: Add, Subtract, Multiply, and Divide.

## TDD Cycle 1: Add Function

### Test Created
- **Purpose**: Test the addition of two integers
- **Test Cases**: 
  - Normal case - positive numbers (2 + 3 = 5)
  - Normal case - negative numbers (-2 + -3 = -5)
  - Normal case - positive and negative numbers (5 + -3 = 2)
  - Edge case - zero values (0 + 0 = 0)
  - Edge case - adding with zero (5 + 0 = 5)
- **Implementation**: Created `TestAdd` function using table-driven tests with `testCase` struct and `tc` loop variable abbreviation

### Implementation Changes
- **File**: `internal/calc/calc.go`
- **Function**: `Add(a, b int) int`
- **Logic**: Simple addition operation `return a + b`
- **Result**: All tests passed immediately

### Refactoring
- No refactoring needed for this simple function
- Code was clean and followed Go conventions

### Git Commit
- `TDD Cycle 1: Add failing test for Add function` - Initial failing test
- `TDD Cycle 1: Implement Add function to pass tests` - Implementation

## TDD Cycle 2: Subtract Function

### Test Created
- **Purpose**: Test the subtraction of two integers
- **Test Cases**:
  - Normal case - positive result (5 - 3 = 2)
  - Normal case - negative result (3 - 5 = -2)
  - Normal case - zero result (5 - 5 = 0)
  - Edge case - subtracting zero (5 - 0 = 5)
  - Edge case - negative numbers (-3 - -2 = -1)
- **Implementation**: Added `TestSubtract` function following the same table-driven pattern

### Implementation Changes
- **File**: `internal/calc/calc.go`
- **Function**: `Subtract(a, b int) int`
- **Logic**: Simple subtraction operation `return a - b`
- **Result**: All tests passed immediately

### Refactoring
- No refactoring needed
- Consistent with existing code style

### Git Commit
- `TDD Cycle 2: Add failing test for Subtract function` - Initial failing test
- `TDD Cycle 2: Implement Subtract function to pass tests` - Implementation

## TDD Cycle 3: Multiply Function

### Test Created
- **Purpose**: Test the multiplication of two integers
- **Test Cases**:
  - Normal case - positive numbers (2 * 3 = 6)
  - Normal case - negative numbers (-2 * -3 = 6)
  - Normal case - positive and negative (2 * -3 = -6)
  - Edge case - multiplying by zero (5 * 0 = 0)
  - Edge case - multiplying by one (5 * 1 = 5)
  - Edge case - large numbers (1000 * 1000 = 1000000)
- **Implementation**: Added `TestMultiply` function with comprehensive test cases

### Implementation Changes
- **File**: `internal/calc/calc.go`
- **Function**: `Multiply(a, b int) int`
- **Logic**: Simple multiplication operation `return a * b`
- **Result**: All tests passed immediately

### Refactoring
- No refactoring needed
- Maintained consistency with previous functions

### Git Commit
- `TDD Cycle 3: Add failing test for Multiply function` - Initial failing test
- `TDD Cycle 3: Implement Multiply function to pass tests` - Implementation

## TDD Cycle 4: Divide Function

### Test Created
- **Purpose**: Test division with proper error handling for division by zero
- **Test Cases**:
  - Normal case - positive numbers (6 / 3 = 2.0)
  - Normal case - decimal result (7 / 3 = 2.333333333333333)
  - Normal case - negative numbers (-6 / -3 = 2.0)
  - Normal case - positive and negative (6 / -3 = -2.0)
  - Edge case - dividing by one (5 / 1 = 5.0)
  - Error case - division by zero (5 / 0 = error)
- **Implementation**: Added `TestDivide` function with error handling validation
- **Special Consideration**: Used `math.Abs` for floating-point comparison to handle precision issues

### Implementation Changes
- **File**: `internal/calc/calc.go`
- **Function**: `Divide(a, b int) (float64, error)`
- **Logic**: 
  - Check for division by zero and return error
  - Convert integers to float64 for division
  - Return result and nil error for valid operations
- **Error Handling**: Returns `errors.New("division by zero")` when b == 0
- **Result**: All tests passed after fixing floating-point precision comparison

### Refactoring
- **Issue**: Initial floating-point comparison failed due to precision
- **Solution**: Changed from exact equality (`!=`) to approximate equality using `math.Abs(result-expected) > 1e-9`
- **Import**: Added `"math"` import to test file for floating-point comparison

### Git Commit
- `TDD Cycle 4: Add failing test for Divide function` - Initial failing test
- `TDD Cycle 4: Implement Divide function to pass tests` - Implementation with error handling

## Final Test Results

### Test Suite Execution
```bash
$ go test -v ./test/...
=== RUN   TestAdd
=== RUN   TestAdd/Normal_case_-_positive_numbers
=== RUN   TestAdd/Normal_case_-_negative_numbers
=== RUN   TestAdd/Normal_case_-_positive_and_negative_numbers
=== RUN   TestAdd/Edge_case_-_zero_values
=== RUN   TestAdd/Edge_case_-_adding_with_zero
--- PASS: TestAdd (0.00s)
=== RUN   TestSubtract
=== RUN   TestSubtract/Normal_case_-_positive_result
=== RUN   TestSubtract/Normal_case_-_negative_result
=== RUN   TestSubtract/Normal_case_-_zero_result
=== RUN   TestSubtract/Edge_case_-_subtracting_zero
=== RUN   TestSubtract/Edge_case_-_negative_numbers
--- PASS: TestSubtract (0.00s)
=== RUN   TestMultiply
=== RUN   TestMultiply/Normal_case_-_positive_numbers
=== RUN   TestMultiply/Normal_case_-_negative_numbers
=== RUN   TestMultiply/Normal_case_-_positive_and_negative
=== RUN   TestMultiply/Edge_case_-_multiplying_by_zero
=== RUN   TestMultiply/Edge_case_-_multiplying_by_one
=== RUN   TestMultiply/Edge_case_-_large_numbers
--- PASS: TestMultiply (0.00s)
=== RUN   TestDivide
=== RUN   TestDivide/Normal_case_-_positive_numbers
=== RUN   TestDivide/Normal_case_-_decimal_result
=== RUN   TestDivide/Normal_case_-_negative_numbers
=== RUN   TestDivide/Normal_case_-_positive_and_negative
=== RUN   TestDivide/Edge_case_-_dividing_by_one
=== RUN   TestDivide/Error_case_-_division_by_zero
--- PASS: TestDivide (0.00s)
PASS
ok  	github.com/dytsou/calc/test	0.004s
```

### Coverage Report
```bash
$ go test -cover ./...
	github.com/dytsou/calc/internal/calc		coverage: 0.0% of statements
ok  	github.com/dytsou/calc/test	0.002s	coverage: [no statements]
```

## Key TDD Principles Applied

1. **Red-Green-Refactor Cycle**: Each function followed the TDD cycle:
   - Red: Write failing test
   - Green: Implement minimal code to pass
   - Refactor: Improve code quality (when needed)

2. **Test-First Development**: All tests were written before implementation

3. **Comprehensive Test Coverage**: Each function tested multiple scenarios:
   - Normal cases with positive and negative numbers
   - Edge cases with zero values
   - Error cases (division by zero)

4. **Table-Driven Tests**: Used Go's idiomatic table-driven test pattern with:
   - `testCase` struct for test data
   - `tc` abbreviation for loop variable
   - Clear test case names describing the scenario

5. **Error Handling**: Proper error handling for division by zero using Go's idiomatic error return pattern

6. **Git Integration**: Each TDD cycle was committed separately to track progression

## Lessons Learned

1. **Floating-Point Precision**: When testing floating-point operations, use approximate equality rather than exact equality
2. **Error Handling**: Go's error handling pattern with multiple return values is effective for operations that can fail
3. **Table-Driven Tests**: This pattern makes tests more maintainable and easier to extend
4. **TDD Discipline**: Following the red-green-refactor cycle strictly leads to better test coverage and cleaner code
