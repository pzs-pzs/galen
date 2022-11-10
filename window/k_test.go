package window

import "testing"

func Test_maximumSubarraySum(t *testing.T) {
	sum := maximumSubarraySum([]int{1, 1, 1, 1, 3, 2}, 3)
	println(sum)
}

func Test_maxFreq(t *testing.T) {
	ans := maxFreq("abcde", 2, 3, 3)
	println(ans)
}

func Test_numOfSubarrays(t *testing.T) {
	ans := numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4)
	println(ans)
	ans = numOfSubarrays([]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5)
	println(ans)
}

func Test_numberOfSubstrings(t *testing.T) {
	ans := numberOfSubstrings("abcabc")
	println(ans)
}
