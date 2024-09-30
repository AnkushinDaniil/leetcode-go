package minimumremovetomakevalidparentheses

const (
	opening byte = '('
	closing byte = ')'
)

type parenthesisWithIndex struct {
	parenthesis byte
	index       int
}

func minRemoveToMakeValid(s string) string {
	stack := findExtraParenthesis(s)
	return deleteFromString(s, stack)
}

func findExtraParenthesis(s string) []parenthesisWithIndex {
	stack := make([]parenthesisWithIndex, 0, len(s))
	for i := range s {
		switch s[i] {
		case opening:
			stack = append(stack, parenthesisWithIndex{parenthesis: opening, index: i})
		case closing:
			if len(stack) > 0 && stack[len(stack)-1].parenthesis == opening {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, parenthesisWithIndex{parenthesis: closing, index: i})
			}
		}
	}
	return stack
}

func deleteFromString(s string, stack []parenthesisWithIndex) string {
	bts := make([]byte, 0, len(s))
	stackIndex := 0
	for i := range s {
		if stackIndex < len(stack) && stack[stackIndex].index == i {
			stackIndex++
		} else {
			bts = append(bts, s[i])
		}
	}
	return string(bts)
}
