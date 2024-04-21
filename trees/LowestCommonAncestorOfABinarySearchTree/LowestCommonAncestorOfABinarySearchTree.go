package lowestCommonAncestorOfABinarySearchTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	l, r := p.Val, q.Val
	if l > r {
		l, r = r, l
	}
	var rec func(*TreeNode) *TreeNode
	rec = func(tn *TreeNode) *TreeNode {
		if tn.Val >= l && tn.Val <= r {
			return tn
		} else if tn.Val > r {
			return rec(tn.Left)
		} else {
			return rec(tn.Right)
		}
	}
	return rec(root)
}
