package minStack

type Value struct {
	minimum int
	value   int
}

type MinStack struct {
	stack []Value
}

func Constructor() MinStack {
	return MinStack{
		stack: []Value{},
	}
}

func (s *MinStack) Push(val int) {
	value := Value{
		minimum: val,
		value:   val,
	}

	if len(s.stack) == 0 {
		s.stack = append(s.stack, value)
		return
	}

	value.minimum = min(val, s.stack[len(s.stack)-1].minimum)
	s.stack = append(s.stack, value)
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].value
}

func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].minimum
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
