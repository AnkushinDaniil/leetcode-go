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
		t := false
		if tn.Val == l || tn.Val == r {
			t = true
		}
		l := dfs(tn.Left)
		r := dfs(tn.Right)
		if (t && l) || (t && r) || (l && r) {
			res = tn
		}
		return t || l || r
	}
	dfs(root)
	return res
}
