package validParentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "()",
			want: true,
		},
		{
			name: "Example 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Example 3",
			s:    "(]",
			want: false,
		},
		{
			name: "Example 4",
			s:    "[",
			want: false,
		},
		{
			name: "Example 5",
			s:    "]",
			want: false,
		},
		{
			name: "Example 6",
			s:    "(){}}{",
			want: false,
		},
		{
			name: "Example 7",
			s:    "[[[[[[[[[[[[[[[[[[[",
			want: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, isValid(testCase.s), testCase.want)
		})
	}
}
