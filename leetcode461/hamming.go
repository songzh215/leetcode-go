package leetcode

func hammingDistance(x int, y int) int {
	res := 0
	judge := x ^ y
	for judge != 0 {
		if judge&1 == 1 {
			res++
		}
		judge >>= 1
	}
	return res
}
