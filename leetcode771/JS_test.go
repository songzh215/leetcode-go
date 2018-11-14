package leetcode

import (
	"testing"
)

func TestJewStone(t *testing.T) {
	J := "aA"
	S := "aAAbbbb"
	if numJewelsInStones(J, S) != 3 {
		t.Errorf("jew in stone is %d", numJewelsInStones(J, S))
	}
}

func TestJewStone1(t *testing.T) {
	J := "z"
	S := "ZZ"
	if numJewelsInStones(J, S) != 0 {
		t.Errorf("jew in stone is %d", numJewelsInStones(J, S))
	}
}
