package doublepointer

import (
	"math"
	"sort"
	"strings"
)

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

func sumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	var ans int
	for i := 1; i < 1<<len(nums); i++ {
		a, b := -1, math.MaxInt
		for j := 0; j < len(nums); j++ {
			if i>>j&1 > 0 {
				a = max(a, nums[j])
				b = min(b, nums[j])
			}
		}
		ans = (ans + (a - b)) % (1e9 + 7)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func numDistinct(s string, t string) int {
	var ans int

	return ans
}

func unequalTriplets(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for z := j + 1; z < len(nums); z++ {
				if nums[i] != nums[j] && nums[j] != nums[z] && nums[i] != nums[z] {
					ans++
				}
			}
		}
	}
	return ans
}

func expand(s, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] != t[j] {
			return false
		}
		ch := s[i]
		cntI := 0
		for i < n && s[i] == ch {
			cntI++
			i++
		}
		cntJ := 0
		for j < m && t[j] == ch {
			cntJ++
			j++
		}
		if cntI < cntJ || cntI > cntJ && cntI < 3 {
			return false
		}
	}
	return i == n && j == m
}

func expressiveWords(s string, words []string) (ans int) {
	for _, word := range words {
		if expand(s, word) {
			ans++
		}
	}
	return
}

func pivotInteger(n int) int {
	p, s := make([]int, n), make([]int, n)
	p[0] = 1
	s[n-1] = n
	for i := 2; i <= n; i++ {
		p[i-1] = p[i-2] + i
	}
	for i := n - 1; i >= 1; i-- {
		s[i-1] = s[i] + i
	}
	for i := 1; i <= n; i++ {
		if p[i-1] == s[i-1] {
			return i
		}
	}
	return -1
}

func appendCharacters(s string, t string) int {
	l, r := 0, 0

	for l < len(s) && r < len(t) {
		if s[l] == t[r] {
			l++
			r++
			continue
		}
		l++
	}
	if r == len(t) {
		return 0
	}
	return len(t) - r
}

func removeNodes(head *ListNode) *ListNode {
	var stack []*ListNode
	n := head
	for n != nil {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top.Val > n.Val {
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
		n = n.Next
	}
	head = stack[0]
	stack[len(stack)-1].Next = nil
	for i := range stack {
		if i == 0 {
			continue
		}
		stack[i-1].Next = stack[i]
	}
	return head
}

func check(nums []int) bool {
	i := 1
	for ; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			break
		}
	}
	if i == len(nums) {
		return true
	}
	for ; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return nums[0] >= nums[len(nums)-1]
}

func countSubarrays(nums []int, k int) int {
	var ans int
	for i, num := range nums {
		if num != k {
			continue
		}
		ans++
		// 偶数
		l, r := i, i+1
		for l >= 0 && r <= len(nums)-1 {
			if nums[l] <= k && nums[r] > k {
				ans++
				l--
				r++
				continue
			}
			break
		}
		// 奇数
		l, r = i-1, i+1
		for l >= 0 && r <= len(nums)-1 {
			if nums[l] < k && nums[r] > k {
				ans++
				l--
				r++
				continue
			}
			break
		}

	}
	return ans
}

type Allocator struct {
	mem   []int
	index int
}

func Constructor(n int) Allocator {
	return Allocator{
		mem:   make([]int, n),
		index: 0,
	}
}

func (a *Allocator) Allocate(size int, mID int) int {
	for i := 0; i < len(a.mem); i++ {
		if a.mem[i] != 0 {
			continue
		}
		end := i
		for ; end < len(a.mem); end++ {
			if a.mem[end] != 0 {
				break
			}
		}
		if size <= end-i+1 {
			for s := i; s < i+size; s++ {
				a.mem[s] = mID
			}
			return i
		}
	}
	return -1
}

