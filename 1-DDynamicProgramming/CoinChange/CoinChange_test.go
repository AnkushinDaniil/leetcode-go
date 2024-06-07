package coinChange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	coins  []int
	amount int
	output int
}{
	{
		coins:  []int{1, 2, 5},
		amount: 11,
		output: 3,
	},
	{
		coins:  []int{2},
		amount: 3,
		output: -1,
	},
	{
		coins:  []int{1},
		amount: 0,
		output: 0,
	},
}

func Test(t *testing.T) {
	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("Test %v", testCase.coins), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				coinChange(testCase.coins, testCase.amount),
			)
		})
	}
}
