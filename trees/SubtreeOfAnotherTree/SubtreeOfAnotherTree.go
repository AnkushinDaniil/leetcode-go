package subtreeOfAnotherTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pIsNil := p == nil
	qIsNil := q == nil

	if pIsNil && qIsNil {
		return true
	}
	if pIsNil != qIsNil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
