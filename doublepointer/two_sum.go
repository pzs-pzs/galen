package sum

func twoSum(nums []int, target int) []int {
	index := make(map[int]int)
	for i, v := range nums {
		if j, ok := index[target-v]; ok {
			return []int{j, i}
		}
		index[v] = i
	}
	return []int{}
}

func reverse(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
