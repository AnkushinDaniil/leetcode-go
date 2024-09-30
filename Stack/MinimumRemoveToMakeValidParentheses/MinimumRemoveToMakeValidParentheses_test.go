package minimumremovetomakevalidparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
		{"(a(b(c)d)", "a(b(c)d)"},
	}

	for _, test := range tests {
		if minRemoveToMakeValid(test.input) != test.expected {
			assert.Equal(t, test.expected, minRemoveToMakeValid(test.input), "Test case failed: %v", test)
		}
	}
}
