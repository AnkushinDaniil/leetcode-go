package MaxConsecutiveOnes

import (
	"fmt"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := FindMaxConsecutiveOnes(tt.nums)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
	}
}
