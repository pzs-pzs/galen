package bfs

func solve(board [][]byte) {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	row := len(board)
	col := len(board[0])

	var bfs func(i, j int)
	type point struct {
		x int
		y int
	}
	bfs = func(i, j int) {
		board[i][j] = 'Z'
		q := append([]*point(nil), &point{x: i, y: j})
		for len(q) > 0 {
			first := q[0]
			q = q[1:]
			for _, dir := range dirs {
				x := first.x + dir[0]
				y := first.y + dir[1]
				if x < 0 || x > row-1 || y < 0 || y > col-1 || board[x][y] != 'O' {
					continue
				}
				board[x][y] = 'Z'
				q = append(q, &point{
					x: x,
					y: y,
				})
			}
		}

	}

	// 行边界
	for i := 0; i < col; i++ {
		if board[0][i] == 'O' {
			bfs(0, i)
		}
		if board[row-1][i] == 'O' {
			bfs(row-1, i)
		}
	}

	// 列边界
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			bfs(i, 0)
		}
		if board[i][col-1] == 'O' {
			bfs(i, col-1)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Z' {
				board[i][j] = 'O'
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	var ans int
	row, col := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var bfs func(i, j int)
	type point struct {
		x int
		y int
	}
	bfs = func(i, j int) {
		q := append([]*point(nil), &point{
			x: i,
			y: j,
		})
		for len(q) > 0 {
			first := q[0]
			q = q[1:]
			for _, dir := range dirs {
				x, y := first.x+dir[0], first.y+dir[1]
				if x < 0 || x > row-1 || y < 0 || y > col-1 || grid[x][y] == '0' {
					continue
				}
				grid[x][y] = '0'
				q = append(q, &point{
					x: x,
					y: y,
				})
			}
		}

	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				ans++
				bfs(i, j)
			}
		}
	}
	return ans
}
