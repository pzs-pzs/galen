package doublepointer

import "sort"

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := head
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		dummy = dummy.Next
	}
	return dummy
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		tmp := numbers[left] + numbers[right]
		if tmp == target {
			return []int{left + 1, right + 1}
		}
		if tmp < target {
			left++
		}
		if tmp > target {
			right--
		}
	}
	return []int{-1, -1}
}

func numMatchingSubseq(s string, words []string) int {
	var ans int
	index := make(map[rune][]int)
	for i, v := range s {
		index[v] = append(index[v], i)
	}
	for _, word := range words {
		lastIndex := -1
		ok := true
		for _, v := range word {
			idxs := index[v]
			l, r := 0, len(idxs)
			for l < r {
				mid := (l + r) / 2
				if idxs[mid] >= lastIndex+1 {
					r = mid
				} else {
					l = mid + 1
				}
			}
			if r == len(idxs) {
				ok = false
				break
			}
			lastIndex = idxs[r]
		}
		if ok {
			ans++
		}

	}
	return ans
}

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	ans := map[float64]struct{}{}
	for l < r {
		ans[float64(nums[l]+nums[r])/float64(2)] = struct{}{}
		l++
		r--
	}
	return len(ans)
}

func topKFrequent(nums []int, k int) []int {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num] = cnt[num] + 1
	}
	var s [][]int
	for i, c := range cnt {
		s = append(s, []int{c, i})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][0] > s[j][0]
	})
	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, s[i][1])
	}
	return ans
}
