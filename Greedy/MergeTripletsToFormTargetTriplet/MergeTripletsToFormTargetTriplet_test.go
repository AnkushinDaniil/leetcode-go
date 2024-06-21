package mergeTripletsToFormTargetTriplet

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	triplets [][]int
	target   []int
	output   bool
}{
	{
		triplets: [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}},
		target:   []int{2, 7, 5},
		output:   true,
	},
	{
		triplets: [][]int{{3, 4, 5}, {4, 5, 6}},
		target:   []int{3, 2, 5},
		output:   false,
	},
	{
		triplets: [][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}},
		target:   []int{5, 5, 5},
		output:   true,
	},
	{
		triplets: [][]int{{1, 3, 1}},
		target:   []int{1, 3, 1},
		output:   true,
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].target), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				mergeTripletsConcur(testTable[i].triplets, testTable[i].target),
			)
		})
	}
}

func Benchmark_mergeTriplets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeTriplets(testTable[2].triplets, testTable[2].target)
	}
}

func Benchmark_mergeTripletsConcur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeTripletsConcur(testTable[2].triplets, testTable[2].target)
	}
}

func Benchmark_mergeTripletsOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeTripletsOld(testTable[2].triplets, testTable[2].target)
	}
}

// Benchmark_mergeTriplets-10              35181759                31.25 ns/op              0 B/op          0 allocs/op
// Benchmark_mergeTripletsConcur-10           95900             12093 ns/op     3259 B/op          74 allocs/op
// Benchmark_mergeTripletsOld-10           24720499                47.60 ns/op              0 B/op          0 allocs/op
