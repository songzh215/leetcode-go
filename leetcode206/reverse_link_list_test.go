package leetcode

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n1.Next = &n2
	n3 := ListNode{Val: 3}
	n2.Next = &n3
	n4 := ListNode{Val: 4}
	n3.Next = &n4
	n5 := ListNode{Val: 5}
	n4.Next = &n5
	nr := reverseList(&n1)
	if nr.Val != 5 {
		t.Errorf("r n1 is %d ,expect 5", nr.Val)
	}
	if nr.Next.Val != 4 {
		t.Errorf("r n1 is %d ,expect 4", nr.Next.Val)
	}
	if nr.Next.Next.Val != 3 {
		t.Errorf("r n1 is %d ,expect 3", nr.Next.Next.Val)
	}
	if nr.Next.Next.Next.Val != 2 {
		t.Errorf("r n1 is %d ,expect 2", nr.Next.Next.Next.Val)
	}
	if nr.Next.Next.Next.Next.Val != 1 {
		t.Errorf("r n1 is %d ,expect 1", nr.Next.Next.Next.Next.Val)
	}
}

func TestReverseListTwoElements(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n1.Next = &n2
	nr := reverseList(&n1)
	if nr.Val != 2 {
		t.Errorf("r n1 is %d ,expect 2", nr.Val)
	}
	if nr.Next.Val != 1 {
		t.Errorf("r n1 is %d ,expect 1", nr.Next.Val)
	}

}
