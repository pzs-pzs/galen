package prefixsum

import "testing"

func TestSum(t *testing.T) {
	ans := subarraySum([]int{1, 1, 1}, 2)
	println(ans)
}
