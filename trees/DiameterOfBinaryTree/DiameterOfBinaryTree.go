package diameterOfBinaryTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var diam func(*TreeNode) int

	diam = func(t *TreeNode) int {
		if t == nil {
			return -1
		}

		l := diam(t.Left)
		r := diam(t.Right)

		res = max(res, 2+l+r)

		return 1 + max(l, r)
	}

	diam(root)

	return res
}
