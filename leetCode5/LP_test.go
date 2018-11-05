package leetcode

import (
	"testing"
)

func TestMaxStr(t *testing.T) {
	s := "babad"
	if longestPalindrome(s) != "bab" {
		t.Errorf("acutal is %s", longestPalindrome(s))
	}
}

func TestTwoMaxStr(t *testing.T) {
	s := "bb"
	if longestPalindrome(s) != "bb" {
		t.Errorf("acutal is %s", longestPalindrome(s))
	}
}
