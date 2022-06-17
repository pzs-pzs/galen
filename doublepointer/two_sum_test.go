package sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	rst := reverse([]int{1, 2, 3, 4})
	assert.Equal(t, rst, []int{4, 3, 2, 1})
}
