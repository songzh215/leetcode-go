package leetcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fastPos := head
	slowPos := head
	for i := 0; i < n; i++ {
		if fastPos.Next != nil {
			fastPos = fastPos.Next
		} else {
			return slowPos.Next
		}
	}
	for fastPos.Next != nil {
		fastPos = fastPos.Next
		slowPos = slowPos.Next
	}
	slowPos.Next = slowPos.Next.Next
	return head
}
