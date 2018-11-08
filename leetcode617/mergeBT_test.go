package leetcode

import (
	"testing"
)

func TestMerge1(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Left = n3
	n1.Right = n2
	n5 := &TreeNode{Val: 5}
	n3.Left = n5

	n11 := &TreeNode{Val: 1}
	n22 := &TreeNode{Val: 2}
	n33 := &TreeNode{Val: 3}
	n44 := &TreeNode{Val: 4}
	n77 := &TreeNode{Val: 7}
	n22.Left = n11
	n22.Right = n33
	n11.Right = n44
	n33.Right = n77

	mergeNode := mergeTrees(n1, n22)
	if mergeNode.Val != 3 {
		t.Errorf("merge tree root %d", mergeNode.Val)
	}
	if mergeNode.Left.Val != 4 {
		t.Errorf("merge tree left %d", mergeNode.Left.Val)
	}
	if mergeNode.Right.Val != 5 {
		t.Errorf("merge tree right %d", mergeNode.Right.Val)
	}
	if mergeNode.Left.Left.Val != 5 {
		t.Errorf("merge tree left.left %d", mergeNode.Left.Left.Val)
	}
	if mergeNode.Left.Right.Val != 4 {
		t.Errorf("merge tree left.left %d", mergeNode.Left.Right.Val)
	}
	if mergeNode.Right.Right.Val != 7 {
		t.Errorf("merge tree right.right %d", mergeNode.Right.Right.Val)
	}
}
