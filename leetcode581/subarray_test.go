package leetcode

import (
	"testing"
)

func TestSubarray(t *testing.T) {
	var nums = []int{2, 6, 4, 8, 10, 9, 15}
	if findUnsortedSubarray(nums) != 5 {
		t.Errorf("subarray is %d", findUnsortedSubarray(nums))
	}
}
