package main

func twoSum(nums []int, target int) []int {
	var res []int
	size := len(nums)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i)
				res = append(res, j)
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	twoSum(nums, 9)
}
