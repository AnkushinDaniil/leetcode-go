package minimumoperationstomakemedianofarrayequaltok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		k      int
		nums   []int
		output int64
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 5, 6, 8, 5},
			k:      4,
			output: 2,
		},
		{
			name:   "Example 2",
			nums:   []int{2, 5, 6, 8, 5},
			k:      7,
			output: 3,
		},
		{
			name:   "Example 3",
			nums:   []int{1, 2, 3, 4, 5, 6},
			k:      4,
			output: 0,
		},
		{
			name:   "Example 4",
			nums:   []int{1},
			k:      1000000000,
			output: 999999999,
		},
		{
			name:   "Example 5",
			nums:   []int{1000000000},
			k:      1,
			output: 999999999,
		},
		{
			name:   "Example 6",
			nums:   []int{1000000000, 1000000000, 1000000000},
			k:      1,
			output: 1999999998,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, minOperationsToMakeMedianK(testCase.nums, testCase.k), testCase.output)
		})
	}
}
