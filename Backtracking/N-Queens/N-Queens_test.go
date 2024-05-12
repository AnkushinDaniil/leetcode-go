package nQueens

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		n      int
		output [][]string
	}{
		{
			n: 4,
			output: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			n: 1,
			output: [][]string{
				{"Q"},
			},
		},
	}

	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, solveNQueens(testCase.n))
		})
	}
}

// func Benchmark1d(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		solveNQueens1d(9)
// 	}
// }

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveNQueens(9)
	}
}

func Benchmark2d(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveNQueens2d(9)
	}
}

func BenchmarkNeetcode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveNQueensNeetcode(9)
	}
}

func BenchmarkLeetcode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveNQueensLeetcode(9)
	}
}

// Benchmark-10                        3308            358937 ns/op           83472 B/op        375 allocs/op
// Benchmark2d-10                       940           1266957 ns/op          133780 B/op       3537 allocs/op
// BenchmarkNeetcode-10                 274           4288334 ns/op          488297 B/op      67518 allocs/op
// BenchmarkLeetcode-10                 620           1916975 ns/op          480186 B/op      25191 allocs/op
