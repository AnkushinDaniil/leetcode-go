package SortedSquares

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := SortedSquares(tt.nums)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortedSquares([]int{1, 1, 0, 1, 1, 1})
	}
}
