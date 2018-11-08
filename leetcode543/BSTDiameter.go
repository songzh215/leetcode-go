package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxHeight(root.Left)
	r := maxHeight(root.Right)
	res := l
	if r > l {
		res = r
	}
	return res + 1
}

func path(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left int
	var right int
	right = maxHeight(root.Right)
	left = maxHeight(root.Left)
	return left + right
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := path(root)
	r := diameterOfBinaryTree(root.Right)
	l := diameterOfBinaryTree(root.Left)
	if r > res {
		res = r
	}
	if l > res {
		res = l
	}
	return res
}
