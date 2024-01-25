package CheckIfNAndItsDoubleExist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckIfExist(t *testing.T) {
	var tests = []struct {
		arr  []int
		want bool
	}{
		{[]int{10, 2, 5, 3}, true},
		{[]int{3, 1, 7, 11}, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.arr)
		t.Run(testname, func(t *testing.T) {
			ans := CheckIfExist(tt.arr)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckIfExist([]int{1, 1, 0, 1, 1, 1})
	}
}
