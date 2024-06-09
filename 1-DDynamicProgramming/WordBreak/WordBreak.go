package wordBreak

func wordBreak(s string, wordDict []string) bool {
	memo := make([]bool, len(s)+1)
	memo[0] = true
	for i := 0; i <= len(s); i++ {
		if memo[i] {
			for j := range wordDict {
				if i+len(wordDict[j]) <= len(s) && !memo[i+len(wordDict[j])] {
					memo[i+len(wordDict[j])] = s[i:i+len(wordDict[j])] == wordDict[j]
				}
			}
		}
	}
	return memo[len(memo)-1]
}

func wordBreakArray(s string, wordDict []string) bool {
	memo := [301]bool{}
	memo[0] = true
	for i := 0; i <= len(s); i++ {
		if memo[i] {
			for j := range wordDict {
				if i+len(wordDict[j]) <= len(s) && !memo[i+len(wordDict[j])] {
					memo[i+len(wordDict[j])] = s[i:i+len(wordDict[j])] == wordDict[j]
				}
			}
		}
	}
	return memo[len(s)]
}

func wordBreakNeetcode(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if (i+len(w)) <= len(s) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}

func wordBreakLeetcode(s string, wordDict []string) bool {
	t := new(trie)
	for _, word := range wordDict {
		t.addWord(word)
	}

	return search(s, 0, t, make(map[int]bool))
}

func search(s string, idx int, t *trie, memo map[int]bool) bool {
	l := len(s)
	if idx == l {
		return true
	} else if v, ok := memo[idx]; ok {
		return v
	}

	res := false
	for i := idx; i < l; i++ {
		prefix := s[idx : i+1]
		if !t.hasPrefix(prefix) {
			break
		}

		if !t.isWord(prefix) {
			continue
		}

		res = res || search(s, i+1, t, memo)

		if res {
			break
		}
	}

	memo[idx] = res
	return res
}

type trie struct {
	root *node
}

func (t *trie) addWord(word string) {
	t.root = t.addWordAt(t.root, word, 0)
}

func (t *trie) addWordAt(curr *node, word string, idx int) *node {
	if curr == nil {
		curr = new(node)
	}

	l := len(word)
	if idx == l {
		curr.isWord = true
	} else {
		p := word[idx] - 'a'
		curr.links[p] = t.addWordAt(curr.links[p], word, idx+1)
	}

	return curr
}

func (t *trie) get(curr *node, word string, idx int) *node {
	if curr == nil {
		return nil
	}

	l := len(word)
	if idx == l {
		return curr
	} else {
		p := word[idx] - 'a'
		return t.get(curr.links[p], word, idx+1)
	}
}

func (t *trie) isWord(word string) bool {
	n := t.get(t.root, word, 0)
	return n != nil && n.isWord
}

func (t *trie) hasPrefix(prefix string) bool {
	n := t.get(t.root, prefix, 0)
	return n != nil
}

const alphabetSize = 26

type node struct {
	isWord bool
	links  [alphabetSize]*node
}
