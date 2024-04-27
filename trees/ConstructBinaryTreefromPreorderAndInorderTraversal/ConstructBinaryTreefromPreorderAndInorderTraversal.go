package constructBinaryTreefromPreorderAndInorderTraversal

import (
	"github.com/AnkushinDaniil/leetcode-go/trees"
)

type TreeNode = trees.TreeNode

func buildTree(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}

	m := 0

	for i, v := range inOrder {
		if v == preOrder[0] {
			m = i
			break
		}
	}

	return &TreeNode{
		Val:   preOrder[0],
		Left:  buildTree(preOrder[1:m+1], inOrder[:m]),
		Right: buildTree(preOrder[m+1:], inOrder[m+1:]),
	}
}
