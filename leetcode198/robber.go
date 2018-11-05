package leetcode

func max(x, y int) int {
	res := x
	if y > x {
		res = y
	}
	return res
}

func rob(nums []int) int {
	// dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) == 1 {
		return dp[0]
	}
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
