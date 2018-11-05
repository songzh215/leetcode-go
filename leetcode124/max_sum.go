package leetcode

import ()

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

var maxValue int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func calcMaxSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := max(calcMaxSum(root.Left), 0)
	rightSum := max(calcMaxSum(root.Right), 0)

	maxValue = max(maxValue, leftSum+rightSum+root.Val)
	return max(0, max(leftSum, rightSum)+root.Val)
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue = INT_MIN
	calcMaxSum(root)
	return maxValue
}
