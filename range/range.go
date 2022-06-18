package merge

import "sort"

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
