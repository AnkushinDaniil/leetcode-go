package jumpGame

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	nums   []int
	output bool
}{
	{
		nums:   []int{2, 3, 1, 1, 4},
		output: true,
	},
	{
		nums:   []int{3, 2, 1, 0, 4},
		output: false,
	},
	{
		nums:   []int{0, 1},
		output: false,
	},
	{
		nums:   []int{0},
		output: true,
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].nums), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				canJump(testTable[i].nums),
			)
		})
	}
}
