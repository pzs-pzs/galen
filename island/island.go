package island

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	c := image[sr][sc]
	row, col := len(image), len(image[0])
	var backTrack func(i, j int)
	backTrack = func(i, j int) {
		if i < 0 || i >= row || j < 0 || j >= col {
			return
		}
		if image[i][j] == newColor || image[i][j] != c {
			return
		}
		image[i][j] = newColor
		backTrack(i+1, j)
		backTrack(i-1, j)
		backTrack(i, j+1)
		backTrack(i, j-1)
	}
	backTrack(sr, sc)
	return image
}
