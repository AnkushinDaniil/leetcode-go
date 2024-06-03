package houseRobberII

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	nums   []int
	output int
}{
	{
		nums:   []int{2, 3, 2},
		output: 3,
	},
	{
		nums:   []int{1, 2, 3, 1},
		output: 4,
	},
	{
		nums:   []int{1, 2, 3},
		output: 3,
	},
}

func Test(t *testing.T) {
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				rob(testCase.nums),
			)
		})
	}
}
