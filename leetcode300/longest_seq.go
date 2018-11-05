package leetcode

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	dp := make([]int, length)
	dp[0] = 1
	maxans := 1
	for i := 1; i < length; i++ {
		maxLength := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > maxLength {
				maxLength = dp[j] + 1
			}
		}
		dp[i] = maxLength
		if dp[i] > maxans {
			maxans = dp[i]
		}
	}
	return maxans
}
