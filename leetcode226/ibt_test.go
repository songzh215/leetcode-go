package leetcode

import (
	"testing"
)

func TestDiameter(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}
	n9 := &TreeNode{Val: 9}
	n4.Left = n2
	n4.Right = n7
	n2.Left = n1
	n2.Right = n3
	n7.Left = n6
	n7.Right = n9
	root := invertTree(n4)
	//		t.Errorf("diameterOfBinaryTree is %d", diameterOfBinaryTree(n1))
	if root.Val != 4 {
		t.Errorf("root is %d", root.Val)
	}
	if root.Left.Val != 7 {
		t.Errorf("root is %d", root.Left.Val)
	}
	if root.Right.Val != 2 {
		t.Errorf("root is %d", root.Right.Val)
	}
	if root.Left.Left.Val != 9 {
		t.Errorf("root is %d", root.Left.Left.Val)
	}
	if root.Left.Right.Val != 6 {
		t.Errorf("root is %d", root.Left.Right.Val)
	}
	if root.Right.Right.Val != 1 {
		t.Errorf("root is %d", root.Right.Right.Val)
	}
	if root.Right.Left.Val != 3 {
		t.Errorf("root is %d", root.Right.Left.Val)
	}
}
