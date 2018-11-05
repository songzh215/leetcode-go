package sum

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	if len(res) != 2 {
		t.Errorf("num is not 2, actual %d", len(res))
	}
	for _, r := range res {
		sum := 0
		if len(r) != 3 {
			t.Errorf("num is not 3, actual %d", len(r))
		}
		for _, n := range r {
			sum = sum + n
			t.Log("number %d", n)
		}
		if sum != 0 {
			t.Errorf("sum not 0, actual %d", sum)
		}
	}
}
