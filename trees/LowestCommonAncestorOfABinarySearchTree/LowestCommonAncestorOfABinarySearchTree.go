package lowestCommonAncestorOfABinarySearchTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	l, r := p.Val, q.Val
	var res *TreeNode

	var dfs func(*TreeNode) bool
	dfs = func(tn *TreeNode) bool {
		if res != nil {
			return false
		}
		if tn == nil {
			return false
		}
		if tn.Val == l || tn.Val == r {
			return true
		}
		if dfs(tn.Right) && dfs(tn.Left) {
			res = tn
		}
		return false
	}
	dfs(root)
	return res
}
