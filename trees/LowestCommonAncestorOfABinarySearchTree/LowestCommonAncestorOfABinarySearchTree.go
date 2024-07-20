package lowestCommonAncestorOfABinarySearchTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode

	var dfs func(*TreeNode) bool
	dfs = func(tn *TreeNode) bool {
		if res != nil || tn == nil {
			return false
		}

		t := tn.Val == p.Val || tn.Val == q.Val
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
