package minimumIntervalToIncludeEachQuery

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	intervals [][]int
	queries   []int
	output    []int
}{
	{
		intervals: [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}},
		queries:   []int{2, 3, 4, 5},
		output:    []int{3, 3, 1, 4},
	},
	{
		intervals: [][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}},
		queries:   []int{2, 19, 5, 22},
		output:    []int{2, -1, 4, 6},
	},
	{
		intervals: [][]int{{9, 9}, {6, 7}, {5, 6}, {2, 5}, {3, 3}},
		queries:   []int{6, 1, 1, 1, 9},
		output:    []int{2, -1, -1, -1, 1},
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].intervals), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				minInterval(testTable[i].intervals, testTable[i].queries),
			)
		})
	}
}

func BenchmarkDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minIntervalDefault([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22})
	}
}

// BenchmarkDefault-10      2997714               394.6 ns/op           328 B/op      14 allocs/op
// Benchmark-10             7757434               153.5 ns/op            64 B/op       4 allocs/op
