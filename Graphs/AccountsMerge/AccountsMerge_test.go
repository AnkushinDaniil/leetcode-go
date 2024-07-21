package accountsMerge

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	accounts [][]string
	output   [][]string
}

var tests = []test{
	// {
	// 	accounts: [][]string{
	// 		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
	// 		{"John", "johnsmith@mail.com", "john00@mail.com"},
	// 		{"Mary", "mary@mail.com"},
	// 		{"John", "johnnybravo@mail.com"},
	// 	},
	// 	output: [][]string{
	// 		{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
	// 		{"Mary", "mary@mail.com"},
	// 		{"John", "johnnybravo@mail.com"},
	// 	},
	// },
	// {
	// 	accounts: [][]string{
	// 		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
	// 		{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
	// 		{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
	// 		{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
	// 		{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
	// 	},
	// 	output: [][]string{
	// 		{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
	// 		{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
	// 		{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
	// 		{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
	// 		{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
	// 	},
	// },
	{
		accounts: [][]string{
			{"David", "David0@m.co", "David1@m.co"},
			{"David", "David3@m.co", "David4@m.co"},
			{"David", "David4@m.co", "David5@m.co"},
			{"David", "David2@m.co", "David3@m.co"},
			{"David", "David1@m.co", "David2@m.co"},
		},
		output: [][]string{
			{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
		},
	},
}

func Test(t *testing.T) {
	for i, testCase := range tests {
		slices.SortFunc(testCase.output, sortFunc)
		res := accountsMerge(testCase.accounts)
		slices.SortFunc(res, sortFunc)
		t.Run(string('0'+byte(i)), func(t *testing.T) {
			assert.Equal(t, testCase.output, res)
		})
	}
}

func sortFunc(a, b []string) int {
	if res := strings.Compare(a[0], b[0]); res == 0 {
		return sortFunc(a[1:], b[1:])
	} else {
		return res
	}
}
