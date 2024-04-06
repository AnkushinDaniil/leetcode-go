package bestTimeToBuyAndSellStock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		nums   []int
		output int
	}{
		{
			name:   "Example 1",
			nums:   []int{7, 1, 5, 3, 6, 4},
			output: 5,
		},
		{
			name:   "Example 2",
			nums:   []int{7, 6, 4, 3, 1},
			output: 0,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, maxProfit(testCase.nums), testCase.output)
		})
	}
}
