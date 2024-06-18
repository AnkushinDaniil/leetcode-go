package jumpGameII

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	nums   []int
	output int
}{
	{
		nums:   []int{2, 3, 1, 1, 4},
		output: 2,
	},
	{
		nums:   []int{2, 3, 0, 1, 4},
		output: 2,
	},
	{
		nums:   []int{0},
		output: 0,
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].nums), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				jump(testTable[i].nums),
			)
		})
	}
}

func Benchmark_jump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3})
	}
}

func Benchmark_jumpLCT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jumpLCT([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3})
	}
}

func Benchmark_jumpLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jumpLCM([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3})
	}
}

// Benchmark_jump-10               131386281                9.018 ns/op             0 B/op          0 allocs/op
// Benchmark_jumpLCT-10            127872496                9.290 ns/op             0 B/op          0 allocs/op
// Benchmark_jumpLCM-10            131428699                9.044 ns/op             0 B/op          0 allocs/op
