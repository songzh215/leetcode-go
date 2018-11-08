package leetcode

func findUnsortedSubarray(nums []int) int {
	len := len(nums)
	var i, j, n, p, q int
	for i = 0; i < len-1; i++ {
		if nums[i] > nums[i+1] {
			break
		}
	}
	if i == len-1 {
		return 0
	}
	for j = len - 1; j > 0; j-- {
		if nums[j] < nums[j-1] {
			break
		}
	}
	max := nums[i]
	min := nums[i]
	for n = i; n <= j; n++ {
		if nums[n] > max {
			max = nums[n]
		}
		if nums[n] < min {
			min = nums[n]
		}
	}
	pos1 := 0
	pos2 := len - 1
	for p = 0; p < len; p++ {
		if nums[p] > min {
			pos1 = p
			break
		}
	}
	for q = len - 1; q > 0; q-- {
		if nums[q] < max {
			pos2 = q
			break
		}
	}
	return pos2 - pos1 + 1
}
