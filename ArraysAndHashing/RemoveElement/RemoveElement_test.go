package removeElement

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		nums []int
		val  int
		want []int
	}{
		{[]int{3, 2, 2, 3}, 3, []int{2, 2}},
		{[]int{3, 2, 2, 3}, 4, []int{3, 2, 2, 3}},
		{[]int{3, 3, 3, 3}, 3, []int{}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}},
		{[]int{2}, 3, []int{2}},
		{[]int{1}, 1, []int{}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d, %d", tt.nums, tt.val)
		t.Run(testname, func(t *testing.T) {
			res := RemoveElement(tt.nums, tt.val)
			if !reflect.DeepEqual(tt.nums[:res], tt.want) {
				t.Errorf("got %d, want %d", tt.nums[:res], tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveElement([]int{3, 2, 2, 3}, 3)
	}
}
