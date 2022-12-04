package dfs

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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
