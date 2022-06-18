package prefixsum

func subarraySum(nums []int, k int) int {
	prefix := make([]int, len(nums)+1, len(nums)+1)
	prefix[0] = 0
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	var ans int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if prefix[j+1]-prefix[i] == k {
				ans++
			}
		}
	}
	return ans
}
