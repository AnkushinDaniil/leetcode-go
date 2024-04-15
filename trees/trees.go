package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeFromArray(arr []int) *TreeNode {
	res := make([]TreeNode, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i].Val = arr[i]
		if j := 2*i + 1; j < len(arr) {
			res[i].Left = &res[j]
		}
		if j := 2*i + 2; j < len(arr) {
			res[i].Right = &res[j]
		}
	}
	return &res[0]
}
