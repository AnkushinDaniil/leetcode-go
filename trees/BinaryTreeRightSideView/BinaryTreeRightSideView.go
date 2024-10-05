package binaryTreeRightSideView

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	level := 0

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if level == len(res) {
			res = append(res, r.Val)
		}
		level++
		dfs(r.Right)
		dfs(r.Left)
		level--
	}
	dfs(root)

	return res
}
