package leetcode

func numJewelsInStones(J string, S string) int {
	t := make([]int, 52)
	for i := 0; i < len(J); i++ {
		if J[i]-'a' >= 0 && J[i]-'a' < 26 {
			t[J[i]-'a'] = 1
		} else {
			t[J[i]-'A'+26] = 1
		}
	}
	count := 0
	for j := 0; j < len(S); j++ {
		if S[j]-'a' >= 0 && S[j]-'a' < 26 {
			if t[S[j]-'a'] == 1 {
				count++
			}
		} else {
			if t[S[j]-'A'+26] == 1 {
				count++
			}
		}
	}
	return count
}
