package dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]Item, n)

	for i, temperature := range temperatures {
		for len(stack) > 0 && temperature > stack[len(stack)-1].Val {
			res[stack[len(stack)-1].Index] = i - stack[len(stack)-1].Index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Item{
			Val:   temperature,
			Index: i,
		})
	}

	return res
}

type Item struct {
	Val   int
	Index int
}
