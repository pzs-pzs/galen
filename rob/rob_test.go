package rob

import "testing"

var action = []struct {
	d []int
	s int
}{
	{
		d: []int{1, 2, 3, 1},
		s: 4,
	},
	{
		d: []int{2, 7, 9, 3, 1},
		s: 12,
	},
}

func TestRob(t *testing.T) {
	for _, a := range action {
		v := rob(a.d)
		if v == a.s {
			continue
		} else {
			panic("eee")
		}
	}
}
