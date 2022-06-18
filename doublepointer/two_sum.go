package doublepointer

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
