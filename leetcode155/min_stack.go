package leetcode

type MinStack struct {
	Min      int
	Size     int
	TopPos   int
	Elements []int
}

const INT_MAX = int(^uint(0) >> 1)

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{Min: INT_MAX, Size: 0, TopPos: -1, Elements: make([]int, 1000)}
	return stack
}

func (this *MinStack) Push(x int) {
	this.TopPos = this.TopPos + 1
	this.Elements[this.TopPos] = x
	if x < this.Min {
		this.Min = x
	}
	this.Size = this.Size + 1
}

func (this *MinStack) Pop() {
	if this.TopPos != -1 {
		if this.Min == this.Elements[this.TopPos] {
			min := INT_MAX
			for i := 0; i < this.Size-1; i++ {
				if this.Elements[i] < min {
					min = this.Elements[i]
				}
			}
			this.Min = min
		}
		this.TopPos = this.TopPos - 1
		this.Size = this.Size - 1
	}
}

func (this *MinStack) Top() int {
	if this.TopPos >= 0 && this.TopPos < len(this.Elements) {
		return this.Elements[this.TopPos]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
