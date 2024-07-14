package minStack

// type MinStack struct {
// 	stack, minStack []int
// }

// var (
// 	arr1 [30000]int
// 	arr2 [30000]int
// )

// func Constructor() MinStack {
// 	return MinStack{
// 		stack:    arr1[:0],
// 		minStack: arr2[:0],
// 	}
// }

// func (ms *MinStack) Push(val int) {
// 	ms.stack = append(ms.stack, val)
// 	if len(ms.minStack) == 0 {
// 		ms.minStack = append(ms.minStack, val)
// 	} else {
// 		ms.minStack = append(ms.minStack, min(ms.GetMin(), val))
// 	}
// }

// func (ms *MinStack) Pop() {
// 	ms.stack = ms.stack[:len(ms.stack)-1]
// 	ms.minStack = ms.minStack[:len(ms.minStack)-1]
// }

// func (ms *MinStack) Top() int {
// 	return ms.stack[len(ms.stack)-1]
// }

// func (ms *MinStack) GetMin() int {
// 	return ms.minStack[len(ms.minStack)-1]
// }

type Node struct {
	val, minimum int
}

type MinStack []Node

func Constructor() MinStack {
	return make(MinStack, 0, 16)
}

func (ms *MinStack) Push(val int) {
	node := Node{
		val:     val,
		minimum: val,
	}

	if len(*ms) > 0 {
		node.minimum = min(val, ms.GetMin())
	}

	*ms = append(*ms, node)
}

func (ms *MinStack) Pop() {
	*ms = (*ms)[:len(*ms)-1]
}

func (ms *MinStack) Top() int {
	return (*ms)[len(*ms)-1].val
}

func (ms *MinStack) GetMin() int {
	return (*ms)[len(*ms)-1].minimum
}
