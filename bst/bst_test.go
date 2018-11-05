package bst

import (
	"testing"
)

func TestIsBST(t *testing.T) {
	root := TreeNode{Val: 2}
	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 3}
	root.Left = &n1
	root.Right = &n2
	if !isValidBST(&root) {
		t.Errorf("is not bst")
	}
}
