package implementTrie

type Trie struct {
	isEnd    bool
	shared   string
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	i := 0
	for ; i < len(word) && i < len(t.shared) && word[i] == t.shared[i]; i++ {
	}
	if i == len(word) && i == len(t.shared) {
		t.isEnd = true
	}

	if i < len(t.shared) {
		idx := t.shared[i] - 'a'
		temp := t.children
		t.children = [26]*Trie{}
		t.children[idx] = &Trie{
			isEnd:    t.isEnd,
			shared:   t.shared[i:],
			children: temp,
		}
		t.shared = t.shared[:i]
		t.isEnd = i == len(word)
	}

	if i < len(word) {
		idx := word[i] - 'a'
		if t.children[idx] == nil {
			t.children[idx] = &Trie{
				isEnd:  true,
				shared: word[i:],
			}
		} else {
			t.children[idx].Insert(word[i:])
		}
	}
}

func (t *Trie) Search(word string) bool {
	if len(t.shared) > len(word) {
		return false
	}
	if t.shared == word[:len(t.shared)] {
		if len(t.shared) == len(word) {
			return t.isEnd
		}
		if t.children[word[len(t.shared)]-'a'] == nil {
			return false
		}
		return t.children[word[len(t.shared)]-'a'].Search(word[len(t.shared):])
	} else {
		return false
	}
}

func (t *Trie) StartsWith(prefix string) bool {
	i := 0
	for ; i < len(prefix) && i < len(t.shared) && prefix[i] == t.shared[i]; i++ {
	}
	if i == len(prefix) {
		return true
	}
	if i == len(t.shared) {
		idx := prefix[i] - 'a'
		if t.children[idx] == nil {
			return false
		}
		return t.children[idx].StartsWith(prefix[i:])
	}
	return false
}
