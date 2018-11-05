package leetcode

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	fast, slow := 1, 0
	for ; slow < len(nums)-1; slow++ {
		if slow >= fast {
			fast = slow + 1
		}
		for fast < len(nums) && nums[fast] == 0 {
			fast++
		}
		if nums[slow] == 0 && fast < len(nums) {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			fast++
		}
	}
}
