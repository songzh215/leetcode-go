package leetcode

import (
	"testing"
)

func TestDiameter(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Left = n2
	n1.Right = n3
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n2.Left = n4
	n2.Right = n5
	if diameterOfBinaryTree(n1) != 3 {
		t.Errorf("diameterOfBinaryTree is %d", diameterOfBinaryTree(n1))
	}
}

func TestDiameterEmpty(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	if diameterOfBinaryTree(n1) != 0 {
		t.Errorf("diameterOfBinaryTree is %d", diameterOfBinaryTree(n1))
	}
}

func TestDiameterCase3(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n5.Left = n2
	n2.Left = n4
	n2.Right = n3
	n4.Left = n1
	if diameterOfBinaryTree(n5) != 3 {
		t.Errorf("diameterOfBinaryTree is %d", diameterOfBinaryTree(n5))
	}
}
