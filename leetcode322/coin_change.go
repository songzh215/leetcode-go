package leetcode

func min(l, r int) int {
	res := l
	if r < l {
		res = r
	}
	return res
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	length := len(coins)
	if amount == 0 {
		return 0
	}
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < length; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
