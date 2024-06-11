package partitionEqualSubsetSum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	nums   []int
	output bool
}{
	{
		nums:   []int{1, 5, 11, 5},
		output: true,
	},
	{
		nums:   []int{1, 2, 3, 5},
		output: false,
	},
	{
		nums:   []int{1, 2, 5},
		output: false,
	},
	{
		nums:   []int{14, 9, 8, 4, 3, 2},
		output: true,
	},
}

func Test(t *testing.T) {
	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("Test %v", testCase.nums), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				canPartition(testCase.nums),
			)
		})
	}
}
