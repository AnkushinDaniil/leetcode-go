package minCostClimbingStairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	cost   []int
	output int
}{
	{
		cost:   []int{10, 15, 20},
		output: 15,
	},
	{
		cost:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		output: 6,
	},
}

func Test(t *testing.T) {
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				minCostClimbingStairs(testCase.cost),
			)
		})
	}
}
