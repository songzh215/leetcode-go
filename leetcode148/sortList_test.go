package leetcode

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	n1 := &ListNode{Val: 4}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	res := sortList(n1)
	if res.Val != 1 {
		t.Errorf("sort list %d", res.Val)
	}
	if res.Next.Val != 2 {
		t.Errorf("sort list %d", res.Next.Val)
	}
	if res.Next.Next.Val != 3 {
		t.Errorf("sort list %d", res.Next.Next.Val)
	}
	if res.Next.Next.Next.Val != 4 {
		t.Errorf("sort list %d", res.Next.Next.Next.Val)
	}
}

func TestSortListLong(t *testing.T) {
	n1 := &ListNode{Val: 8}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n5 := &ListNode{Val: 10}
	n6 := &ListNode{Val: 5}
	n7 := &ListNode{Val: 4}
	n8 := &ListNode{Val: 11}
	n9 := &ListNode{Val: 7}
	n10 := &ListNode{Val: 5}
	n11 := &ListNode{Val: 9}
	n12 := &ListNode{Val: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	n7.Next = n8
	n8.Next = n9
	n9.Next = n10
	n10.Next = n11
	n11.Next = n12
	res := sortList(n1)
	for res != nil {
		fmt.Printf("val %d \n", res.Val)
		res = res.Next
	}
}
