package leetcode

func partition(nums []int, p, q int) int {
	pivot := nums[p]
	i := p
	j := q
	for j > i {
		for nums[j] <= pivot && i < j {
			j--
		}
		for nums[i] >= pivot && i < j {
			i++
		}
		if i < j {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
		}
	}
	tmp := nums[i]
	nums[i] = pivot
	nums[p] = tmp
	return i
}

func findKthLargest(nums []int, k int) int {
	pos := partition(nums, 0, len(nums)-1)
	for pos != k-1 {
		if pos > k-1 {
			pos = partition(nums, 0, pos-1)
		} else if pos < k-1 {
			pos = partition(nums, pos+1, len(nums)-1)
		} else {
			break
		}
	}
	return nums[pos]
}
