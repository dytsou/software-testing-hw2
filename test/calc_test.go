package calc_test

import (
	"math"
	"testing"

	"github.com/dytsou/calc/internal/calc"
)

func TestAdd(t *testing.T) {
	// Define test case structure
	type Params struct {
		a, b int
	}
	type testCase struct {
		name        string
		params      Params
		expected    int
		expectedErr bool
	}

	// Define test case table
	testCases := []testCase{
		{
			name:        "Normal case - positive numbers",
			params:      Params{a: 2, b: 3},
			expected:    5,
			expectedErr: false,
		},
		{
			name:        "Normal case - negative numbers",
			params:      Params{a: -2, b: -3},
			expected:    -5,
			expectedErr: false,
		},
		{
			name:        "Normal case - positive and negative numbers",
			params:      Params{a: 5, b: -3},
			expected:    2,
			expectedErr: false,
		},
		{
			name:        "Edge case - zero values",
			params:      Params{a: 0, b: 0},
			expected:    0,
			expectedErr: false,
		},
		{
			name:        "Edge case - adding with zero",
			params:      Params{a: 5, b: 0},
			expected:    5,
			expectedErr: false,
		},
	}

	// Execute tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Execute the function under test
			result := calc.Add(tc.params.a, tc.params.b)

			// Verify result
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	// Define test case structure
	type Params struct {
		a, b int
	}
	type testCase struct {
		name        string
		params      Params
		expected    int
		expectedErr bool
	}

	// Define test case table
	testCases := []testCase{
		{
			name:        "Normal case - positive result",
			params:      Params{a: 5, b: 3},
			expected:    2,
			expectedErr: false,
		},
		{
			name:        "Normal case - negative result",
			params:      Params{a: 3, b: 5},
			expected:    -2,
			expectedErr: false,
		},
		{
			name:        "Normal case - zero result",
			params:      Params{a: 5, b: 5},
			expected:    0,
			expectedErr: false,
		},
		{
			name:        "Edge case - subtracting zero",
			params:      Params{a: 5, b: 0},
			expected:    5,
			expectedErr: false,
		},
		{
			name:        "Edge case - negative numbers",
			params:      Params{a: -3, b: -2},
			expected:    -1,
			expectedErr: false,
		},
	}

	// Execute tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Execute the function under test
			result := calc.Subtract(tc.params.a, tc.params.b)

			// Verify result
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	// Define test case structure
	type Params struct {
		a, b int
	}
	type testCase struct {
		name        string
		params      Params
		expected    int
		expectedErr bool
	}

	// Define test case table
	testCases := []testCase{
		{
			name:        "Normal case - positive numbers",
			params:      Params{a: 2, b: 3},
			expected:    6,
			expectedErr: false,
		},
		{
			name:        "Normal case - negative numbers",
			params:      Params{a: -2, b: -3},
			expected:    6,
			expectedErr: false,
		},
		{
			name:        "Normal case - positive and negative",
			params:      Params{a: 2, b: -3},
			expected:    -6,
			expectedErr: false,
		},
		{
			name:        "Edge case - multiplying by zero",
			params:      Params{a: 5, b: 0},
			expected:    0,
			expectedErr: false,
		},
		{
			name:        "Edge case - multiplying by one",
			params:      Params{a: 5, b: 1},
			expected:    5,
			expectedErr: false,
		},
		{
			name:        "Edge case - large numbers",
			params:      Params{a: 1000, b: 1000},
			expected:    1000000,
			expectedErr: false,
		},
	}

	// Execute tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Execute the function under test
			result := calc.Multiply(tc.params.a, tc.params.b)

			// Verify result
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	// Define test case structure
	type Params struct {
		a, b int
	}
	type testCase struct {
		name        string
		params      Params
		expected    float64
		expectedErr bool
	}

	// Define test case table
	testCases := []testCase{
		{
			name:        "Normal case - positive numbers",
			params:      Params{a: 6, b: 3},
			expected:    2.0,
			expectedErr: false,
		},
		{
			name:        "Normal case - decimal result",
			params:      Params{a: 7, b: 3},
			expected:    2.333333333333333,
			expectedErr: false,
		},
		{
			name:        "Normal case - negative numbers",
			params:      Params{a: -6, b: -3},
			expected:    2.0,
			expectedErr: false,
		},
		{
			name:        "Normal case - positive and negative",
			params:      Params{a: 6, b: -3},
			expected:    -2.0,
			expectedErr: false,
		},
		{
			name:        "Edge case - dividing by one",
			params:      Params{a: 5, b: 1},
			expected:    5.0,
			expectedErr: false,
		},
		{
			name:        "Error case - division by zero",
			params:      Params{a: 5, b: 0},
			expected:    0.0,
			expectedErr: true,
		},
	}

	// Execute tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Execute the function under test
			result, err := calc.Divide(tc.params.a, tc.params.b)

			// Verify error
			if tc.expectedErr {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
				return
			}

			// Verify result
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if math.Abs(result-tc.expected) > 1e-9 {
				t.Errorf("Expected %f, but got %f", tc.expected, result)
			}
		})
	}
}

func TestAddWithLargeNumbers(t *testing.T) {
	// This test will fail because it expects a result that the current implementation can't provide
	result := calc.Add(1000000, 2000000)
	expected := 3000001  // Intentionally wrong expected value to break the build
	
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
