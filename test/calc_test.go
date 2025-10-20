package calc_test

import (
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
