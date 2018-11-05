package leetcode

func exists(tgt string, wordDict []string) bool {
	if tgt == "" {
		return true
	}
	for _, e := range wordDict {
		if tgt == e {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	// dp[i] = dp[i] || dp[j] && sub(i-j)
	length := len(s)
	if length == 0 {
		return true
	}
	dp := make([]bool, length+1)
	dp[0] = true
	for i := 1; i <= length; i++ {
		for j := 0; !dp[i] && j < i; j++ {
			dp[i] = dp[i] || dp[j] && exists(s[j:i], wordDict)
		}
	}
	return dp[length]
}
