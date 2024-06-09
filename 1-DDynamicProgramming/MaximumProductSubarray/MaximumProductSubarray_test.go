package maximumProductSubarray

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
		nums:   []int{2, 3, -2, 4},
		output: 6,
	},
	{
		nums:   []int{-2, 0, -1},
		output: 0,
	},
	{
		nums:   []int{-2, 3, -4},
		output: 24,
	},
	{
		nums:   []int{2, 0, -1},
		output: 2,
	},
	{
		nums: []int{
			0,
			10,
			10,
			10,
			10,
			10,
			10,
			10,
			10,
			10,
			-10,
			10,
			10,
			10,
			10,
			10,
			10,
			10,
			10,
			10,
			0,
		},
		output: 1000000000,
	},
}

func Test(t *testing.T) {
	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("Test %v", testCase.nums), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				maxProduct(testCase.nums),
			)
		})
	}
}

func Benchmark_maxProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProduct([]int{2, 3, -2, 4})
	}
}
