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
		if sum >= threshold*k {
			ans++
		}
		r++
		sum -= arr[l]
		l++
	}
	return ans
}

func numberOfSubstrings(s string) int {
	var ans int
	l, r := 0, 0
	cnt := map[uint8]int{}
	for ; r < len(s); r++ {
		cnt[s[r]] = cnt[s[r]] + 1
		for ; cnt['a'] > 0 && cnt['b'] > 0 && cnt['c'] > 0; l++ {
			ans += len(s) - r
			cnt[s[l]] = cnt[s[l]] - 1
		}
	}
	return ans
}

func maxScore(cardPoints []int, k int) int {
	s := len(cardPoints) + k - 1
	l, r := len(cardPoints)-k, len(cardPoints)-k
	tmp := append(cardPoints, cardPoints...)
	var ans int
	var sum int
	for r <= s {
		sum += tmp[r]
		if r-l == k-1 {
			ans = max(ans, sum)
			sum -= tmp[l]
			l++
		}
		r++
	}
	return ans
}

func maxVowels(s string, k int) int {
	l, r := 0, 0
	var ans, cnt int
	for r < len(s) {
		if is(s[r]) {
			cnt++
		}
		if r-l == k-1 {
			ans = max(ans, cnt)
			if is(s[l]) {
				cnt--
			}
			l++
		}
		r++
	}
	return ans
}
func is(c uint8) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	l, r := -1, 0
	cnt := map[int]int{}
	var ans int
	for r < len(nums) {
		if nums[r] >= left && nums[r] <= right {
			cnt[nums[r]]++
		} else {

		}
		for l < r && nums[r] >= left && nums[r] <= right {
			if len(cnt) > 0 {
				ans++
			}
			l++
			if nums[l] >= left && nums[l] <= right {
				cnt[nums[l]]--
				if cnt[nums[l]] == 0 {
					delete(cnt, nums[l])
				}
			}
		}
		r++
	}
	return ans
}
