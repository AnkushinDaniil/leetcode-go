package carFleet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name     string
		target   int
		position []int
		speed    []int
		output   int
	}{
		{
			name:     "Example 1",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			output:   3,
		},
		{
			name:     "Example 2",
			target:   10,
			position: []int{3},
			speed:    []int{3},
			output:   1,
		},
		{
			name:     "Example 3",
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			output:   1,
		},
		{
			name:     "Example 4",
			target:   20,
			position: []int{6, 2, 17},
			speed:    []int{3, 9, 2},
			output:   2,
		},
		{
			name:     "Example 5",
			target:   13,
			position: []int{10, 2, 5, 7, 4, 6, 11},
			speed:    []int{7, 5, 10, 5, 9, 4, 1},
			output:   2,
		},
		{
			name:     "Example 6",
			target:   21,
			position: []int{1, 15, 6, 8, 18, 14, 16, 2, 19, 17, 3, 20, 5},
			speed:    []int{8, 5, 5, 7, 10, 10, 7, 9, 3, 4, 4, 10, 2},
			output:   7,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				carFleet(testCase.target, testCase.position, testCase.speed),
				testCase.output,
			)
		})
	}
}
