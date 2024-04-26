package kthSmallestElementInBST

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0, 64)
	n := -1
	cur := root

	for {
		for cur != nil {
			stack = append(stack, cur)
			n++
			cur = cur.Left
		}
		cur = stack[n]
		stack = stack[:n]
		n--
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
}
