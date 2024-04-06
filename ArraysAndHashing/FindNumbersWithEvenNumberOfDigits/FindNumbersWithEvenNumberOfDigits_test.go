package FindNumbersWithEvenNumberOfDigits

import (
	"fmt"
	"testing"
)

func TestFindNumbers(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := FindNumbers(tt.nums)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindNumbers([]int{1, 1, 0, 1, 1, 1})
	}
}
