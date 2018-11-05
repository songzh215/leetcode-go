package leetcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	other_half := reverse(slow.Next)
	return compare(head, other_half)
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var last *ListNode = nil
	cur := head
	next := cur.Next
	for cur.Next != nil {
		next = cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	cur.Next = last
	return cur
}

func compare(l1, l2 *ListNode) bool {
	for l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
