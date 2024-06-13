package mergeIntervals

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	intervals [][]int
	output    [][]int
}{
	{
		intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		output:    [][]int{{1, 6}, {8, 10}, {15, 18}},
	},
	{
		intervals: [][]int{{1, 4}, {4, 5}},
		output:    [][]int{{1, 5}},
	},
	{
		intervals: [][]int{{1, 4}, {0, 4}},
		output:    [][]int{{0, 4}},
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].intervals), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				merge(testTable[i].intervals),
			)
		})
	}
}

func forBenchmark(function func([][]int) [][]int) {
	for i := range testTable {
		function(testTable[i].intervals)
	}
}

func Benchmark_merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forBenchmark(merge)
	}
}

func Benchmark_mergeJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forBenchmark(mergeJ)
	}
}

func Benchmark_mergeInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forBenchmark(mergeInsert)
	}
}

// Benchmark_merge-10              24146710                44.78 ns/op            0 B/op        0 allocs/op
// Benchmark_mergeJ-10             21600393                54.78 ns/op            0 B/op        0 allocs/op
// Benchmark_mergeInsert-10         7925052               153.3 ns/op           192 B/op        3 allocs/op
