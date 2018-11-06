package leetcode

import (
	"testing"
)

func TestSum(t *testing.T) {
	n := TreeNode{Val: 10}
	n1 := TreeNode{Val: 5}
	n2 := TreeNode{Val: -3}
	n.Left = &n1
	n.Right = &n2
	n3 := TreeNode{Val: 3}
	n4 := TreeNode{Val: 2}
	n5 := TreeNode{Val: 11}
	n1.Left = &n3
	n1.Right = &n4
	n2.Right = &n5
	n6 := TreeNode{Val: 3}
	n7 := TreeNode{Val: -2}
	n3.Left = &n6
	n3.Right = &n7
	n8 := TreeNode{Val: 1}
	n4.Right = &n8
	res := pathSum(&n, 8)
	if res != 3 {
		t.Errorf("path sum is %d:", res)
	}
}
