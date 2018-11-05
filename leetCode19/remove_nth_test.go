package leetcode

import (
	"testing"
)

func TestRemoveNth(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 3}
	n4 := ListNode{Val: 4}
	n5 := ListNode{Val: 5}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	nx := removeNthFromEnd(&n1, 2)
	if nx.Val != 1 {
		t.Errorf("n1 acutal is %d", nx.Val)
	}
	if nx.Next.Val != 2 {
		t.Errorf("n2 acutal is %d", nx.Next.Val)
	}
	if nx.Next.Next.Val != 3 {
		t.Errorf("n3 acutal is %d", nx.Next.Next.Val)
	}
	if nx.Next.Next.Next.Val != 5 {
		t.Errorf("n5 acutal is %d", nx.Next.Next.Next.Val)
	}
	/*if nx.Next.Next.Next.Next.Val != 1 {
		t.Errorf("n1 acutal is %d", nx.Next.Next.Next.Next.Val)
	}*/
}

func TestRemovelastOne(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n1.Next = &n2
	nx := removeNthFromEnd(&n1, 1)
	if nx.Val != 1 {
		t.Errorf("n1 acutal is %d", nx.Val)
	}
	if nx.Next != nil {
		t.Errorf("n2 acutal is %d", nx.Next.Val)
	}
	/*if nx.Next.Next.Next.Next.Val != 1 {
		t.Errorf("n1 acutal is %d", nx.Next.Next.Next.Next.Val)
	}*/
}

func TestRemoveOne(t *testing.T) {
	n1 := ListNode{Val: 1}
	nx := removeNthFromEnd(&n1, 1)
	if nx != nil {
		t.Errorf("n1 acutal is %d", nx.Val)
	}
}

func TestRemovefirstOne(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n1.Next = &n2
	nx := removeNthFromEnd(&n1, 2)
	if nx.Val != 2 {
		t.Errorf("n1 acutal is %d", nx.Val)
	}
	if nx.Next != nil {
		t.Errorf("n2 acutal is %d", nx.Next.Val)
	}
	/*if nx.Next.Next.Next.Next.Val != 1 {
		t.Errorf("n1 acutal is %d", nx.Next.Next.Next.Next.Val)
	}*/
}
