package binaryTreeLevelOrderTraversal

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	var rec func(*TreeNode, int)
	rec = func(tn *TreeNode, level int) {
		if tn == nil {
			return
		}
		if len(res) < level+1 {
			res = append(res, []int{tn.Val})
		} else {
			res[level] = append(res[level], tn.Val)
		}

		rec(tn.Left, level+1)
		rec(tn.Right, level+1)
	}

	rec(root, 0)

	return res
}
