package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	var res ListNode
	p_res := &res
	add := 0
	for ; p1 != nil && p2 != nil; p1 = p1.Next {
		tmp := p1.Val + p2.Val + add
		add = tmp / 10
		newNode := ListNode{Val: tmp % 10}
		p_res.Next = &newNode
		p_res = p_res.Next
		p2 = p2.Next
	}

	if p1 == nil && p2 != nil {
		newNode := ListNode{Val: add}
		p_res.Next = addTwoNumbers(&newNode, p2)
		for ; p_res.Next != nil; p_res = p_res.Next {
		}
	}
	if p2 == nil && p1 != nil {
		newNode := ListNode{Val: add}
		p_res.Next = addTwoNumbers(&newNode, p1)
		for ; p_res.Next != nil; p_res = p_res.Next {
		}
	}
	if p2 == nil && p1 == nil && add != 0 {
		newNode := ListNode{Val: add}
		p_res.Next = &newNode
		p_res = p_res.Next
	}
	return res.Next
}
