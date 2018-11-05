package leetcode

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	v := findKthLargest(nums, 2)
	if v != 5 {
		t.Errorf("2 largest not 5:", v)
	}
}
func TestKthLargestcase1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	v := findKthLargest(nums, 2)
	if v != 5 {
		t.Errorf("2 largest not 5:", v)
	}
}
