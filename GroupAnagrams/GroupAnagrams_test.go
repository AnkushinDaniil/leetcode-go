package groupAnagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	testTable := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "Example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "Example 2",
			strs: []string{""},
			want: [][]string{
				{""},
			},
		},
		{
			name: "Example 2",
			strs: []string{"a"},
			want: [][]string{
				{"a"},
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, groupAnagrams(testCase.strs), testCase.want)
		})
	}
}
