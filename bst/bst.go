package bst

import ()

/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowerThan(root *TreeNode, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max {
		return false
	}
	if root.Right != nil && !lowerThan(root.Right, max) {
		return false
	}
	return true
}

func largerThan(root *TreeNode, min int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Left != nil && !largerThan(root.Left, min) {
		return false
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && !lowerThan(root.Left, root.Val) {
		return false
	}
	if root.Right != nil && !largerThan(root.Right, root.Val) {
		return false
	}
	if root.Left != nil && !isValidBST(root.Left) {
		return false
	}
	if root.Right != nil && !isValidBST(root.Right) {
		return false
	}
	return true
}
