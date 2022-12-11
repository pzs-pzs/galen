package merge

import (
	"container/list"
	"math"
	"sort"
	"strings"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// sort start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if ans[len(ans)-1][1] < intervals[i][0] {
			ans = append(ans, intervals[i])
			continue
		}
		if intervals[i][1] > ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = intervals[i][1]
		}
	}
	return ans
}

type RangeModule struct {
	item *list.List
}

type rangeItem struct {
	l, r int
}

func Constructor() RangeModule {
	return RangeModule{
		item: list.New(),
	}
}

func (r *RangeModule) AddRange(left int, right int) {
	if r.item.Len() == 0 {
		r.item.PushBack(&rangeItem{
			l: left,
			r: right,
		})
		return
	}

	if r.item.Front().Value.(*rangeItem).l > right {
		r.item.PushFront(&rangeItem{
			l: left,
			r: right,
		})
		return
	}

	if r.item.Back().Value.(*rangeItem).r < left {
		r.item.PushBack(&rangeItem{
			l: left,
			r: right,
		})
		return
	}

	cur := r.item.Front()
	for cur != nil && cur.Value.(*rangeItem).l < right {
		cur = cur.Next()
	}
	if cur == nil {
		return
	}

	if cur.Value.(*rangeItem).r < right {
		r.item.InsertBefore(&rangeItem{
			l: left,
			r: right,
		}, cur)
		return
	}
	// merge
	if left < cur.Value.(*rangeItem).l {
		cur.Value.(*rangeItem).l = left
	}
	if right > cur.Value.(*rangeItem).r {
		cur.Value.(*rangeItem).r = right
	}
	for cur != nil && cur.Next() != nil && cur.Next().Value.(*rangeItem).r < right {
		r.item.Remove(cur.Next())
	}
	if cur != nil && cur.Next() != nil && cur.Next().Value.(*rangeItem).l > right {
		return
	}

	// last
	cur = cur.Next()
	if cur != nil && left < cur.Value.(*rangeItem).l {
		cur.Value.(*rangeItem).l = left
	}
	if cur != nil && right > cur.Value.(*rangeItem).r {
		cur.Value.(*rangeItem).r = right
	}
}

func (r *RangeModule) QueryRange(left int, right int) bool {
	cur := r.item.Front()
	for cur != nil {
		if cur.Value.(*rangeItem).l <= left && cur.Value.(*rangeItem).r >= right {
			return true
		}
		cur = cur.Next()
	}
	return false
}

func (r *RangeModule) RemoveRange(left int, right int) {
	if r.item.Len() == 0 {
		return
	}
	if right < r.item.Front().Value.(*rangeItem).l {
		return
	}
	if left > r.item.Back().Value.(*rangeItem).r {
		return
	}
	if left < r.item.Front().Value.(*rangeItem).l && right > r.item.Back().Value.(*rangeItem).r {
		r.item = list.New()
		return
	}

	cur := r.item.Front()
	for cur != nil && left >= cur.Value.(*rangeItem).r {
		cur = cur.Next()
	}

	if cur.Value.(*rangeItem).l > right {
		return
	}

	if left > cur.Value.(*rangeItem).l {
		r.item.InsertBefore(&rangeItem{l: cur.Value.(*rangeItem).l, r: left}, cur)
	}

	for cur != nil && cur.Value.(*rangeItem).r <= right {
		next := cur.Next()
		r.item.Remove(cur)
		cur = next
	}

	if cur != nil && cur.Value.(*rangeItem).l < right {
		cur.Value = &rangeItem{
			l: right,
			r: cur.Value.(*rangeItem).r,
		}
	}

}

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words)-1; i++ {
		c := words[i]
		n := words[i+1]
		if c[len(c)-1] != n[0] {
			return false
		}
	}
	last := words[len(words)-1]
	first := words[0]
	return last[len(last)-1] == first[0]
}

func dividePlayers(skill []int) int64 {
	var sum int
	cache := map[int]int{}
	for _, s := range skill {
		sum += s
		cache[s]++
	}
	avg := sum / (len(skill) >> 1)
	var rst int
	for _, s := range skill {
		if cache[s] > 0 {
			v, ok := cache[avg-s]
			if !ok {
				return -1
			}
			if v <= 0 {
				return -1
			}
			rst += (avg - s) * s
			cache[avg-s]--
			cache[s]--
		}
	}
	return int64(rst)
}

func minScore(n int, roads [][]int) int {
	type edge struct{ to, d int }
	g := make([][]edge, n)
	for _, e := range roads {
		x, y, d := e[0]-1, e[1]-1, e[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}
	ans := math.MaxInt32
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, e := range g[x] {
			ans = min(ans, e.d)
			if !vis[e.to] {
				dfs(e.to)
			}
		}
	}
	dfs(0)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numDifferentIntegers(word string) int {
	ans := map[string]int{}
	for i := 0; i < len(word); {
		if !digit(word[i]) {
			i++
			continue
		}
		s := i
		for ; s < len(word); s++ {
			if word[s] != '0' {
				break
			}
		}

		j := i
		for ; j < len(word); j++ {
			if !digit(word[j]) {
				break
			}
		}
		ans[word[s:j]]++
		i = j
		continue
	}
	return len(ans)
}

func digit(b byte) bool {
	return b >= '0' && b <= '9'
}

func deleteGreatestValue(grid [][]int) int {
	for _, g := range grid {
		sort.Slice(g, func(i, j int) bool {
			return g[i] < g[j]
		})
	}
	var ans int
	for j := 0; j < len(grid[0]); j++ {
		m := -1
		for i := 0; i < len(grid); i++ {
			m = max(m, grid[i][j])
		}
		ans += m
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSquareStreak(nums []int) int {
	cache := map[int]bool{}
	for _, num := range nums {
		cache[num] = true
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := -1
	for _, num := range nums {
		cnt := 0
		for true {
			if !cache[num] {
				break
			}
			num = num * num
			cnt++
		}
		if cnt >= 2 {
			ans = max(ans, cnt)
		}

	}
	return ans
}
