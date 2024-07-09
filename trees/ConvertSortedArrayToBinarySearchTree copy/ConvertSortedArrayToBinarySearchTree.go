package convertSortedArrayToBinarySearchTree

import (
	"github.com/AnkushinDaniil/leetcode-go/trees"
)

type TreeNode = trees.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(int, int) *TreeNode
	dfs = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		t := &TreeNode{}
		if l == r {
			t.Val = nums[l]
			return t
		}
		m := (l + r) / 2
		t.Val = nums[m]
		t.Left = dfs(l, m-1)
		t.Right = dfs(m+1, r)
		return t
	}
	return dfs(0, len(nums)-1)
}

func sortedArrayToBSTfast(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[0:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}
