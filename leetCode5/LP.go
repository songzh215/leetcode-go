package leetcode

func isPalindrome(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	for i := 0; i < length/2+1; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := s[0:1]
	for i := 2; i <= len(s); i++ { // window length.
		for j := 0; j < len(s)-i+1; j++ { // start pos.
			if isPalindrome(s[j : j+i]) {
				res = s[j : j+i]
				break
			}
		}
	}
	return res
}
