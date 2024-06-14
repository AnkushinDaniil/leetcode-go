package mergeIntervals

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	intervals [][]int
	output    int
}{
	{
		intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		output:    1,
	},
	{
		intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
		output:    2,
	},
	{
		intervals: [][]int{{1, 2}, {2, 3}},
		output:    0,
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].intervals), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				eraseOverlapIntervals(testTable[i].intervals),
			)
		})
	}
}