func (a *Allocator) Free(mID int) int {
	var cnt int
	for i := 0; i < len(a.mem); i++ {
		if a.mem[i] == mID {
			a.mem[i] = 0
			cnt++
		}
	}
	return cnt
}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j, q int, ans []int, index int, visited [][]bool)
	dfs = func(i, j, q int, ans []int, index int, visited [][]bool) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			return
		}
		if visited[i][j] {
			return
		}
		if grid[i][j] >= q {
			return
		}
		ans[index]++
		visited[i][j] = true
		for _, d := range dirs {
			dfs(i+d[0], j+d[1], q, ans, index, visited)
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		visited := make([][]bool, len(grid))
		for i := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}
		dfs(0, 0, q, ans, i, visited)
	}
	return ans

}

func beautySum(s string) (ans int) {
	for i := range s {
		cnt := [26]int{}
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a']++
			minFreq := len(s)
			maxFreq := 0
			for _, c := range cnt {
				if c > 0 {
					minFreq = min(minFreq, c)
				}
				maxFreq = max(maxFreq, c)
			}
			ans += maxFreq - minFreq
		}
	}
	return
}

func similarPairs(words []string) int {
	var ans int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if is(words[i], words[j]) {
				ans++
			}
		}
	}
	return ans
}

func is(s1, s2 string) bool {
	a := make([]int, 26)
	for _, v := range s1 {
		a[v-'a'] = 1
	}
	b := make([]int, 26)
	for _, v := range s2 {
		b[v-'a'] = 1
	}
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func smallestValue(n int) int {
	for {
		x, s := n, 0
		for i := 2; i*i <= x; i++ {
			for ; x%i == 0; x /= i {
				s += i
			}
		}
		if x > 1 {
			s += x
		}
		if s == n {
			return n
		}
		n = s
	}
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	var ans bool
	e := make([][]int, n)
	for _, edge := range edges {
		start, end := edge[0], edge[1]
		e[start] = append(e[start], end)
		e[end] = append(e[end], start)
	}
	visited := make([]bool, n)
	var dfs func(s int)
	dfs = func(p int) {
		visited[p] = true
		if p == destination {
			ans = true
			return
		}
		for _, c := range e[p] {
			if visited[c] {
				continue
			}
			dfs(c)
		}
	}
	dfs(source)
	return ans
}

func maximumScore(a int, b int, c int) int {
	var ans int
	rec := []int{a, b, c}
	sort.Slice(rec, func(i, j int) bool {
		return rec[i] < rec[j]
	})
	for !(rec[0] == 0 && rec[1] == 0) {
		ans++
		rec[1]--
		rec[2]--
		sort.Slice(rec, func(i, j int) bool {
			return rec[i] < rec[j]
		})
	}
	return ans
}

func finalValueAfterOperations(operations []string) int {
	var ans int
	for _, operation := range operations {
		if strings.Contains(operation, "++") {
			ans++
			continue
		}
		ans--
	}
	return ans
}

func minimumMoves(s string) int {
	var ans int
	for i := 0; i < len(s); {
		if s[i] == 'X' {
			i += 3
			ans++
			continue
		}
		i++
	}
	return ans
}

func closetTarget(words []string, target string, startIndex int) int {
	ans := len(words)
	for i, word := range words {
		if word == target {
			t := min(abs(startIndex-i), len(words)-abs(startIndex-i))
			ans = min(t, ans)
		}
	}
	if ans == len(words) {
		return -1
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func takeCharacters(s string, k int) int {
	s = s + s
	t := len(s)
	for i := len(s)/2 - 1; i >= 0; i-- {
		ans := make([]int, 3, 3)
		for j := i; j < len(s) && j-i < len(s)/2; j++ {
			ans[s[j]-'a']++
			if ans[0] >= k && ans[1] >= k && ans[2] >= k {
				t = min(t, j-i+1)
				break
			}
		}
	}
	if t == len(s) {
		return 0
	}
	return t
}
