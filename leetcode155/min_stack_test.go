package leetcode

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	if obj.GetMin() != -3 {
		t.Errorf("get min should be -3, actual %d", obj.GetMin())
	}
	obj.Pop()
	if obj.Top() != 0 {
		t.Errorf("get top should be 0, actual %d", obj.Top())
	}
	if obj.GetMin() != -2 {
		t.Errorf("get min should be -2, actual %d", obj.GetMin())
	}

}
