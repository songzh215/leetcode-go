package leetcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var last *ListNode = nil
	next := head.Next
	for cur.Next != nil {
		next = cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	cur.Next = last
	return cur
}
