package window

func maximumSubarraySum(nums []int, k int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	left, right, sum, n := 0, 0, 0, len(nums)
	var ans int
	cnt := map[int]struct{}{}
	for right < n {

		for {
			_, ok := cnt[nums[right]]
			if !(ok || right-left >= k) {
				break
			}
			sum -= nums[left]
			delete(cnt, nums[left])
			left++
		}

		sum += nums[right]
		cnt[nums[right]] = struct{}{}
		right++
		if right-left == k {
			ans = max(ans, sum)
		}

	}
	return int64(ans)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
