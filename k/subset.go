package k

// subsets https://leetcode.cn/problems/subsets/
func subsets(nums []int) [][]int {
	var ans [][]int
	ans = append(ans, []int{})
	var f func(nums []int, pre []int)
	f = func(nums []int, pre []int) {
		if len(nums) == 0 {
			return
		}
		//tmp := append([]int(nil), pre...)
		for i, num := range nums {
			tmp := append([]int(nil), pre...)
			ans = append(ans, append(tmp, num))
			f(nums[i+1:], append(tmp, num))
		}
	}

	for i, num := range nums {
		ans = append(ans, []int{num})
		f(nums[i+1:], []int{num})
	}
	return ans
}

// subsetsMask
func subsetsMask(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	for i := 0; i < 1<<n; i++ {
		var tmp []int
		for index, num := range nums {
			if i>>index&1 > 0 {
				tmp = append(tmp, num)
			}
		}
		ans = append(ans, tmp)
	}

	return ans
}
