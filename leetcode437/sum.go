package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cnt = 0

func sumOk(root *TreeNode, target int) {
	if root == nil {
		return
	}
	if root.Val == target {
		cnt++
	}
	sumOk(root.Left, target-root.Val)
	sumOk(root.Right, target-root.Val)
}

func dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sumOk(root, sum)
	dfs(root.Right, sum)
	dfs(root.Left, sum)
}

func pathSum(root *TreeNode, sum int) int {
	dfs(root, sum)
	return cnt
}
