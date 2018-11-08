package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isTheSameTree(s *TreeNode, t *TreeNode) bool {
	if t == nil && s == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}
	if s != nil && t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}
	if !isTheSameTree(s.Left, t.Left) {
		return false
	}
	if !isTheSameTree(s.Right, t.Right) {
		return false
	}
	return true
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isTheSameTree(s, t) {
		return true
	}

	if s != nil {
		if isSubtree(s.Left, t) {
			return true
		}
		if isSubtree(s.Right, t) {
			return true
		}
	}
	return false
}
