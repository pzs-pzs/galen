package union

func longestConsecutive(nums []int) int {
	var (
		ans int
	)
	parent := make(map[int]int, len(nums))
	for _, num := range nums {
		parent[num] = num
	}

	for _, num := range nums {
		_, ok := parent[num+1]
		if ok {
			unionRoot(parent, num, num+1)
		}
	}

	for _, num := range nums {
		right := findRoot(parent, num)
		ans = max(ans, right-num+1)
	}

	return ans
}

func findRoot(parent map[int]int, k int) int {
	if parent[k] == k {
		return k
	}
	return findRoot(parent, parent[k])
}

func unionRoot(parent map[int]int, a, b int) {
	rootA := findRoot(parent, a)
	rootB := findRoot(parent, b)
	if rootA != rootB {
		parent[rootA] = rootB
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	size := row * col
	parent := make([]int, size)
	var ans int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			parent[i*row+j] = i*row + 1
			ans++
		}
	}
	var find func(k int) int
	find = func(k int) int {
		if parent[k] == k {
			return k
		}
		return find(parent[k])
	}
	var union func(i, j int)
	union = func(i, j int) {
		rootA := find(i)
		rootB := find(j)
		if rootA == rootB {
			return
		}
		ans--
		parent[rootA] = rootB
	}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x >= 0 && x <= row-1 && y >= 0 && y <= col-1 && grid[x][y] == 1 {
						union(i*row+j, x*row+y)
					}
				}
			}
		}
	}

	return ans
}
