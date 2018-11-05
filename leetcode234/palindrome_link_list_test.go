package leetcode

import (
	"testing"
)

func TestIsPalindromeLink(t *testing.T) {
	// 6 3 5 5 3 6
	n1 := &ListNode{Val: 6}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 5}
	n5 := &ListNode{Val: 3}
	n6 := &ListNode{Val: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	if isPalindrome(n1) != true {
		t.Errorf("6 3 5 5 3 6 shold be Palindrome")
	}
}

func TestIsPalindromeLinkFalseCase(t *testing.T) {
	// 6 3 5 5 3 6
	n1 := &ListNode{Val: 6}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 3}
	n5 := &ListNode{Val: 3}
	n6 := &ListNode{Val: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	if isPalindrome(n1) != false {
		t.Errorf("6 3 5 5 3 7 shold not be Palindrome")
	}
}

func TestIsPalindromeLinkFalseCase1(t *testing.T) {
	// 6 3 5 5 3 6
	n1 := &ListNode{Val: 6}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 3}
	n5 := &ListNode{Val: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	if isPalindrome(n1) != true {
		t.Errorf("6 3 5 5 3 7 shold not be Palindrome")
	}
}

func TestIsPalindromeLinkOneElement(t *testing.T) {
	// 6
	n1 := &ListNode{Val: 6}
	if isPalindrome(n1) != true {
		t.Errorf("6 shold not be Palindrome")
	}
}
