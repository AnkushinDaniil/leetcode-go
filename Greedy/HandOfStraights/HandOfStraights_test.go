package handOfStraights

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	hand      []int
	groupSize int
	output    bool
}{
	{
		hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
		groupSize: 3,
		output:    true,
	},
	{
		hand:      []int{11, 2, 3, 4, 5},
		groupSize: 4,
		output:    false,
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].hand), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				isNStraightHand(testTable[i].hand, testTable[i].groupSize),
			)
		})
	}
}
