package doublepointer

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestReverse(t *testing.T) {
	rst := reverse([]int{1, 2, 3, 4})
	assert.Equal(t, rst, []int{4, 3, 2, 1})
}

func TestTwoSum2(t *testing.T) {
	rst := twoSum2([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, rst, []int{1, 2})
}

func Test_numMatchingSubseq(t *testing.T) {
	numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"})

	nums := []int{1, 3, 3, 6, 6, 8}
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 5
	})

	l, r := 0, len(nums)
	var mid int
	target := 3
	for l < r {
		mid = (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}

	}
	println(l - 1)
	println(i)
}
