package binaryTreeMaximumPathSum

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func maxPathSum(root *TreeNode) int {
	res := root.Val

	var rec func(*TreeNode) int
	rec = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}

		l := rec(tn.Left)
		if l < 0 {
			l = 0
		}

		r := rec(tn.Right)
		if r < 0 {
			r = 0
		}

		if t := l + r + tn.Val; t > res {
			res = t
		}

		return max(l, r) + tn.Val
	}

	rec(root)

	return res
}

func maxPathSum_(root *TreeNode) int {
	res := root.Val

	var rec func(*TreeNode) int
	rec = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}

		l := max(0, rec(tn.Left))
		r := max(0, rec(tn.Right))

		res = max(res, l+r+tn.Val)

		return max(l, r) + tn.Val
	}

	rec(root)

	return res
}
