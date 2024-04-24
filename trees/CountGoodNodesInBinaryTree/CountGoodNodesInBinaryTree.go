package countGoodNodesInBinaryTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func goodNodes(root *TreeNode) int {
	res := 0

	var rec func(*TreeNode, int)
	rec = func(tn *TreeNode, maximum int) {
		if tn == nil {
			return
		}
		if tn.Val >= maximum {
			maximum = tn.Val
			res++
		}

		rec(tn.Left, maximum)
		rec(tn.Right, maximum)
	}

	rec(root, root.Val)

	return res
}
