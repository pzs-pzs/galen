package backpack

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	println(strongPasswordCheckerII("a1A!A!A!"))
}

func TestB(t *testing.T) {
	successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7)
}

func TestC(t *testing.T) {
	// 10 150 + 70 = 210 3*25=75
	in := [][]int{{3, 50}, {7, 10}, {12, 25}}
	fmt.Printf("%f", calculateTax(in, 10))
}
