package leetcode

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	if hammingDistance(1, 4) != 2 {
		t.Errorf("hamming distance is %d, expect %d", hammingDistance(1, 4), 2)
	}
}
