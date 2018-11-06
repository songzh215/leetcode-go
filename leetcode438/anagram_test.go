package leetcode

import (
	"testing"
)

func TestIsAna(t *testing.T) {
	s1 := "cbaebabacd"
	s2 := "abc"
	res := findAnagrams(s1, s2)
	if len(res) != 2 {
		t.Errorf("ana wrong lenght is %d", len(res))
	}
	if res[0] != 0 {
		t.Errorf("ana wrong is %d", res[0])
	}
	if res[1] != 6 {
		t.Errorf("ana wrong is %d", res[1])
	}
}

func TestIsAna1(t *testing.T) {
	s1 := "abab"
	s2 := "ab"
	res := findAnagrams(s1, s2)
	if len(res) != 3 {
		t.Errorf("ana wrong lenght is %d", len(res))
	}
	if res[0] != 0 {
		t.Errorf("ana wrong is %d", res[0])
	}
	if res[1] != 1 {
		t.Errorf("ana wrong is %d", res[1])
	}
	if res[2] != 2 {
		t.Errorf("ana wrong is %d", res[2])
	}
}
