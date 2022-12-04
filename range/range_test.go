package merge

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestMerge(t *testing.T) {
	ans := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	println(ans)
}

// ["RangeModule","addRange","addRange","addRange","queryRange","queryRange","queryRange","removeRange","queryRange"]
// [[],[10,180],[150,200],[250,500],[50,100],[180,300],[600,1000],[50,150],[50,100]]
var action = []struct {
	op   string
	data []int
	rst  bool
}{
	{
		op:   "addRange",
		data: []int{10, 180},
		rst:  false,
	},
	{
		op:   "addRange",
		data: []int{150, 200},
		rst:  false,
	},
	{
		op:   "addRange",
		data: []int{250, 500},
		rst:  true,
	},
	{
		op:   "queryRange",
		data: []int{50, 200},
		rst:  false,
	},
	{
		op:   "queryRange",
		data: []int{180, 300},
		rst:  true,
	},
	{
		op:   "queryRange",
		data: []int{600, 1000},
		rst:  false,
	},
	{
		op:   "removeRange",
		data: []int{50, 150},
		rst:  false,
	},
	{
		op:   "queryRange",
		data: []int{50, 100},
		rst:  false,
	},
}

func TestRangeModule(t *testing.T) {
	rangeModule := Constructor()
	for _, v := range action {
		if v.op == "queryRange" {
			queryRange := rangeModule.QueryRange(v.data[0], v.data[1])
			if queryRange != v.rst {
				panic("err")
			}
			continue
		}
		if v.op == "addRange" {
			rangeModule.AddRange(v.data[0], v.data[1])
			continue
		}
		if v.op == "removeRange" {
			rangeModule.RemoveRange(v.data[0], v.data[1])
		}
	}
}

func TestEmptySlice(t *testing.T) {
	a := make([]int, 0, 0)
	println((*reflect.SliceHeader)(unsafe.Pointer(&a)).Data)
	b := make([]int, 0, 0)
	println((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	var c []int
	println((*reflect.SliceHeader)(unsafe.Pointer(&c)).Data)
	d := new([]int)
	println((*reflect.SliceHeader)(unsafe.Pointer(d)).Data)
	*d = append(*d, 1)
	println((*reflect.SliceHeader)(unsafe.Pointer(d)).Data)
	println((*reflect.SliceHeader)(unsafe.Pointer(d)).Len)
	println((*reflect.SliceHeader)(unsafe.Pointer(d)).Cap)
}

func Test_dividePlayers(t *testing.T) {
	println(dividePlayers([]int{3, 2, 5, 1, 3, 4}))
}
