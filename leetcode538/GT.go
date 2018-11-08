package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumUp(node *TreeNode) int {
	if node == nil {
		return 0
	}
	res := node.Val
	if node.Right != nil {
		res += sumUp(node.Right)
	}
	if node.Left != nil {
		res += sumUp(node.Left)
	}
	return res
}

func calcNode(base int, root *TreeNode) {
	if root == nil {
		return
	}
	root.Val = root.Val + base
	if root.Right != nil {
		root.Val += sumUp(root.Right)
	}
	calcNode(base, root.Right)
	calcNode(root.Val, root.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	calcNode(0, root)
	return root
}
