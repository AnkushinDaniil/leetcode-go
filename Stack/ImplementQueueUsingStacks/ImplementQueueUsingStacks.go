package implementQueueUsingStacks

const INITIAL_STACK_CAPACITY = 2

type Stack []int

func NewStack() *Stack {
	stack := make(Stack, 0, INITIAL_STACK_CAPACITY)
	return &stack
}

func (s *Stack) Append(x int) {
	*s = append(*s, x)
}

func (s Stack) Peek() int {
	return s[len(s)-1]
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

type MyQueue [2]Stack

func Constructor() MyQueue {
	return MyQueue{*NewStack(), *NewStack()}
}

func (q *MyQueue) Push(x int) {
	q[0].Append(x)
}

func (q *MyQueue) reallocate() {
	if q[1].Empty() {
		for !q[0].Empty() {
			q[1].Append(q[0].Pop())
		}
	}
}

func (q *MyQueue) Pop() int {
	q.reallocate()
	return q[1].Pop()
}

func (q *MyQueue) Peek() int {
	q.reallocate()
	return q[1].Peek()
}

func (q *MyQueue) Empty() bool {
	return q[0].Empty() && q[1].Empty()
}
