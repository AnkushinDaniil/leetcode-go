package ValidMountainArray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidMountainArray(t *testing.T) {
	var tests = []struct {
		arr  []int
		want bool
	}{
		{[]int{2}, false},
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{0, 3, 2, 1}, true},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.arr)
		t.Run(testname, func(t *testing.T) {
			ans := ValidMountainArray(tt.arr)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidMountainArray([]int{1, 1, 0, 1, 1, 1})
	}
}
