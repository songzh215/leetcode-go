package leetcode

import (
	"testing"
)

func TestRobber(t *testing.T) {
	nums := make([]int, 4)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 1
	if rob(nums) != 4 {
		t.Errorf("rob num should be 4, actual %d", rob(nums))
	}
}
