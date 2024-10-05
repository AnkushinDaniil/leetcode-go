package stringtointeger

import "testing"

func TestMyAToi(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words",   4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"9223372036854775808", 2147483647},
		{"000000000001", 1},
		{"-000000000001", -1},
		{"-a", 0},
		{"-1", -1},
    {"+1", 1},
    {"+1a", 1},
    {"+1 1", 1},
    {"1-1", 1},
	}

	for _, test := range tests {
		if got := myAtoi(test.input); got != test.output {
			t.Errorf("myAtoi(%s) = %d; want %d", test.input, got, test.output)
		}
	}
}
