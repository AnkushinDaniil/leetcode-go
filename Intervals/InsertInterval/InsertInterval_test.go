package insertInterval

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	intervals   [][]int
	newInterval []int
	output      [][]int
}{
	{
		intervals:   [][]int{{1, 3}, {6, 9}},
		newInterval: []int{2, 5},
		output:      [][]int{{1, 5}, {6, 9}},
	},
	{
		intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
		newInterval: []int{4, 8},
		output:      [][]int{{1, 2}, {3, 10}, {12, 16}},
	},
	{
		intervals:   [][]int{{6, 7}, {8, 10}, {12, 16}},
		newInterval: []int{1, 2},
		output:      [][]int{{1, 2}, {6, 7}, {8, 10}, {12, 16}},
	},
	{
		intervals:   [][]int{{1, 5}},
		newInterval: []int{5, 7},
		output:      [][]int{{1, 7}},
	},
	{
		intervals:   [][]int{{1, 5}},
		newInterval: []int{0, 3},
		output:      [][]int{{0, 5}},
	},
	{
		intervals:   [][]int{{0, 5}, {9, 12}},
		newInterval: []int{7, 16},
		output:      [][]int{{0, 5}, {7, 16}},
	},
	{
		intervals:   [][]int{{1, 5}, {9, 12}},
		newInterval: []int{0, 4},
		output:      [][]int{{0, 5}, {9, 12}},
	},
	{
		intervals:   [][]int{},
		newInterval: []int{5, 7},
		output:      [][]int{{5, 7}},
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].newInterval), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				insert(testTable[i].intervals, testTable[i].newInterval),
			)
		})
	}
}

func Benchmark_insertOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range testTable {
			insertOld(testTable[j].intervals, testTable[j].newInterval)
		}
	}
}

func Benchmark_insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range testTable {
			insert(testTable[j].intervals, testTable[j].newInterval)
		}
	}
}

func Benchmark_insertLCT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range testTable {
			insertLCT(testTable[j].intervals, testTable[j].newInterval)
		}
	}
}

func Benchmark_insertLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range testTable {
			insertLCM(testTable[j].intervals, testTable[j].newInterval)
		}
	}
}

// Benchmark_insertOld-10           6704220               172.8 ns/op           168 B/op          2 allocs/op
// Benchmark_insert-10              6734072               177.5 ns/op           168 B/op          2 allocs/op
// Benchmark_insertLCT-10           1366977               860.8 ns/op           928 B/op         23 allocs/op
// Benchmark_insertLCM-10           5761681               205.4 ns/op           248 B/op          3 allocs/op
