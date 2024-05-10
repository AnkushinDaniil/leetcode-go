package palindromePartitioning

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		s      string
		output [][]string
	}{
		{
			s:      "aab",
			output: [][]string{{"aa", "b"}, {"a", "a", "b"}},
		},
		{
			s:      "a",
			output: [][]string{{"a"}},
		},
	}

	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, partitionIter(testCase.s))
		})
	}
}

func BenchmarkIter(b *testing.B) {
	s := "aabaa"
	for i := 0; i < b.N; i++ {
		partitionIter(s)
	}
}

func BenchmarkBacktrack(b *testing.B) {
	s := "aabaa"
	for i := 0; i < b.N; i++ {
		partition(s)
	}
}
