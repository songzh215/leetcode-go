package leetcode

import (
	"testing"
)

func TestLongestSeq(t *testing.T) {
	var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	if lengthOfLIS(nums) != 4 {
		t.Errorf("lengthOfLIS is %d", lengthOfLIS(nums))
	}
}

func TestLongestSeqCase1(t *testing.T) {
	var nums = []int{4, 10, 4, 3, 8, 9}
	if lengthOfLIS(nums) != 3 {
		t.Errorf("lengthOfLIS is %d", lengthOfLIS(nums))
	}
}
