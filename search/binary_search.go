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

func mySqrt(x int) int {
	l, r := 1, x
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			l = mid + 1
		}
		if mid*mid > x {
			r = mid - 1
		}
	}
	return r
}
