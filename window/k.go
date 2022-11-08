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

// s = "aabcabcab", maxLetters = 2, minSize = 2, maxSize = 3
// 3
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	l, r := 0, 0
	letters := map[uint8]int{}
	cnt := map[string]int{}
	for r < len(s) {
		// 初始化窗口
		if r-l < minSize-1 {
			letters[s[r]] = letters[s[r]] + 1
			r++
			continue
		}
		letters[s[r]] = letters[s[r]] + 1
		if len(letters) <= maxLetters {
			cnt[s[l:r+1]] = cnt[s[l:r+1]] + 1
		}
		r++
		letters[s[l]] = letters[s[l]] - 1
		if letters[s[l]] == 0 {
			delete(letters, s[l])
		}
		l++
	}
	var ans int
	for _, v := range cnt {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func numOfSubarrays(arr []int, k int, threshold int) int {
	var sum int
	l, r := 0, 0
	var ans int
	for r < len(arr) {
		if r-l < k-1 {
			sum += arr[r]
			r++
			continue
		}
		sum += arr[r]
		if (float64(sum) / float64(k)) >= float64(threshold) {
			ans++
		}
		r++
		sum -= arr[l]
		l++
	}
	return ans
}
