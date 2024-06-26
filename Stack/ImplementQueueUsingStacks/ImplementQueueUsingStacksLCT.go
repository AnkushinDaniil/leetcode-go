package implementQueueUsingStacks

type MyQueueLCT struct {
	arr []int
}

func ConstructorLCT() MyQueueLCT {
	my_arr := make([]int, 0)
	return MyQueueLCT{
		arr: my_arr,
	}
}

func (this *MyQueueLCT) Push(x int) {
	this.arr = append([]int{x}, this.arr...)
}

func (this *MyQueueLCT) Pop() int {
	last_elem := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return last_elem
}

func (this *MyQueueLCT) Peek() int {
	return this.arr[len(this.arr)-1]
}

func (this *MyQueueLCT) Empty() bool {
	return len(this.arr) == 0
}
