package leetcode

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	start := 0
	end := length - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[start] <= nums[mid] {
			if target <= nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		if nums[mid] <= nums[end] {
			if target <= nums[end] && target >= nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
