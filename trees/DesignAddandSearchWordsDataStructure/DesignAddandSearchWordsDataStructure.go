package designAddAndSearchWordsDataStructure

type WordDictionary struct {
	links [26]*WordDictionary
	end   bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

var index uint8

func (this *WordDictionary) AddWord(word string) {
	cur := this
	_ = word[len(word)-1]
	for i := 0; i < len(word); i++ {
		index = word[i] - 'a'
		if cur.links[index] == nil {
			cur.links[index] = &WordDictionary{}
		}
		cur = cur.links[index]
	}
	cur.end = true
}

func (this *WordDictionary) Search(word string) bool {
	_ = word[len(word)-1]
	var dfs func(wd *WordDictionary, i int) bool
	dfs = func(wd *WordDictionary, i int) bool {
		if wd == nil {
			return false
		}
		if i == len(word) {
			return wd.end
		}
		if word[i] == '.' {
			for j := 0; j < 26; j++ {
				if dfs(wd.links[j], i+1) {
					return true
				}
			}
			return false
		}
		return dfs(wd.links[word[i]-'a'], i+1)
	}

	return dfs(this, 0)
}
