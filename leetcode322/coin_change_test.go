package leetcode

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	var nums = []int{1, 2, 5}
	amount := 11
	if coinChange(nums, amount) != 3 {
		t.Errorf("lengthOfLIS is %d", coinChange(nums, amount))
	}
}

func TestCoinChangeCase1(t *testing.T) {
	var nums = []int{2}
	amount := 3
	if coinChange(nums, amount) != -1 {
		t.Errorf("lengthOfLIS is %d", coinChange(nums, amount))
	}
}
