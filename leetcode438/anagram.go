package leetcode

func isAnagram(s1 string, cache []int) (bool, int) {
	sCache := make([]int, 26)
	for j := len(s1) - 1; j >= 0; j-- {
		if cache[s1[j]-'a'] == 0 {
			return false, j
		}
		sCache[s1[j]-'a'] = sCache[s1[j]-'a'] + 1
		if sCache[s1[j]-'a'] > cache[s1[j]-'a'] {
			return false, 0
		}
	}
	return true, 0
}

func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) < len(p) {
		return res
	}
	cache := make([]int, 26)
	for i := 0; i < len(p); i++ {
		cache[p[i]-'a'] = cache[p[i]-'a'] + 1
	}

	for i := 0; i <= len(s)-len(p); i++ {
		ok, pos := isAnagram(s[i:i+len(p)], cache)
		if ok {
			res = append(res, i)
		} else {
			i = i + pos
		}
	}
	return res
}
