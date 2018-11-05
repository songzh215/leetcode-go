package leetcode

import (
	"testing"
)

func TestSearch(t *testing.T) {
	s := []int{0, 1, 2}
	if search(s, 0) != 0 {
		t.Errorf("acutal is %d", search(s, 0))
	}
}

func TestSearch1(t *testing.T) {
	s := []int{4, 7, 8, 9, 11, 0, 1, 2}
	if search(s, 8) != 2 {
		t.Errorf("acutal is %d", search(s, 8))
	}
}
