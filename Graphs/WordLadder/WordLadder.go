package wordLadder

import "strings"

func ladderLengthSB(beginWord string, endWord string, wordList []string) int {
	isValid := false
	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			isValid = true
			break
		}
	}
	if !isValid {
		return 0
	}

	sb := &strings.Builder{}
	sb.Grow(len(wordList[0]))
	wordList = append(wordList, beginWord)
	neighbors := make(map[string][]string, len(wordList)*len(wordList[0]))
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < len(wordList[i]); j++ {
			sb.WriteString(wordList[i][:j])
			sb.WriteByte('.')
			sb.WriteString(wordList[i][j+1:])
			neighbors[sb.String()] = append(neighbors[sb.String()], wordList[i])
			sb.Reset()
		}
	}

	length := 1
	visited := make(map[string]struct{}, len(wordList))
	queue := make([]string, 0, len(wordList))

	visited[beginWord] = struct{}{}
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		for t := len(queue); t > 0; t-- {
			word := queue[0]
			queue = queue[1:]

			if word == endWord {
				return length
			}

			for i := 0; i < len(word); i++ {
				sb.WriteString(word[:i])
				sb.WriteByte('.')
				sb.WriteString(word[i+1:])
				for _, neighbor := range neighbors[sb.String()] {
					if _, ok := visited[neighbor]; !ok {
						visited[neighbor] = struct{}{}
						queue = append(queue, neighbor)
					}
				}
				sb.Reset()
			}
		}
		length++
	}

	return 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	isValid := false
	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			isValid = true
			break
		}
	}
	if !isValid {
		return 0
	}

	wordList = append(wordList, beginWord)
	neighbors := make(map[string][]string, len(wordList)*len(wordList[0]))
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < len(wordList[i]); j++ {
			pattern := wordList[i][:j] + "." + wordList[i][j+1:]
			neighbors[pattern] = append(neighbors[pattern], wordList[i])
		}
	}

	length := 1
	visited := make(map[string]struct{}, len(wordList))
	queue := make([]string, 0, len(wordList))

	visited[beginWord] = struct{}{}
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		for t := len(queue); t > 0; t-- {
			word := queue[0]
			queue = queue[1:]

			if word == endWord {
				return length
			}

			for i := 0; i < len(word); i++ {
				pattern := word[:i] + "." + word[i+1:]
				for _, neighbor := range neighbors[pattern] {
					if _, ok := visited[neighbor]; !ok {
						visited[neighbor] = struct{}{}
						queue = append(queue, neighbor)
					}
				}
			}
		}
		length++
	}

	return 0
}

func ladderLengthBytes(beginWord string, endWord string, wordList []string) int {
	isValid := false
	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			isValid = true
			break
		}
	}
	if !isValid {
		return 0
	}

	wordList = append(wordList, beginWord)
	neighbors := make(map[string][]string, len(wordList)*len(wordList[0]))
	pattern := make([]byte, len(wordList[0]))
	for i := 0; i < len(wordList); i++ {
		pattern = []byte(wordList[i])
		for j := 0; j < len(pattern); j++ {
			ch := pattern[j]
			pattern[j] = '.'
			neighbors[string(pattern)] = append(neighbors[string(pattern)], wordList[i])
			pattern[j] = ch
		}
	}

	length := 1
	visited := make(map[string]struct{}, len(wordList))
	queue := make([]string, 0, len(wordList))

	visited[beginWord] = struct{}{}
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		for t := len(queue); t > 0; t-- {
			word := queue[0]
			queue = queue[1:]

			if word == endWord {
				return length
			}

			for i := 0; i < len(word); i++ {
				pattern = []byte(word)
				for j := 0; j < len(pattern); j++ {
					ch := pattern[j]
					pattern[j] = '.'
					for _, neighbor := range neighbors[string(pattern)] {
						if _, ok := visited[neighbor]; !ok {
							visited[neighbor] = struct{}{}
							queue = append(queue, neighbor)
						}
					}
					pattern[j] = ch
				}
			}
		}
		length++
	}

	return 0
}
