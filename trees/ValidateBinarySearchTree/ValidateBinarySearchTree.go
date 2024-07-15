package validateBinarySearchTree

import (
	"math"

	"github.com/AnkushinDaniil/leetcode-go/trees"
)

type TreeNode = trees.TreeNode

func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(tn *TreeNode, l, r int) bool {
		if tn == nil {
			return true
		}
		if !(tn.Val > l && tn.Val < r) {
			return false
		}
		return dfs(tn.Left, l, tn.Val) && dfs(tn.Right, tn.Val, r)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	var rec func(*TreeNode, *int, *int) bool
// 	rec = func(tn *TreeNode, minimum, maximum *int) bool {
// 		if tn == nil {
// 			return true
// 		} else if minimum != nil && tn.Val <= *minimum || maximum != nil && tn.Val >= *maximum {
// 			return false
// 		} else {
// 			return rec(tn.Left, minimum, &tn.Val) && rec(tn.Right, &tn.Val, maximum)
// 		}
// 	}

// 	return rec(root, nil, nil)
// }
