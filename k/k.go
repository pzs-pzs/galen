package k

import (
	"math/bits"
	"sort"
)

func ambiguousCoordinates(s string) []string {
	f := func(i, j int) []string {
		var res []string
		for k := 1; k <= j-i; k++ {
			l, r := s[i:i+k], s[i+k:j]
			ok := (l == "0" || l[0] != '0') && (r == "" || r[len(r)-1] != '0')
			if ok {
				t := ""
				if k < j-i {
					t = "."
				}
				res = append(res, l+t+r)
			}
		}
		return res
	}

	n := len(s)
	var ans []string
	for i := 2; i < n-1; i++ {
		for _, x := range f(1, i) {
			for _, y := range f(i, n-1) {
				ans = append(ans, "("+x+", "+y+")")
			}
		}
	}
	return ans
}

func countBalls(lowLimit int, highLimit int) int {
	box := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		var t int
		k := i
		for k != 0 {
			t += k % 10
			k = k / 10
		}
		box[t] = box[t] + 1
	}
	var ans [][]int
	for k, v := range box {
		ans = append(ans, []int{k, v})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][1] > ans[j][1]
	})
	return ans[0][1]
}

func numWays(n int, relation [][]int, k int) int {
	edge := make([][]int, n)
	for _, r := range relation {
		edge[r[0]] = append(edge[r[0]], r[1])
	}
	var ans int
	var dfs func(x, step int)
	dfs = func(x, step int) {
		if step == k && x == n-1 {
			ans++
			return
		}
		for _, e := range edge[x] {
			dfs(e, step+1)
		}
	}
	dfs(0, 1)
	bits.OnesCount(1)
	return ans
}
