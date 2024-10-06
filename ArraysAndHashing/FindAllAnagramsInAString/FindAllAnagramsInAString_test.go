package findallanagramsinastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
    {"aaaaaaaaaa", "aaaaaaaaaaaaa", []int{}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("findAnagrams(%q, %q)", test.s, test.p), func(t *testing.T) {
			assert.Equal(t, test.want, findAnagrams(test.s, test.p))
		})
	}
}
