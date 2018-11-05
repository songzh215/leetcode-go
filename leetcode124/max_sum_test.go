package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	n1 := ListNode{Val: 2}
	n2 := ListNode{Val: 4}
	n3 := ListNode{Val: 3}
	n1.Next = &n2
	n2.Next = &n3

	n4 := ListNode{Val: 5}
	n5 := ListNode{Val: 6}
	n6 := ListNode{Val: 4}
	n4.Next = &n5
	n5.Next = &n6
	node := addTwoNumbers(&n1, &n4)
	if node.Val != 7 {
		t.Errorf("addTwoNumbers expect 7, actually: [%d]", node.Val)
	}
	if node.Next.Val != 0 {
		t.Errorf("addTwoNumbers expect 0, actually: [%d]", node.Next.Val)
	}
	if node.Next.Next.Val != 8 {
		t.Errorf("addTwoNumbers expect 8, actually: [%d]", node.Next.Next.Val)
	}
}

func TestAddTwoNumbersOvers(t *testing.T) {
	n1 := ListNode{Val: 8}
	n2 := ListNode{Val: 3}
	n3 := ListNode{Val: 2}
	n4 := ListNode{Val: 9}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4

	n5 := ListNode{Val: 4}
	n6 := ListNode{Val: 7}
	n7 := ListNode{Val: 9}
	n5.Next = &n6
	n6.Next = &n7
	node := addTwoNumbers(&n1, &n5)
	if node.Val != 2 {
		t.Errorf("addTwoNumbers expect 2, actually: [%d]", node.Val)
	}
	if node.Next.Val != 1 {
		t.Errorf("addTwoNumbers expect 1, actually: [%d]", node.Next.Val)
	}
	if node.Next.Next.Val != 2 {
		t.Errorf("addTwoNumbers expect 2, actually: [%d]", node.Next.Next.Val)
	}
	if node.Next.Next.Next.Val != 0 {
		t.Errorf("addTwoNumbers expect 0, actually: [%d]", node.Next.Next.Next.Val)
	}
	if node.Next.Next.Next.Next.Val != 1 {
		t.Errorf("addTwoNumbers expect 1, actually: [%d]", node.Next.Next.Next.Next.Val)
	}
}

func TestAddTwoNumbers5_5(t *testing.T) {
	n1 := ListNode{Val: 5}

	n2 := ListNode{Val: 5}
	node := addTwoNumbers(&n1, &n2)
	if node.Val != 0 {
		t.Errorf("addTwoNumbers expect 0, actually: [%d]", node.Val)
	}
	if node.Next.Val != 1 {
		t.Errorf("addTwoNumbers expect 1, actually: [%d]", node.Next.Val)
	}
}

func TestAddTwoNumbersSpec(t *testing.T) {
	n1 := ListNode{Val: 9}
	n2 := ListNode{Val: 1}
	n3 := ListNode{Val: 6}
	n1.Next = &n2
	n2.Next = &n3

	n4 := ListNode{Val: 0}

	node := addTwoNumbers(&n1, &n4)
	if node.Val != 9 {
		t.Errorf("addTwoNumbers expect 9, actually: [%d]", node.Val)
	}
	if node.Next.Val != 1 {
		t.Errorf("addTwoNumbers expect 1, actually: [%d]", node.Next.Val)
	}
	if node.Next.Next.Val != 6 {
		t.Errorf("addTwoNumbers expect 6, actually: [%d]", node.Next.Next.Val)
	}
}

/*
func main() {
}*/
