package leetcode

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil || left != nil && right == nil {
		return false
	}
	if left != nil && right != nil && left.Val != right.Val {
		return false
	}
	if !isSymmetricHelper(left.Right, right.Left) {
		return false
	}
	if !isSymmetricHelper(left.Left, right.Right) {
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}
