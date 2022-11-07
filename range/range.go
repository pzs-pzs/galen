package merge

import (
	"container/list"
	"sort"
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
