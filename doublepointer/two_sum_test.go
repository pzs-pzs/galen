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

func Test_sumSubseqWidths(t *testing.T) {
	sumSubseqWidths([]int{2, 1, 3})
}

func Test_expressiveWords(t *testing.T) {
	expressiveWords("heeellooo", []string{"hello"})
}

func Test_pivotInteger(t *testing.T) {
	println(pivotInteger(8))
}

func Test_appendCharacters(t *testing.T) {
	println(appendCharacters("coaching", "coding"))
}

func Test_removeNodes(t *testing.T) {
	a := &ListNode{
		Val:  5,
		Next: nil,
	}
	b := &ListNode{
		Val:  2,
		Next: nil,
	}
	c := &ListNode{
		Val:  13,
		Next: nil,
	}
	d := &ListNode{
		Val:  3,
		Next: nil,
	}
	e := &ListNode{
		Val:  8,
		Next: nil,
	}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	head.Next = a
	removeNodes(a)
}

func Test_check(t *testing.T) {
	println(check([]int{6, 10, 6}))
}

func Test_countSubarrays(t *testing.T) {
	println(countSubarrays([]int{2, 5, 1, 4, 3, 6}, 1))
}
