package dfs

import (
	"container/heap"
	"math"
	"sort"
	"strings"
)

func solve(board [][]byte) {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	row := len(board)
	col := len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i > row || j < 0 || j > col || board[i][j] == 'X' {
			return
		}
		board[i][j] = 'Z'
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1])
		}

	}

	// 行边界
	for i := 0; i < col; i++ {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[row-1][i] == 'O' {
			dfs(row-1, i)
		}
	}

	// 列边界
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][col-1] == 'O' {
			dfs(i, col-1)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Z' {
				board[i][j] = '0'
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	var ans int
	row, col := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i > row-1 || j < 0 || j > col-1 || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1])
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	ans := baseCosts[0]
	for _, c := range baseCosts {
		ans = min(ans, c)
	}
	var dfs func(int, int)
	dfs = func(p, curCost int) {
		if abs(ans-target) < curCost-target {
			return
		} else if abs(ans-target) >= abs(curCost-target) {
			if abs(ans-target) > abs(curCost-target) {
				ans = curCost
			} else {
				ans = min(ans, curCost)
			}
		}
		if p == len(toppingCosts) {
			return
		}
		dfs(p+1, curCost+toppingCosts[p]*2)
		dfs(p+1, curCost+toppingCosts[p])
		dfs(p+1, curCost)
	}
	for _, c := range baseCosts {
		dfs(0, c)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 0; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left <= 0 {
		return this.prefixSum[right]
	}
	return this.prefixSum[right] - this.prefixSum[left-1]
}

func isSubstringPresent(s string) bool {
	for i := len(s) - 1; i >= 1; i-- {
		if strings.Contains(s, string(s[i])+string(s[i-1])) {
			return true
		}
	}
	return false
}

func countSubstrings(s string, c byte) int64 {
	var cnt int
	for _, i := range s {
		if byte(i) == c {
			cnt++
		}
	}
	return int64(cnt) + int64(cnt*(cnt-1))/2
}

func minimumDeletions(word string, k int) int {
	var freq []int = make([]int, 26)
	for _, w := range word {
		freq[w-'a']++
	}
	sort.Ints(freq)
	var l int
	for l = 0; l < len(freq); l++ {
		if freq[l] > 0 {
			break
		}
	}
	r := len(freq) - 1
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > r {
			return 0
		}
		if freq[r]-freq[l] <= k {
			return 0
		}
		return min(freq[l]+dfs(l+1, r), freq[r]-freq[l]-k+dfs(l, r-1))
	}
	return dfs(l, r)
}

func minStoneSum(piles []int, k int) int {
	tmp := IntHeap(piles)
	heap.Init(&tmp)
	for i := 0; i < k; i++ {
		val := heap.Pop(&tmp).(int)
		heap.Push(&tmp, val/2)
		heap.Fix(&tmp, 0)
	}
	var ans int
	for _, pile := range piles {
		ans += pile
	}
	return ans
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func triangleType(nums []int) string {
	a, b, c := nums[0], nums[1], nums[2]
	if a == b && a == c {
		return "equilateral"
	}
	if !(a+b > c && a+c > b && b+c > a) {
		return "none"
	}
	if a == b || a == c || b == c {
		return "isosceles"
	}
	return "scalene"
}

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		if a[0] == b[0] {
			return a[1] > b[1]
		}
		return a[0] < b[0]
	})
	var ans int
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			a, b := points[i], points[j]
			if a[0] <= b[0] && b[1] > a[1] {
				continue
			}
			var f bool
			for k := 0; k < len(points); k++ {
				if k == i || k == j {
					continue
				}
				tmp := points[k]
				if tmp[0] >= a[0] && tmp[0] <= b[0] && tmp[1] <= a[1] && tmp[1] >= b[1] {
					f = true
					break
				}
			}
			if f {
				continue
			}
			ans++
		}
	}
	return ans
}

func maximumSubarraySum(nums []int, k int) int64 {
	ans := math.MinInt
	minS := map[int]int{}
	sum := 0
	for _, x := range nums {
		s, ok := minS[x+k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x-k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x]
		if !ok || sum < s {
			minS[x] = sum
		}

		sum += x
	}
	if ans == math.MinInt {
		return 0
	}
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maximumLength(nums []int) int {
	exist := map[int][]int{}
	for i, num := range nums {
		exist[num] = append(exist[num], i)
	}
	var ans int = 1
	if exist[1] != nil {
		ans = len(exist[1]) - 1 | 1
	}
	for _, num := range nums {
		if num == 1 {
			continue
		}
		if len(exist[num]) < 2 {
			continue
		}
		var tmp int = 2
		for {
			num = num * num
			if len(exist[num]) >= 2 {
				tmp += 2
				continue
			} else {
				break
			}
		}
		if exist[num] != nil {
			ans = max(ans, tmp+1)
		} else {
			ans = max(ans, tmp-1)
		}
	}
	return ans
}

func flowerGame(n int, m int) int64 {
	if n <= 1 && m <= 1 {
		return 0
	}
	if n == 1 {
		return int64(m / 2)
	}
	if m == 1 {
		return int64(n / 2)
	}
	if n%2 == 0 && m%2 == 0 {
		return int64((n / 2) * (m / 2) * 2)
	}
	if n%2 == 0 && m%2 != 0 {
		return int64((n/2)*(m/2+1) + (n/2)*(m/2))
	}
	if n%2 != 0 && m%2 == 0 {
		return int64((n/2+1)*(m/2) + (n/2)*(m/2))
	}
	return int64((n/2+1)*(m/2) + (n/2)*(m/2+1))
}
