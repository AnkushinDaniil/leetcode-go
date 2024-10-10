package sortanarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Test case 1",
			input:    []int{5, 2, 3, 1},
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "Test case 2",
			input:    []int{5, 1, 1, 2, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name:     "Test case 3",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, sortArray(test.input))
		})
	}
}
