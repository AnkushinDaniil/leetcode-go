package symmetricTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left, root.Right)
}

func dfs(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return (l == nil) == (r == nil)
	}
	return l.Val == r.Val && dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
}
