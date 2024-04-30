package implementTrie

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

var (
	index uint8
	i     int
)

func (this *Trie) Insert(word string) {
	cur := this
	_ = word[len(word)-1]
	for i = 0; i < len(word); i++ {
		index = word[i] - 'a'
		if cur.links[index] == nil {
			cur.links[index] = &Trie{}
		}
		cur = cur.links[index]
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	_ = word[len(word)-1]
	for i = 0; i < len(word); i++ {
		index = word[i] - 'a'
		if cur.links[index] == nil {
			return false
		}
		cur = cur.links[index]
	}
	return cur.end
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	_ = prefix[len(prefix)-1]
	for i = 0; i < len(prefix); i++ {
		index = prefix[i] - 'a'
		if cur.links[index] == nil {
			return false
		}
		cur = cur.links[index]
	}
	return true
}
