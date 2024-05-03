package wordSearchII

// type TrieNode struct {
// 	val   struct{}
// 	links [26]*TrieNode
// }

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func Constructor() TrieNode {
	return TrieNode{}
}

func (tn *TrieNode) Insert(word string) {
	cur := tn
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if cur.children[index] == nil {
			cur.children[index] = &TrieNode{}
		}
		cur = cur.children[index]
	}
	cur.isEnd = true
}

func (tn *TrieNode) IsEmpty() bool {
	for i := 0; i < 26; i++ {
		if tn.children[i] != nil {
			return false
		}
	}
	return true
}

func (tn *TrieNode) Delete(word string) bool {
	if len(word) == 0 {
		tn.isEnd = false
		return tn.IsEmpty()
	}
	if tn.children[word[0]-'a'].Delete(word[1:]) {
		tn.children[word[0]-'a'] = nil
		return tn.IsEmpty()
	}
	return false
}

func findWords(board [][]byte, words []string) []string {
	rows, cols := len(board), len(board[0])
	trie := Constructor()
	res := make([]string, 0, 64)

	var dfs func(int, int, *TrieNode)
	stack := make([]byte, len(board)*len(board[0]))
	stackLen := 0

	dfs = func(c, r int, t *TrieNode) {
		if c < 0 || c >= cols || r < 0 ||
			r >= rows || board[r][c] == 0 {
			return
		}
		t = t.children[board[r][c]-'a']

		if t == nil {
			return
		}
		ch := board[r][c]
		board[r][c] = 0

		stack[stackLen] = ch
		stackLen++

		if t.isEnd {
			res = append(res, string(stack[:stackLen]))
			t.isEnd = false
			trie.Delete(string(stack[:stackLen]))
		}

		dfs(c+1, r, t)
		dfs(c, r+1, t)
		dfs(c-1, r, t)
		dfs(c, r-1, t)
		board[r][c] = ch
		stackLen--
	}

	table := make(map[[2]byte]struct{}, rows*(cols-1)+(rows-1)*cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols-1; c++ {
			table[[2]byte{board[r][c], board[r][c+1]}] = struct{}{}
		}
	}
	for r := 0; r < rows-1; r++ {
		for c := 0; c < cols; c++ {
			table[[2]byte{board[r][c], board[r+1][c]}] = struct{}{}
		}
	}

	var (
		ok1, ok2 bool
		i, j     int
	)
	for i = 0; i < len(words); i++ {
	WORD:
		for j = 0; j < len(words[i])-1; j++ {
			_, ok1 = table[[2]byte{words[i][j], words[i][j+1]}]
			_, ok2 = table[[2]byte{words[i][j+1], words[i][j]}]
			if !ok1 && !ok2 {
				break WORD
			}
		}
		trie.Insert(words[i])
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			dfs(x, y, &trie)
		}
	}

	return res
}

// func (tn *Trie) Search(word string) bool {
// 	cur := tn
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

// func (tn *Trie) StartsWith(prefix string) bool {
// 	cur := tn
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
