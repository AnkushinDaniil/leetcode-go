package validateBinarySearchTree

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var rec func(*TreeNode, *int, *int) bool
	rec = func(tn *TreeNode, minimum, maximum *int) bool {
		if tn == nil {
			return true
		} else if minimum != nil && tn.Val <= *minimum || maximum != nil && tn.Val >= *maximum {
			return false
		} else {
			return rec(tn.Left, minimum, &tn.Val) && rec(tn.Right, &tn.Val, maximum)
		}
	}

	return rec(root, nil, nil)
}
