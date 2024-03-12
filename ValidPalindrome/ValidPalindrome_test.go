package validPalindrome

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
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Example 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "Example 3",
			s:    "0P",
			want: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, isPalindrome(testCase.s), testCase.want)
		})
	}
}
