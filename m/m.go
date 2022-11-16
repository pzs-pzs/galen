package m

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	m := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > m {
			return false
		}
		m = min(nums[i], nums[i+1])
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
