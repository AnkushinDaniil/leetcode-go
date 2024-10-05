package subsets

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		nums   []int
		output [][]int
	}{
		{
			nums:   []int{1, 2, 3},
			output: [][]int{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}},
		},
	}

  cmp := func(a, b []int) int {
    m := min(len(a), len(b))
    for i := 0; i < m; i++ {
      if a[i] < b[i] {
        return -1
      } else if a[i] > b[i] {
        return 1
      }
    }
    return len(a) - len(b)
  }
    
	for i, testCase := range testTable {
    slices.SortFunc(testCase.output, cmp)
    res := subsets(testCase.nums)
    slices.SortFunc(res, cmp)
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, res)
		})
	}
}
