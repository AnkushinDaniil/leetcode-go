package wordSearchII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		board  [][]byte
		words  []string
		output []string
	}{
		{
			name: "Example 1",
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words:  []string{"oath", "pea", "eat", "rain"},
			output: []string{"eat", "oath"},
		},
		{
			name: "Example 2",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words:  []string{"abcb"},
			output: []string{},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.output, findWords(testCase.board, testCase.words))
		})
	}
}
