package bfs

import "testing"

func Test_solve(t *testing.T) {
	in := [][]byte{
		{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'},
	}
	solve(in)
}
