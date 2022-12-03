package stack

import (
	"strings"
)

func isValid(s string) bool {
	stk := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stk = append(stk, v)
			continue
		}
		if len(stk) == 0 {
			return false
		}
		if v == ')' && stk[len(stk)-1] == '(' {
			stk = stk[:len(stk)-1]
			continue
		}
		if v == ']' && stk[len(stk)-1] == '[' {
			stk = stk[:len(stk)-1]
			continue
		}
		if v == '}' && stk[len(stk)-1] == '{' {
			stk = stk[:len(stk)-1]
			continue
		}
		return false
	}
	return len(stk) == 0
}

func nextGreaterElements(nums []int) []int {
	stack := []int{}
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	nums = append(nums, nums...)
	for i := range nums {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if nums[top] >= nums[i] {
				break
			}
			if top%n != i%n {
				ans[top%n] = nums[i%n]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	ans, n := 0, len(nums)
	l := make([]int, n+10)
	r := make([]int, n+10)
	for i := range l {
		l[i] = -1
		r[i] = n
	}
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if nums[top] >= nums[i] {
				break
			}
			r[top] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if nums[top] > nums[i] {
				break
			}
			l[top] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		if nums[i] < left || nums[i] > right {
			continue
		}
		ans += (i - l[i]) * (r[i] - i)
	}
	return ans
}

func removeDuplicateLetters(s string) string {
	var ans strings.Builder
	var stack []int
	r := make([]int, 26)
	for i := 0; i < len(s); i++ {
		r[s[i]-'a'] = i
	}
	visited := make([]bool, 26)
	for i := 0; i < len(s); i++ {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if s[i] > s[top] {
				break
			}
			if visited[s[i]-'a'] {
				break
			}
			if r[s[top]-'a'] > i {
				stack = stack[:len(stack)-1]
				visited[s[top]-'a'] = false
				continue
			}
			break
		}
		if !visited[s[i]-'a'] {
			visited[s[i]-'a'] = true
			stack = append(stack, i)
		}
	}
	for _, v := range stack {
		ans.WriteByte(s[v])
	}
	return ans.String()
}

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

type FreqStack struct {
	cnt    map[int]int
	stacks [][]int
}

func Constructor() FreqStack {
	return FreqStack{cnt: map[int]int{}}
}

func (f *FreqStack) Push(val int) {
	c := f.cnt[val]
	if c == len(f.stacks) { // 这个元素的频率已经是目前最多的，现在又出现了一次
		f.stacks = append(f.stacks, []int{val}) // 那么必须创建一个新栈
	} else {
		f.stacks[c] = append(f.stacks[c], val) // 否则就压入对应的栈
	}
	f.cnt[val]++ // 更新频率
}

func (f *FreqStack) Pop() int {
	back := len(f.stacks) - 1
	st := f.stacks[back]
	bk := len(st) - 1
	val := st[bk] // 弹出最右侧栈的栈顶
	if bk == 0 {  // 栈为空
		f.stacks = f.stacks[:back] // 删除
	} else {
		f.stacks[back] = st[:bk]
	}
	f.cnt[val]-- // 更新频率
	return val
}

func secondHighest(s string) int {
	ans := make([]int, 10)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			ans[s[i]-'0']++
		}
	}
	for i := 9; i >= 0; i-- {
		if ans[i] > 0 {
			for j := i - 1; j >= 0; j-- {
				if ans[j] > 0 {
					return j
				}
			}
			break
		}
	}
	return -1
}

func makeFancyString(s string) string {
	ans := strings.Builder{}
	for i := 0; i < len(s); {
		ans.WriteByte(s[i])
		k := 0
		j := i + 1
		for ; j < len(s); j++ {
			if s[j] == s[i] {
				k++
				if k <= 1 {
					ans.WriteByte(s[j])
				}
				continue
			}
			break
		}
		i = j
	}
	return ans.String()
}
