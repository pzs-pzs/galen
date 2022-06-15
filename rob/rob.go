package rob

// rob 198. 打家劫舍
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	tmp := make([]int, len(nums), len(nums))

	tmp[0] = nums[0]
	tmp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		if tmp[i-2]+nums[i] > tmp[i-1] {
			tmp[i] = tmp[i-2] + nums[i]
			continue
		}
		tmp[i] = tmp[i-1]
	}
	return tmp[len(tmp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
