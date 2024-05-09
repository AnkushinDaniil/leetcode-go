package wordSearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		board  [][]byte
		word   string
		output bool
	}{
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "ABCCED",
			output: true,
		},
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "SEE",
			output: true,
		},
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "ABCB",
			output: false,
		},
	}
	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, exist(testCase.board, testCase.word))
		})
	}
}

var board = [][]byte{
	{'A', 'A', 'A', 'A', 'A', 'A'},
	{'A', 'A', 'A', 'A', 'A', 'A'},
	{'A', 'A', 'A', 'A', 'A', 'A'},
	{'A', 'A', 'A', 'A', 'A', 'A'},
	{'A', 'A', 'A', 'A', 'A', 'B'},
	{'A', 'A', 'A', 'A', 'B', 'A'},
}

var word = "AAAAAAAAAAAAABB"

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		existSimple(board, word)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exist(board, word)
	}
}

// BenchmarkSimple-10            73          16159390 ns/op               0 B/op          0 allocs/op
// Benchmark-10             5760918               206.1 ns/op             0 B/op          0 allocs/op
