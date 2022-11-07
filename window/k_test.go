package window

import "testing"

func Test_maximumSubarraySum(t *testing.T) {
	sum := maximumSubarraySum([]int{1, 1, 1, 1, 3, 2}, 3)
	println(sum)
}
