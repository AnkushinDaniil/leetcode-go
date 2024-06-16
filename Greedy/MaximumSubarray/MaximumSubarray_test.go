package maximumSubarray

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	nums   []int
	output int
}{
	{
		nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		output: 6,
	},
	{
		nums:   []int{1},
		output: 1,
	},
	{
		nums:   []int{5, 4, -1, 7, 8},
		output: 23,
	},
	{
		nums:   []int{-1},
		output: -1,
	},
	{
		nums:   []int{-2, 1},
		output: 1,
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].nums), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				maxSubArray(testTable[i].nums),
			)
		})
	}
}
