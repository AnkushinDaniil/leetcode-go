package mergesortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{5, 6, 7, 0, 0, 0}, 3, []int{2, 3, 4}, 3, []int{2, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{3, 4, 5}, 3, []int{1, 2, 3, 3, 4, 5}},
		{[]int{1, 2, 2, 0, 0, 0}, 3, []int{2, 2, 4}, 3, []int{1, 2, 2, 2, 2, 4}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%+v", tt)
		t.Run(testname, func(t *testing.T) {
			Merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("got %d, want %d", tt.nums1, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	}
}
