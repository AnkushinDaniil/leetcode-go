package longestIncreasingSubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	nums     []int
	wordDict []string
	output   int
}{
	{
		nums:   []int{10, 9, 2, 5, 3, 7, 101, 18},
		output: 4,
	},
	{
		nums:   []int{0, 1, 0, 3, 2, 3},
		output: 4,
	},
	{
		nums:   []int{7, 7, 7, 7, 7, 7, 7},
		output: 1,
	},
	{
		nums:   []int{4, 10, 4, 3, 8, 9},
		output: 3,
	},
	{
		nums:   []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
		output: 6,
	},
}

func Test(t *testing.T) {
	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("Test %v", testCase.nums), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				lengthOfLIS(testCase.nums),
			)
		})
	}
}
