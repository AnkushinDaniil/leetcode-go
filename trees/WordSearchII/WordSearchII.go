package wordSearchII

// type TrieNode struct {
// 	val   struct{}
// 	links [26]*TrieNode
// }

type Trie struct {
	links [26]*Trie
	end   bool
}

func Constructor() Trie {
	return Trie{}
}

var i, j int

func (this *Trie) Insert(word string) {
	cur := this
	_ = word[len(word)-1]
	for i = 0; i < len(word); i++ {
		index := word[i] - 'a'
		if cur.links[index] == nil {
			cur.links[index] = &Trie{}
		}
		cur = cur.links[index]
	}
	cur.end = true
}

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for j = len(words) - 1; j >= 0; j-- {
		trie.Insert(words[j])
	}

	var dfs func(int, int, *Trie)
	table := make(map[string]struct{}, 64)
	checked := make(map[[2]int]struct{}, len(board)*len(board[0]))
	stack := make([]byte, len(board)*len(board[0]))
	stackLen := 0

	dfs = func(c, r int, t *Trie) {
		if _, ok := checked[[2]int{c, r}]; ok || c < 0 || c >= len(board[0]) || r < 0 ||
			r >= len(board) || t == nil {
			return
		}
		t = t.links[board[r][c]-'a']
		if t == nil {
			return
		}
		stack[stackLen] = board[r][c]
		stackLen++
		if t.end {
			table[string(stack[:stackLen])] = struct{}{}
		}
		checked[[2]int{c, r}] = struct{}{}
		dfs(c+1, r, t)
		dfs(c, r+1, t)
		dfs(c-1, r, t)
		dfs(c, r-1, t)
		delete(checked, [2]int{c, r})
		stackLen--
	}
	for y := len(board) - 1; y >= 0; y-- {
		for x := len(board[0]) - 1; x >= 0; x-- {
			dfs(x, y, &trie)
		}
	}
	res := make([]string, 0, len(table))

	for key := range table {
		res = append(res, key)
	}

	return res
}

// func (this *Trie) Search(word string) bool {
// 	cur := this
// 	_ = word[len(word)-1]
// 	for i = 0; i < len(word); i++ {
// 		index = word[i] - 'a'
// 		if cur.links[index] == nil {
// 			return false
// 		}
// 		cur = cur.links[index]
// 	}
// 	return cur.end
// }

// func (this *Trie) StartsWith(prefix string) bool {
// 	cur := this
// 	_ = prefix[len(prefix)-1]
// 	for i = 0; i < len(prefix); i++ {
// 		index = prefix[i] - 'a'
// 		if cur.links[index] == nil {
// 			return false
// 		}
// 		cur = cur.links[index]
// 	}
// 	return true
// }
