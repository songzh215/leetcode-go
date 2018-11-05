package leetcode

import (
	"testing"
)

func TestMove0(t *testing.T) {
	s := []int{0, 1, 0, 3, 12}
	moveZeroes(s)
	if s[0] != 1 {
		t.Errorf("acutal is %d", s[0])
	}
	if s[1] != 3 {
		t.Errorf("acutal is %d", s[1])
	}
	if s[2] != 12 {
		t.Errorf("acutal is %d", s[2])
	}
	if s[3] != 0 {
		t.Errorf("acutal is %d", s[3])
	}
	if s[4] != 0 {
		t.Errorf("acutal is %d", s[4])
	}
}
