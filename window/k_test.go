package window

import (
	"strings"
	"testing"
)

func Test_firstCompleteIndex(t *testing.T) {
	println(-6 % 4)
}

func Test_areSimilar(t *testing.T) {
	println(areSimilar([][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, 2))
}

func Test_nextLargerNodes(t *testing.T) {
	a := &ListNode{
		Val: 2,
	}
	b := &ListNode{
		Val: 7,
	}
	c := &ListNode{
		Val: 4,
	}
	d := &ListNode{
		Val: 3,
	}
	e := &ListNode{
		Val: 5,
	}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	nextLargerNodes(a)
}

func Test_minimumFuelCost(t *testing.T) {
	println(minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2))
}

func Test_minReorder(t *testing.T) {
	println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
}

func Test_maximizeTheProfit(t *testing.T) {
	println(maximizeTheProfit(5, [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}))
}

func Test_minimumEffortPath(t *testing.T) {
	//println(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
	//println(minimumEffortPath([][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}))
	println(minimumEffortPath([][]int{{1, 10, 6, 7, 9, 10, 4, 9}}))
}

func Test_secondGreaterElement(t *testing.T) {
	secondGreaterElement([]int{2, 4, 0, 9, 6})
}

func Test_removeDuplicates(t *testing.T) {
	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}

func Test_reverseOddLevels(t *testing.T) {

	root := &TreeNode{Val: 2}
	a := &TreeNode{Val: 3}
	b := &TreeNode{Val: 5}
	c := &TreeNode{Val: 8}
	d := &TreeNode{Val: 13}
	e := &TreeNode{Val: 12}
	f := &TreeNode{Val: 34}
	root.Left = a
	root.Right = b
	a.Left = c
	a.Right = d
	b.Left = e
	b.Right = f
	reverseOddLevels(root)
}

func Test_numOfBurgers(t *testing.T) {
	burgers := numOfBurgers(16, 7)
	println(burgers[0], burgers[1])
}

func Test_numberOfBoomerangs(t *testing.T) {
	numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}})
}

func Test_addMinimum(t *testing.T) {
	println(addMinimum("aaa"))
}

func Test_sumIndicesWithKSetBits(t *testing.T) {
	a := 3
	var ans int
	for a > 0 {
		if a&1 == 1 {
			ans++
		}
		a = a >> 1
	}
}

func Test_distinctDifferenceArray(t *testing.T) {
	distinctDifferenceArray([]int{3, 2, 3, 4, 2})
}

func Test_maxResult(t *testing.T) {
	maxResult([]int{1, -5, -20, 4, -1, 3, -6, -3}, 2)
}

func Test_magicTower(t *testing.T) {
	//println(magicTower([]int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}))
	//println(magicTower([]int{-200, -300, 400, 0}))
	println(magicTower([]int{-1, -1, 10}))
}

func Test_validPartition(t *testing.T) {
	println(validPartition([]int{993335, 993336, 993337, 993338, 993339, 993340, 993341}))
}

func Test_minimumPossibleSum(t *testing.T) {
	println(minimumPossibleSum(39636, 49035))
	println('A' - 'a')
}

func Test_capitalizeTitle(t *testing.T) {
	capitalizeTitle("First leTTeR of EACH Word")
}

func Test_maximumOddBinaryNumber(t *testing.T) {
	println(maximumOddBinaryNumber("0101"))
	println('0' - '0')
	println('1' - '0')
}

func Test_maximumHappinessSum(t *testing.T) {
	maximumHappinessSum([]int{1, 2, 3}, 2)
}

func Test_shortestSubstrings(t *testing.T) {
	//shortestSubstrings([]string{"gfnt", "xn", "mdz", "yfmr", "fi", "wwncn", "hkdy"})
	rst := strings.Split("ironworks-api,photovoltaic-api,building-api,wind-api,pv-api,oil-api,hps-api,cement-api,ewb-api,neb-api,cr-api,mine-api,rsms-api,klrsms-api,xycdrsms-api,bhrsms-api,zscdrsms-api,gszjrsms-api,ship-api,syjzrsms-api,gfcdrsms-api,mr-api,thermal-api,oe-api,pyjzrsms-api,casicflcrsms-api,fertilizer-api,cpv-api", ",")
	println(len(rst))
}

func resultArray(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	for i := 2; i < len(nums); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	return append(arr1, arr2...)
}

func countSubmatrices(grid [][]int, k int) int {
	return 0
}

func Test_minOperations(t *testing.T) {
	minOperations([]int{1, 1, 2, 4, 9}, 20)
}

func Test_isPossibleToSplit(t *testing.T) {
	isPossibleToSplit([]int{6, 1, 3, 1, 1, 8, 9, 2})
}

func Test_countMatchingSubarrays(t *testing.T) {
	countMatchingSubarrays([]int{1, 2, 3, 4, 5, 6}, []int{1, 1})
}

func Test_sellingWood(t *testing.T) {
	sellingWood(3, 5, [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}})
}
