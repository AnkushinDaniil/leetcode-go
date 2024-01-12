package DuplicateZeros

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 0, 1, 2}},
		{[]int{0, 1, 0, 0, 2, 3}, []int{0, 0, 1, 0, 0, 0}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := DuplicateZeros(tt.nums)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DuplicateZeros([]int{1, 1, 0, 1, 1, 1})
	}
}
