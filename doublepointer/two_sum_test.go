package doublepointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	rst := reverse([]int{1, 2, 3, 4})
	assert.Equal(t, rst, []int{4, 3, 2, 1})
}

func TestTwoSum2(t *testing.T) {
	rst := twoSum2([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, rst, []int{1, 2})
}
