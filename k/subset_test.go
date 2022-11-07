package k

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	ans := subsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 10, 0})
	fmt.Printf("%v", ans)
}

func Test_subsetsMask(t *testing.T) {
	ans := subsetsMask([]int{1, 2, 3})
	fmt.Printf("%v", ans)
}
