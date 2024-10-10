package sortanarray

type heap []int

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return i*2 + 1 }
func right(i int) int  { return i*2 + 2 }

func (h heap) cmp(i, j int) bool {
	return h[i] > h[j]
}

func (h heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// func (h heap) heapifyUp(i int) {
// 	for p := parent(i); i > 0 && h.cmp(h[i], h[p]); i = p {
// 		h.swap(i, p)
// 	}
// }

func (h heap) heapifyDown(i int) {
	for {
		hightest := i
		l := left(i)
		r := right(i)

		if l < len(h) && h.cmp(l, hightest) {
			hightest = l
		}

		if r < len(h) && h.cmp(r, hightest) {
			hightest = r
		}

		if hightest == i {
			break
		} else {
			h.swap(hightest, i)
			i = hightest
		}
	}
}

func (h heap) heapify() {
	for i := parent(len(h)); i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// func (h *heap) insert(item int) {
//   h = append(h, item)
//   h.heapifyUp(len(h)-1)
// }

func (h *heap) pop() int {
	if len(*h) == 0 {
		return 0
	}
	res := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	h.heapifyDown(0)
	return res
}

func sortArray(nums []int) []int {
	h := heap(nums)
	h.heapify()
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = h.pop()
	}
	return nums
}
