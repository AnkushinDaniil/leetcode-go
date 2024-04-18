package balancedBinaryTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func isBalanced(root *TreeNode) bool {
	res, _ := rec(root)
	return res
}

func rec(t *TreeNode) (bool, int) {
	if t == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := rec(t.Left)
	isRightBalanced, rightHeight := rec(t.Right)

	return isLeftBalanced &&
			isRightBalanced &&
			abs(leftHeight-rightHeight) <= 1,
		max(
			leftHeight,
			rightHeight,
		) + 1
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}
