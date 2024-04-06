package validAnagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	testTable := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "Example 2",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "Example 3",
			s:    "a",
			t:    "ab",
			want: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, isAnagram(testCase.s, testCase.t), testCase.want)
		})
	}
}
