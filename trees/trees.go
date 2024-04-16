package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeFromArray(arr []interface{}) *TreeNode {
	res := make([]TreeNode, len(arr))
	for i := 0; i < len(arr); i++ {
		if val, ok := arr[i].(int); ok {
			res[i].Val = val
		}
		if j := 2*i + 1; j < len(arr) {
			if _, ok := arr[j].(int); ok {
				res[i].Left = &res[j]
			}
		}
		if j := 2*i + 2; j < len(arr) {
			if _, ok := arr[j].(int); ok {
				res[i].Right = &res[j]
			}
		}
	}
	return &res[0]
}
