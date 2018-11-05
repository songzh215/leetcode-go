package sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var all_res [][]int
	sort.Ints(nums)
	length := len(nums)
	tracer := make(map[string]bool)
	if length < 3 {
		return all_res
	}
	for i := 0; i < length; i++ {
		if i > 1 && nums[i] == nums[i-1] {
			continue
		}
		r := length - 1
		for j := i + 1; j < r; {
			key := fmt.Sprintf("%d-%d", nums[i], nums[j])
			_, ok := tracer[key]
			if ok {
				j++
				continue
			}
			if nums[i]+nums[j]+nums[r] > 0 {
				r--
			} else if nums[i]+nums[j]+nums[r] == 0 {
				tracer[key] = true
				var res []int
				res = append(res, nums[i])
				res = append(res, nums[j])
				res = append(res, nums[r])
				all_res = append(all_res, res)
				r--
				j++
			} else {
				j++
			}
		}
	}
	return all_res
}
