package leetcode

import ()

type ListNode struct {
	Val  int
	Next *ListNode
}

// merge sort
func merge(head *ListNode, mid *ListNode) *ListNode {
	if head == nil || mid == nil {
		return head
	}
	l1 := head
	l2 := mid
	newP := &ListNode{}
	res := newP

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newP.Next = l1
			l1 = l1.Next
		} else {
			newP.Next = l2
			l2 = l2.Next
		}
		newP = newP.Next
	}
	if l1 == nil {
		newP.Next = l2
	} else {
		newP.Next = l1
	}
	return res.Next
}

func sort(head *ListNode) *ListNode {
	sp := head
	qp := head
	if head == nil || head.Next == nil {
		return head
	}
	pre := sp
	for {
		pre = sp
		sp = sp.Next
		if qp.Next == nil || qp.Next.Next == nil {
			break
		}
		qp = qp.Next.Next
	}
	pre.Next = nil
	res := merge(sort(head), sort(sp))
	return res
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := sort(head)
	return res
}
