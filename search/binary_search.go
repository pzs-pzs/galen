package search

func isBadVersion(version int) bool { return true }

func firstBadVersion(n int) int {
	l, r := 1, n
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
