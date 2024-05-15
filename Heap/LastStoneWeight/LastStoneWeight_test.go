package lastStoneWeight

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		stones []int
		output int
	}{
		{
			stones: []int{2, 7, 4, 1, 8, 1},
			output: 1,
		},
		{
			stones: []int{1},
			output: 1,
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, lastStoneWeight(testCase.stones))
		})
	}
}

// func BenchmarkHeap(b *testing.B) {
// 	params := []int{0, 3, 5, 10, 9, 4}
// 	heap := Constructor(3, []int{4, 5, 8, 2})
// 	for i := 0; i < b.N; i++ {
// 		heap.Add(params[i%len(params)])
// 	}
// }
