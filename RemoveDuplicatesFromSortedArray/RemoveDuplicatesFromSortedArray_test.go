package removeDuplicatesFromSortedArray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 2}, []int{1, 2}},
		{[]int{1, 1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{0}, []int{0}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.nums)
		t.Run(testname, func(t *testing.T) {
			res := RemoveDuplicates(tt.nums)
			if !reflect.DeepEqual(tt.nums[:res], tt.want) {
				t.Errorf("got %d, want %d", tt.nums[:res], tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDuplicates([]int{3, 2, 2, 3})
	}
}
