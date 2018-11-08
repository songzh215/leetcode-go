package leetcode

import (
	"testing"
)

func TestSubtree(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n3.Left = n4
	n3.Right = n5
	n4.Left = n1
	n4.Right = n2

	n11 := &TreeNode{Val: 1}
	n22 := &TreeNode{Val: 2}
	n44 := &TreeNode{Val: 4}
	n44.Left = n11
	n44.Right = n22

	if !isSubtree(n3, n44) {
		t.Errorf("is not subtree")
	}
}
