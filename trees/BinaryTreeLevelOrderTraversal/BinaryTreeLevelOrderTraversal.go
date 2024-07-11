package binaryTreeLevelOrderTraversal

import "github.com/AnkushinDaniil/leetcode-go/trees"

type TreeNode = trees.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0, 4)
	queue := make([]*TreeNode, 0, 4)
	queue = append(queue, root)

	for level := 0; len(queue) > 0; level++ {
		res = append(res, make([]int, 0, len(queue)))
		l := len(queue)
		for range l {
			res[level] = append(res[level], queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}

	return res
}
