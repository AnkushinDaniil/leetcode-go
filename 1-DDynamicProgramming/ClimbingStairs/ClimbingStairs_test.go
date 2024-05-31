package climbingStairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	n      int
	output int
}{
	{
		n:      2,
		output: 2,
	},
	{
		n:      3,
		output: 3,
	},
}

func Test(t *testing.T) {
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				climbStairs(testCase.n),
			)
		})
	}
}
