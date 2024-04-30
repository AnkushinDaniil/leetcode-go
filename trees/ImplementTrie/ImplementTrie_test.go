package implementTrie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		funcs  []string
		params []string
		output []interface{}
	}{
		{
			name: "Example 1",
			funcs: []string{
				"Trie",
				"insert",
				"search",
				"search",
				"startsWith",
				"insert",
				"search",
			},
			params: []string{"", "apple", "apple", "app", "app", "app", "app"},
			output: []interface{}{nil, nil, true, false, true, nil, true},
		},
		{
			name: "Example 2",
			funcs: []string{
				"Trie",
				"insert",
				"insert",
				"insert",
				"insert",
				"insert",
				"insert",
				"search",
				"search",
				"search",
				"search",
				"search",
				"search",
				"search",
				"search",
				"search",
				"startsWith",
				"startsWith",
				"startsWith",
				"startsWith",
				"startsWith",
				"startsWith",
				"startsWith",
				"startsWith",
				"startsWith",
			},
			params: []string{
				"",
				"app",      // insert nil
				"apple",    // insert nil
				"beer",     // insert nil
				"add",      // insert nil
				"jam",      // insert nil
				"rental",   // insert nil
				"apps",     // search false
				"app",      // search true
				"ad",       // search false
				"applepie", // search false
				"rest",     // search false
				"jan",      // search false
				"rent",     // search false
				"beer",     // search true
				"jam",      // search true
				"apps",     // startsWith false
				"app",      // startsWith true
				"ad",       // startsWith true
				"applepie", // startsWith false
				"rest",     // startsWith false
				"jan",      // startsWith false
				"rent",     // startsWith true
				"beer",     // startsWith true
				"jam",      // startsWith true
			},
			output: []interface{}{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				false,
				true,
				false,
				false,
				false,
				false,
				false,
				true,
				true,
				false,
				true,
				true,
				false,
				false,
				false,
				true,
				true,
				true,
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var trie Trie
			res := make([]interface{}, len(testCase.funcs))
			for i := 0; i < len(testCase.funcs); i++ {
				switch testCase.funcs[i] {
				case "Trie":
					trie = Constructor()
					res[i] = nil
				case "insert":
					trie.Insert(testCase.params[i])
					res[i] = nil
				case "search":
					res[i] = trie.Search(testCase.params[i])
				case "startsWith":
					res[i] = trie.StartsWith(testCase.params[i])
				}
			}
			assert.Equal(t, testCase.output, res)
		})
	}
}
