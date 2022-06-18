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

var action2 = []struct {
	d []int
	s int
}{
	{
		d: []int{1, 2, 3, 1},
		s: 4,
	},
	{
		d: []int{1, 2, 3},
		s: 3,
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

func TestRob2(t *testing.T) {
	for _, a := range action2 {
		v := rob2(a.d)
		if v == a.s {
			println(v)
			continue
		} else {
			panic("eee")
		}
	}
}
