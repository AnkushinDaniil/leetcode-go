package ReplaceElementsWithGreatestElementOnRightSide

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReplaceElements(t *testing.T) {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
		{[]int{400}, []int{-1}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.arr)
		t.Run(testname, func(t *testing.T) {
			ans := ReplaceElements(tt.arr)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReplaceElements([]int{1, 1, 0, 1, 1, 1})
	}
}
