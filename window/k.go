package window

import (
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func maximumSubarraySum(nums []int, k int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	left, right, sum, n := 0, 0, 0, len(nums)
	var ans int
	cnt := map[int]struct{}{}
	for right < n {

		for {
			_, ok := cnt[nums[right]]
			if !(ok || right-left >= k) {
				break
			}
			sum -= nums[left]
			delete(cnt, nums[left])
			left++
		}

		sum += nums[right]
		cnt[nums[right]] = struct{}{}
		right++
		if right-left == k {
			ans = max(ans, sum)
		}

	}
	return int64(ans)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// s = "aabcabcab", maxLetters = 2, minSize = 2, maxSize = 3
// 3
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	l, r := 0, 0
	letters := map[uint8]int{}
	cnt := map[string]int{}
	for r < len(s) {
		// 初始化窗口
		if r-l < minSize-1 {
			letters[s[r]] = letters[s[r]] + 1
			r++
			continue
		}
		letters[s[r]] = letters[s[r]] + 1
		if len(letters) <= maxLetters {
			cnt[s[l:r+1]] = cnt[s[l:r+1]] + 1
		}
		r++
		letters[s[l]] = letters[s[l]] - 1
		if letters[s[l]] == 0 {
			delete(letters, s[l])
		}
		l++
	}
	var ans int
	for _, v := range cnt {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func numOfSubarrays(arr []int, k int, threshold int) int {
	var sum int
	l, r := 0, 0
	var ans int
	for r < len(arr) {
		if r-l < k-1 {
			sum += arr[r]
			r++
			continue
		}
		sum += arr[r]
		if sum >= threshold*k {
			ans++
		}
		r++
		sum -= arr[l]
		l++
	}
	return ans
}

func numberOfSubstrings(s string) int {
	var ans int
	l, r := 0, 0
	cnt := map[uint8]int{}
	for ; r < len(s); r++ {
		cnt[s[r]] = cnt[s[r]] + 1
		for ; cnt['a'] > 0 && cnt['b'] > 0 && cnt['c'] > 0; l++ {
			ans += len(s) - r
			cnt[s[l]] = cnt[s[l]] - 1
		}
	}
	return ans
}

func maxScore(cardPoints []int, k int) int {
	s := len(cardPoints) + k - 1
	l, r := len(cardPoints)-k, len(cardPoints)-k
	tmp := append(cardPoints, cardPoints...)
	var ans int
	var sum int
	for r <= s {
		sum += tmp[r]
		if r-l == k-1 {
			ans = max(ans, sum)
			sum -= tmp[l]
			l++
		}
		r++
	}
	return ans
}

func maxVowels(s string, k int) int {
	l, r := 0, 0
	var ans, cnt int
	for r < len(s) {
		if is(s[r]) {
			cnt++
		}
		if r-l == k-1 {
			ans = max(ans, cnt)
			if is(s[l]) {
				cnt--
			}
			l++
		}
		r++
	}
	return ans
}
func is(c uint8) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	l, r := -1, 0
	cnt := map[int]int{}
	var ans int
	for r < len(nums) {
		if nums[r] >= left && nums[r] <= right {
			cnt[nums[r]]++
		} else {

		}
		for l < r && nums[r] >= left && nums[r] <= right {
			if len(cnt) > 0 {
				ans++
			}
			l++
			if nums[l] >= left && nums[l] <= right {
				cnt[nums[l]]--
				if cnt[nums[l]] == 0 {
					delete(cnt, nums[l])
				}
			}
		}
		r++
	}
	return ans
}

func maxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}
	var ans int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if masks[i]&masks[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func vowelStrings(words []string, left int, right int) int {
	var ans int
	for i := left; i <= right; i++ {
		if vowel(words[i][0]) && vowel(words[i][len(words[i])-1]) {
			ans++
		}
	}
	return ans
}

func vowel(r byte) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func findTheLongestBalancedSubstring(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if (j-i)%2 == 0 {
				continue
			}
			var cnt int
			m, n := i, j
			for m < n {
				if s[m] == '0' && s[n] == '1' {
					cnt++
					m++
					n--
					continue
				}
				break
			}
			if cnt > ans {
				ans = cnt
			}
		}

	}
	return ans
}

func findChampion(n int, edges [][]int) int {
	weak := make([]bool, n)
	for _, e := range edges {
		weak[e[1]] = true
	}
	ans := -1
	for i, w := range weak {
		if !w {
			if ans != -1 {
				return -1
			}
			ans = i
		}
	}
	return ans
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	ans := make([]int, len(spells))
	sort.Ints(potions)
	for i, spell := range spells {
		index := sort.Search(len(potions), func(j int) bool {
			return int64(spell*potions[j]) >= success
		})
		ans[i] = len(potions) - index
	}
	return ans
}

type NumArray struct {
	data      []int
	prefixSum []int
}

//func Constructor(nums []int) NumArray {
//	prefixSum := make([]int, len(nums)+1)
//	for i, num := range nums {
//		prefixSum[i+1] = prefixSum[i] + num
//	}
//	return NumArray{
//		data:      nums,
//		prefixSum: prefixSum,
//	}
//}

func (this *NumArray) Update(index int, val int) {
	oldVal := this.data[index]
	for i := index + 1; i < len(this.prefixSum); i++ {
		this.prefixSum[i] = this.prefixSum[i] + (val - oldVal)
	}
	this.data[index] = val

}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	return 0
}

func longestAlternatingSubarray(nums []int, threshold int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			continue
		}
	A:
		for j := i; j < len(nums); j++ {
			var tmp int
			var f bool = true
			for k := i; k <= j; k++ {
				if nums[k] > threshold {
					break A
				}
				if k != j && nums[k]%2 == nums[k+1]%2 {
					f = false
					break A
				}
			}
			tmp = j - i + 1
			if f && tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}

func maximumStrongPairXor(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if abs(nums[i], nums[j]) > min(nums[i], nums[j]) {
				continue
			}
			tmp := nums[i] ^ nums[j]
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func findHighAccessEmployees(access_times [][]string) []string {
	t := make(map[string][]int)
	for _, v := range access_times {
		t[v[0]] = append(t[v[0]], convert(v[1]))
	}
	var ans []string
	for name, ts := range t {
		sort.Ints(ts)
		for i := 0; i < len(ts); {
			j := i + 1
			for ; j < len(ts); j++ {
				if ts[j]-ts[i] >= 60 {
					break
				}
			}
			if j-i > 2 {
				ans = append(ans, name)
				break
			}
			i = i + 1

		}
	}
	return ans
}

func convert(s string) int {
	a, _ := strconv.Atoi(s[0:2])
	b, _ := strconv.Atoi(s[2:])
	return a*60 + b
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))

	for i := 0; i < len(nums1); i++ {
		var st []int
		for j := 0; j < len(nums2); j++ {
			// pop
			for len(st) > 0 && nums2[j] > st[len(st)-1] {
				st = st[:len(st)-1]
			}
			// push
			st = append(st, nums2[j])
			if len(st) > 1 && st[len(st)-2] == nums1[i] {
				ans[i] = st[len(st)-1]
			} else {
				ans[i] = -1
			}
		}
	}
	return ans
}

func maxSubArray(nums []int) int {
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	for j := 1; j < len(nums); j++ {
		ans[j] = max(nums[j], nums[j]+ans[j-1])
	}
	var r int = -100000
	for _, an := range ans {
		if an > r {
			r = an
		}
	}
	return r
}

func minPathCost(grid [][]int, moveCost [][]int) int {
	ans := make([][]int, len(grid))
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, len(grid[0]))
	}
	ans[0] = grid[0]

	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			tmp := 100000
			for k := 0; k < len(grid[i-1]); k++ {
				if moveCost[grid[i-1][k]][j]+ans[i-1][k]+grid[i][k] < tmp {
					tmp = moveCost[grid[i-1][k]][j] + ans[i-1][k] + grid[i][k]
				}
			}
			ans[i][j] = tmp
		}
	}
	var rst = 100000
	for i := 0; i < len(ans[0]); i++ {
		if rst > ans[len(ans)-1][i] {
			rst = ans[len(ans)-1][i]
		}
	}
	return rst
}

func entityParser(text string) string {
	for i := 0; i < len(text); i++ {
		if text[i] != '&' {
			continue
		}
		for j := i + 1; j < len(text); j++ {
			if text[i:j+1] == "&quot;" {
				text = text[:i] + "\"" + text[j+1:]
				break
			}

			if text[i:j+1] == "&apos;" {
				text = text[:i] + "'" + text[j+1:]

				break
			}
			if text[i:j+1] == "&amp;" {
				text = text[:i] + "&" + text[j+1:]

				break
			}
			if text[i:j+1] == "&gt;" {
				text = text[:i] + ">" + text[j+1:]

				break
			}
			if text[i:j+1] == "&lt;" {
				text = text[:i] + "<" + text[j+1:]

				break
			}
			if text[i:j+1] == "&frasl;" {
				text = text[:i] + "/" + text[j+1:]

				break
			}
		}
	}
	return text
}

func countPairs(nums []int, target int) int {
	var ans int
	sort.Ints(nums)
	l, r := 0, len(nums)
	for l > r {
		if nums[l]+nums[r] < target {
			l++
			ans += r - l
			continue
		}
		r--
	}
	return ans
}

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	i := 0
	for ; i < len(s1); i++ {
		if i >= len(s2) || i >= len(s3) || s2[i] != s1[i] || s3[i] != s1[i] {
			break
		}
	}
	if i == 0 {
		return -1
	}
	return len(s1) - i + len(s2) - i + len(s3) - i
}

func minimumSteps(s string) int64 {
	var ans int64
	var cntZero int64
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ans += cntZero
		} else {
			cntZero++
		}
	}
	return ans
}

func maximumXorProduct(a int64, b int64, n int) int {
	var m int64
	var x int64
	for x = 0; x < 1<<n; x++ {
		if (((a^x)%(1000000000+7))*((b^x)%(1000000000+7)))%(1000000000+7) > m {
			m = (((a ^ x) % (1000000000 + 7)) * ((b ^ x) % (1000000000 + 7))) % (1000000000 + 7)
		}
	}
	println(m)
	return int(m)
}

func sumSubarrayMins(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	monoStack := []int{}
	for i, x := range arr {
		for len(monoStack) > 0 && x <= arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && arr[i] < arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n - i
		} else {
			right[i] = monoStack[len(monoStack)-1] - i
		}
		monoStack = append(monoStack, i)
	}
	for i, x := range arr {
		ans = (ans + left[i]*right[i]*x) % mod
	}
	return
}

type FrontMiddleBackQueue struct {
	data []int
}

//func Constructor() FrontMiddleBackQueue {
//	return FrontMiddleBackQueue{}
//}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.data = append([]int{val}, this.data...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	var tmp []int
	tmp = append(tmp, this.data[0:len(this.data)/2]...)
	tmp = append(tmp, val)
	tmp = append(tmp, this.data[len(this.data)/2:]...)
	this.data = tmp

}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.data = append(this.data, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.data) == 0 {
		return -1
	}
	rst := this.data[0]
	this.data = this.data[1:]
	return rst
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.data) == 0 {
		return -1
	}
	if len(this.data)%2 != 0 {
		rst := this.data[len(this.data)/2]
		tmp := this.data[0 : len(this.data)/2]
		tmp = append(tmp, this.data[len(this.data)/2+1:]...)
		this.data = tmp
		return rst
	}
	rst := this.data[len(this.data)/2-1]
	tmp := this.data[0 : len(this.data)/2-1]
	tmp = append(tmp, this.data[len(this.data)/2:]...)
	this.data = tmp
	return rst
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.data) == 0 {
		return -1
	}
	rst := this.data[len(this.data)-1]
	this.data = this.data[0 : len(this.data)-1]
	return rst
}

type SmallestInfiniteSet struct {
	left  int
	data  []int
	exist map[int]interface{}
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		left:  1,
		data:  []int{},
		exist: map[int]interface{}{},
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.data) == 0 {
		ans := this.left
		this.left++
		return ans
	}
	ans := this.data[0]
	this.data = this.data[1:]
	delete(this.exist, ans)
	return ans
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.left {
		return
	}
	if _, ok := this.exist[num]; ok {
		return
	}
	this.data = append(this.data, num)
	this.exist[num] = struct{}{}
	sort.Ints(this.data)
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	for i := range word1 {
		cnt1[word1[i]-'a']++
		cnt2[word2[i]-'a']++
	}

	a := map[int]int{}
	b := map[int]int{}
	for i := range cnt1 {
		if cnt1[i] != 0 && cnt2[i] == 0 {
			return false
		}
		if cnt1[i] == 0 && cnt2[i] != 0 {
			return false
		}
		a[cnt1[i]] = a[cnt1[i]] + 1
		b[cnt2[i]] = b[cnt2[i]] + 1
	}
	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	type pair struct {
		i, j int
	}
	var v = make(map[int]*pair)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			v[mat[i][j]] = &pair{
				i: i,
				j: j,
			}
		}
	}
	var col = make([]int, len(mat))
	var row = make([]int, len(mat[0]))

	for i, val := range arr {
		m := v[val].i
		n := v[val].j
		col[m]++
		row[n]++
		if col[m] == len(mat[0]) {
			return i
		}
		if row[n] == len(mat) {
			return i
		}
	}
	return -1
}

func areSimilar(mat [][]int, k int) bool {
	for i := 0; i < len(mat); i++ {
		// 奇数行
		if i%2 == 0 {
			for j := 0; j < len(mat[i]); j++ {
				if mat[i][j] != mat[i][(j+k)%len(mat[i])] {
					return false
				}
			}
			continue
		}
		for j := 0; j < len(mat[i]); j++ {
			var n int
			if (j-k)%len(mat[i]) >= 0 {
				n = (j - k) % len(mat[i])
			} else {
				n = (j-k)%len(mat[i]) + len(mat[i])
			}
			if mat[i][j] != mat[i][n] {
				return false
			}
		}
	}
	return true
}

func beautifulSubstrings(s string, k int) int {
	var ans int
	for i := 0; i < len(s); i++ {
		a, b := 0, 0
		for j := i; j < len(s); j++ {
			if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
				a++
			} else {
				b++
			}
			if (a == b) && (a+b)%k == 0 {
				ans++
			}
		}
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var cc []int
	var f func(t *TreeNode)
	f = func(t *TreeNode) {
		if t == nil {
			return
		}
		f(t.Left)
		cc = append(cc, t.Val)
		f(t.Right)
	}

	sum := make([]int, len(cc))
	sum[len(sum)-1] = cc[len(cc)-1]
	for j := len(sum) - 2; j >= 0; j-- {
		sum[j] = sum[j+1] + cc[j]
	}
	index := map[int]int{}
	for i, v := range cc {
		index[v] = i
	}
	var k func(t *TreeNode)
	k = func(t *TreeNode) {
		if t == nil {
			return
		}
		k(t.Left)
		t.Val = sum[index[t.Val]]
		k(t.Right)
	}
	return root
}

func nextLargerNodes(head *ListNode) []int {
	type pair struct {
		val   int
		index int
	}
	var (
		stack []*pair
		tmp   []int
	)
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	var ans = make([]int, len(tmp))
	for i, t := range tmp {
		for len(stack) > 0 && stack[len(stack)-1].val < t {
			ans[len(stack)-1] = t
			stack = stack[:len(stack)-1]

		}
		stack = append(stack, &pair{
			val:   t,
			index: i,
		})
	}
	for _, v := range stack {
		ans[v.index] = 0
	}

	return ans
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	if len(roads) == 0 {
		return 0
	}
	dag := map[int][]int{}
	var maxN int = -1
	for _, road := range roads {
		dag[road[0]] = append(dag[road[0]], road[1])
		dag[road[1]] = append(dag[road[1]], road[0])
		if road[0] > maxN {
			maxN = road[0]
		}
		if road[1] > maxN {
			maxN = road[1]
		}
	}
	visited := make([]int, maxN+1)
	var f func(i int) (int, int)
	f = func(i int) (int, int) {
		visited[i] = 1
		var oilSum, seatsSum int
		if i != 0 {
			seatsSum = 1
		}
		for _, next := range dag[i] {
			if visited[next] == 0 {
				oil, seat := f(next)
				oilSum += oil
				seatsSum += seat
			}
		}
		if i == 0 {
			return oilSum, seatsSum
		}
		if seatsSum%seats == 0 {
			oilSum += seatsSum / seats
		} else {
			oilSum += seatsSum/seats + 1
		}
		return oilSum, seatsSum
	}
	oil, _ := f(0)
	return int64(oil)
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	cnt := make([]int, n)
	var dfs func(int, int, int) bool
	dfs = func(i, fa, k int) bool {
		cnt[i]++
		if i == k {
			return true
		}
		ok := false
		for _, j := range g[i] {
			if j != fa {
				ok = dfs(j, i, k)
				if ok {
					break
				}
			}
		}
		if !ok {
			cnt[i]--
		}
		return ok
	}
	for _, t := range trips {
		start, end := t[0], t[1]
		dfs(start, -1, end)
	}
	var dfs2 func(int, int) (int, int)
	dfs2 = func(i, fa int) (int, int) {
		a := price[i] * cnt[i]
		b := a >> 1
		for _, j := range g[i] {
			if j != fa {
				x, y := dfs2(j, i)
				a += min(x, y)
				b += x
			}
		}
		return a, b
	}
	a, b := dfs2(0, -1)
	return min(a, b)
}

func minReorder(n int, connections [][]int) int {
	rag := map[int][]int{}
	direction := map[string]struct{}{}
	for _, c := range connections {
		rag[c[0]] = append(rag[c[0]], c[1])
		rag[c[1]] = append(rag[c[1]], c[0])
		direction[fmt.Sprintf("%d-%d", c[0], c[1])] = struct{}{}
	}
	var ans int
	var dfs func(i int, fa int)
	dfs = func(i int, fa int) {
		if _, ok := direction[fmt.Sprintf("%d-%d", fa, i)]; ok {
			ans++
		}
		for _, next := range rag[i] {
			if next != fa {
				dfs(next, i)
			}
		}
	}
	dfs(0, -1)
	return ans
}

func maximizeTheProfit(n int, offers [][]int) int {
	type pair struct {
		start int
		val   int
	}
	var end = make(map[int][]*pair)
	for _, offer := range offers {
		end[offer[1]] = append(end[offer[1]], &pair{
			start: offer[0],
			val:   offer[2],
		})
	}
	var ans []int = make([]int, n+1)
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1]
		for _, p := range end[i-1] {
			if (ans[p.start] + p.val) > ans[i] {
				ans[i] = ans[p.start] + p.val
			}
		}
	}
	return ans[n]
}

func minimumEffortPath(heights [][]int) int {
	type pair struct{ x, y int }
	var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, m := len(heights), len(heights[0])
	return sort.Search(1e6, func(maxHeightDiff int) bool {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[0][0] = true
		queue := []pair{{}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if p.x == n-1 && p.y == m-1 {
				return true
			}
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && absx(heights[x][y]-heights[p.x][p.y]) <= maxHeightDiff {
					vis[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return false
	})
}

func absx(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	st1, st2 := []int{}, []int{}

	for i := 0; i < n; i++ {
		v := nums[i]
		for len(st2) > 0 && nums[st2[len(st2)-1]] < v {
			res[st2[len(st2)-1]] = v
			st2 = st2[:len(st2)-1]
		}
		pos := len(st1) - 1
		for pos >= 0 && nums[st1[pos]] < v {
			pos--
		}
		st2 = append(st2, st1[pos+1:]...)
		st1 = append(st1[:pos+1], i)
	}
	return res
}

func makeSmallestPalindrome(s string) string {
	l, r := 0, len(s)-1
	ans := make([]byte, len(s))
	for l <= r {
		if s[l] <= s[r] {
			ans[l] = s[l]
			ans[r] = s[l]
		} else {
			ans[l] = s[r]
			ans[r] = s[r]
		}
		l++
		r--
	}
	return bytes.NewBuffer(ans).String()
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	for i := 0; i < len(grid)-stampHeight; i++ {
		for j := 0; j < len(grid[i])-stampWidth; j++ {
			for m := 0; m < stampHeight; m++ {
				for n := 0; n < stampWidth; n++ {
					if grid[i+m][j+n] == 1 {
						continue
					}
				}

			}
		}
	}
	return true
}

func carPooling(trips [][]int, capacity int) bool {
	up := map[int]int{}
	down := map[int]int{}
	for _, t := range trips {
		up[t[1]] = up[t[1]] + t[0]
		down[t[2]] = down[t[2]] - t[0]
	}
	s := 0
	for i := 0; i < 1001; i++ {
		s = up[i] - down[i]
		if s > capacity {
			return false
		}
	}
	return true
}

func corpFlightBookings(bookings [][]int, n int) []int {
	var cnt = make([]int, n+1)
	for _, b := range bookings {
		cnt[b[0]] = cnt[b[0]] + b[2]
		cnt[b[1]-1] = cnt[b[0]-1] - b[2]
	}
	var s int
	var ans = make([]int, n+1)
	for i, c := range cnt {
		s += c
		ans[i] = s
	}
	return ans[0:n]
}

// https://leetcode.cn/problems/divide-intervals-into-minimum-number-of-groups/description/
func minGroups(intervals [][]int) int {
	return 0
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	cur := []*TreeNode{root}
	flag := false
	for len(cur) > 0 {
		var next []*TreeNode
		// 计数层 1，3，5....
		if !flag {
			for i := len(cur) - 1; i >= 0; i-- {
				next = append(next, cur[i].Right, cur[i].Left)
			}
			flag = true
		} else {
			for i := len(cur) - 1; i >= 0; i-- {
				if cur[i] == nil {
					continue
				}
				next = append(next, cur[i].Left, cur[i].Right)
			}
			flag = false
		}
		k := 0
		for _, node := range cur {
			if k > len(next) {
				continue
			}
			node.Left = next[k]
			node.Right = next[k+1]
			k += 2
		}
		// 重置当前层
		cur = []*TreeNode{}
		for _, node := range next {
			if node != nil {
				cur = append(cur, node)
			}
		}
	}
	return root
}

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, word := range words {
		if len(word) == 0 {
			return false
		}
		if word[0] != s[i] {
			return false
		}
	}
	return true
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	cnt := map[int]int{}
	for _, row := range grid {
		for _, col := range row {
			cnt[col] = cnt[col] + 1
		}
	}
	var ans []int = make([]int, 2)
	for i := 1; i <= len(grid)<<1; i++ {
		if cnt[i] > 1 {
			ans[0] = i
		}
		if cnt[i] == 2 {
			ans[1] = i
		}
	}
	return ans
}

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 != 0 || tomatoSlices < cheeseSlices*2 || cheeseSlices*4 < tomatoSlices {
		return nil
	}
	return []int{tomatoSlices/2 - cheeseSlices, cheeseSlices*2 - tomatoSlices/2}
}

func isWinner(player1 []int, player2 []int) int {
	d1 := false
	s1 := 0
	for i, v := range player1 {
		if d1 {
			s1 += 2 * v
		} else {
			s1 += v
		}
		if i <= 1 && v == 10 {
			d1 = true
		}
	}
	d2 := false
	s2 := 0
	for i, v := range player2 {
		if d2 {
			s2 += 2 * v
		} else {
			s2 += v
		}
		if i <= 1 && v == 10 {
			d2 = true
		}
	}
	if s1 > s2 {
		return 1
	}
	if s1 < s2 {
		return 2
	}
	return 0

}

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	ans := money - prices[0] - prices[1]
	if ans < 0 {
		return money
	}
	return ans
}

func removeNodes(head *ListNode) *ListNode {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	var st []int
	for _, d := range data {
		for len(st) > 0 && st[len(st)-1] < d {
			st = st[:len(st)-1]
		}
		st = append(st, d)
	}
	var ans *ListNode
	var pre *ListNode
	for i, v := range st {
		if i == 0 {
			ans = &ListNode{
				Val:  v,
				Next: nil,
			}
			pre = ans
			continue
		}
		pre.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		pre = pre.Next
	}
	return ans
}

func numberOfBoomerangs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	var ans int
	for i := 1; i < len(points)-1; i++ {
		l := i - 1
		r := i + 1
		for l >= 0 && r <= len(points)-1 {
			if distance(points[l][0], points[l][1], points[i][0], points[i][1]) == distance(points[r][0], points[r][1], points[i][0], points[i][1]) {
				ans += 2
				l--
				r++
				continue
			}
			if distance(points[l][0], points[l][1], points[i][0], points[i][1]) < distance(points[r][0], points[r][1], points[i][0], points[i][1]) {
				l--
				continue
			}
			if distance(points[l][0], points[l][1], points[i][0], points[i][1]) > distance(points[r][0], points[r][1], points[i][0], points[i][1]) {
				r++
				continue
			}
		}
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
	for i := 1; i < len(points)-1; i++ {
		l := i - 1
		r := i + 1
		for l >= 0 && r <= len(points)-1 {
			a := distance(points[l][0], points[l][1], points[i][1], points[i][0])
			b := distance(points[l][0], points[l][1], points[i][1], points[i][0])
			if a == b {
				ans += 2
				l--
				r++
				continue
			}
			if a < b {
				l--
				continue
			}
			if a > b {
				r++
				continue
			}
		}
	}
	println(ans)
	return ans
}

func distance(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
}

func minExtraChar(s string, dictionary []string) int {
	ss := map[string]bool{}
	for _, w := range dictionary {
		ss[w] = true
	}
	n := len(s)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] + 1
		for j := 0; j < i; j++ {
			if ss[s[j:i]] && f[j] < f[i] {
				f[i] = f[j]
			}
		}
	}
	return f[n]
}

func addMinimum(word string) int {
	var ans int
	for i := 0; i < len(word)-2; {
		if word[i] == 'a' && word[i+1] == 'b' && word[i+2] == 'c' {
			i += 3
			continue
		}
		if word[i] == 'a' && word[i+1] != 'b' {
			ans += 2
			i++
			continue
		}
		if word[i] == 'b' && word[i+1] != 'c' {
			ans += 2
			i++
			continue
		}
		if word[i] == 'c' {
			ans += 2
			i++
			continue
		}
		if word[i] == 'a' && word[i+1] == 'b' {
			ans += 1
			i += 2
			continue
		}
		if word[i] == 'b' && word[i+1] == 'c' {
			ans += 1
			i += 2
			continue
		}
	}
	return ans
}

func deleteDuplicates(head *ListNode) *ListNode {
	before := &ListNode{}
	ans := before
	cur := head
	next := head.Next
	for cur != nil && next != nil && next.Next != nil {
		if cur.Val == next.Val {
			cur = next
			next = next.Next
			continue
		}
		if next.Val == next.Next.Val {
			next = next.Next
			continue
		}
		before.Next = cur
		before = before.Next
		cur = next
		next = next.Next
	}
	return ans.Next
}

func maximumNumberOfStringPairs(words []string) int {
	visited := make([]bool, len(words))
	var ans int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if visited[i] == false && visited[j] == false && words[i][0] == words[j][1] && words[i][1] == words[j][0] {
				visited[i] = true
				visited[j] = true
				ans += 1
			}
		}
	}
	return ans
}

func minimumRemoval(beans []int) int64 {
	var sum int
	for _, bean := range beans {
		sum += bean
	}
	avg := sum / len(beans)
	var (
		ans int
		min = 10000
	)
	for _, bean := range beans {
		if bean < avg {
			avg += bean
			continue
		}
		if bean < min {
			min = bean
		}
	}
	for _, bean := range beans {
		if bean > min {
			ans += bean - min
		}
	}
	return int64(ans)
}

func maximumSwap(num int) int {
	s := fmt.Sprintf("%d", num)
	m := '1'
	index := 0
	for i, v := range s {
		if v > m {
			m = v
			index = i
		}
	}
	var ans []byte = []byte(s)
	ans[index], ans[0] = ans[0], ans[index]
	a, _ := strconv.Atoi(string(ans))
	return a
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	f := func(a int) int {
		var ans int
		for a > 0 {
			if a&1 == 1 {
				ans++
			}
			a = a >> 1
		}
		return ans
	}
	var ans int
	for _, num := range nums {
		if f(num) == k {
			ans += num
		}
	}
	return ans
}

func distinctDifferenceArray(nums []int) []int {
	f := func(a int64) int {
		var ans int
		for a > 0 {
			if a&1 == 1 {
				ans++
			}
			a = a >> 1
		}
		return ans
	}
	var exist int64
	suffix := make([]int, len(nums))
	suffix[len(suffix)-1] = 0
	exist = exist | (1 << nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = f(exist)
		exist = exist | (1 << nums[i])
	}
	exist = 0
	ans := make([]int, len(nums))
	for i, _ := range ans {
		exist = exist | (1 << nums[i])
		ans[i] = f(exist) - suffix[i]
	}
	return ans
}

func maxResult(nums []int, k int) int {
	rst := make([]int, len(nums))
	rst[0] = nums[0]
	for i := 1; i < len(rst); i++ {
		tmp := math.MinInt
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				tmp = max(tmp, rst[0]+nums[i])
				continue
			}
			tmp = max(tmp, rst[i-j]+nums[i])
		}
		rst[i] = tmp
	}
	return rst[len(rst)-1]
}

func magicTower(nums []int) int {
	var prefix []int
	var sum int = 1
	var ans int
	var tmp []int
	for _, num := range nums {
		if num < 0 {
			//prefix = append(prefix, num)
			//sort.Ints(prefix)

			if len(prefix) == 0 {
				prefix = append(prefix, num)
			} else {

				if num < prefix[0] {
					prefix = append([]int{num}, prefix...)
				} else {
					i := 0
					for ; i < len(prefix)-1; i++ {
						if prefix[i] < num && num < prefix[i+1] {
							break
						}
					}
					var t []int
					for k := 0; k <= i; k++ {
						t = append(t, prefix[k])
					}
					t = append(t, num)
					for k := i + 1; k <= len(prefix); k++ {
						t = append(t, prefix[k])
					}
					prefix = t
				}
			}

		}
		sum += num
		if sum <= 0 {
			t := prefix[0]
			prefix = prefix[1:]
			sum = sum - t
			tmp = append(tmp, t)
			ans++
		}
	}
	if sum <= 0 {
		return -1
	}
	for _, t := range tmp {
		sum += t
		if sum <= 0 {
			return -1
		}
	}
	return ans
}
func replaceValueInTree(root *TreeNode) *TreeNode {

	var q []*TreeNode
	root.Val = 0
	q = append(q, root)
	for len(q) > 0 {
		var next []*TreeNode
		var sum int
		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
				sum += node.Left.Val
			}
			if node.Right != nil {
				next = append(next, node.Right)
				sum += node.Right.Val
			}

		}
		for _, fa := range q {
			childSum := 0
			if fa.Left != nil {
				childSum += fa.Left.Val
			}
			if fa.Right != nil {
				childSum += fa.Right.Val
			}
			if fa.Left != nil {
				fa.Left.Val = sum - childSum
			}
			if fa.Right != nil {
				fa.Right.Val = sum - childSum
			}
		}
		q = next
	}
	return root
}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var ans []int
	var f func(n *Node)
	f = func(n *Node) {
		if n == nil {
			return
		}
		for _, child := range n.Children {
			f(child)
		}
		ans = append(ans, n.Val)
	}
	f(root)
	return ans
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 {
//		return nil
//	}
//	root := &TreeNode{preorder[0], nil, nil}
//	i := 0
//	for ; i < len(inorder); i++ {
//		if inorder[i] == preorder[0] {
//			break
//		}
//	}
//	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
//	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
//	return root
//}

//func buildTree(inorder []int, postorder []int) *TreeNode {
//	if len(inorder) == 0 {
//		return nil
//	}
//	root := &TreeNode{preorder[0], nil, nil}
//	i := 0
//	for ; i < len(inorder); i++ {
//		if inorder[i] == preorder[0] {
//			break
//		}
//	}
//	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
//	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
//	return root
//}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	var ans []int64
	for len(q) > 0 {
		var sum int64
		var tmp []*TreeNode
		for _, node := range q {
			sum += int64(node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		ans = append(ans, sum)
		q = tmp
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	if k > len(ans) {
		return -1
	}
	return ans[k-1]
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var dfs func(r *TreeNode)
	var ans int
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Val >= low && r.Val <= high {
			ans += r.Val
		}
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return ans
}

func validPartition(nums []int) bool {
	var ans = make([]bool, len(nums))
	if len(nums) < 2 {
		return false
	}
	ans[0] = false
	if len(nums) == 2 {
		return nums[0] == nums[1]
	}
	ans[1] = nums[0] == nums[1]
	if len(nums) == 3 {
		return (nums[0] == nums[1] && nums[1] == nums[2]) || (nums[1]-nums[0] == 1 && nums[2]-nums[1] == 1)
	}
	ans[2] = (nums[0] == nums[1] && nums[1] == nums[2]) || (nums[1]-nums[0] == 1 && nums[2]-nums[1] == 1)
	for i := 3; i < len(nums); i++ {
		ans[i] = (ans[i-2] && nums[i-1] == nums[i]) || (ans[i-3] && ((nums[i] == nums[i-1] && nums[i-2] == nums[i-1]) || (nums[i]-nums[i-1] == 1 && nums[i-1]-nums[i-2] == 1)))
	}
	return ans[len(ans)-1]
}

func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	x := 0
	for i, c := range word {
		x = (x*10 + int(c-'0')) % m
		if x == 0 {
			ans[i] = 1
		}
	}
	return ans
}

func minimumPossibleSum(n int, target int) int {
	l, r := 1, n
	var ans int
	m := n + 1
	if m < target {
		m = target
	}
	for l <= r {
		if l+r == target {
			if l == r {
				ans += l
				break
			}
			ans += m
			ans += l
			m++
			r--
			l++
			continue
		}
		if l+r < target {
			ans += l
			l++
			continue
		}
		if l+r > target {
			ans += r
			r--
			continue
		}
	}
	return ans % 1_000_000_007
}

func capitalizeTitle(title string) string {
	var ans []byte
	title = strings.ToUpper(title)
	ans = append(ans, title[0])
	for i := 1; i < len(title)-1; {
		if title[i] == ' ' {
			ans = append(ans, ' ', title[i+1])
			i = i + 2
			continue
		}
		ans = append(ans, title[i]+32)
		i++
	}
	ans = append(ans, title[len(title)-1]+32)
	return string(ans)
}

type FindElements struct {
	vals map[int]bool
}

func Constructor1(root *TreeNode) FindElements {
	vals := make(map[int]bool)
	var dfs func(r *TreeNode, val int)
	dfs = func(r *TreeNode, val int) {
		if r == nil {
			return
		}
		r.Val = val
		vals[val] = true
		dfs(r.Left, 2*val+1)
		dfs(r.Right, 2*val+2)
	}
	dfs(root, 0)
	return FindElements{
		vals: vals,
	}
}

func (this *FindElements) Find(target int) bool {
	return this.vals[target]
}

func maximumOddBinaryNumber(s string) string {
	var cnt = make([]int, 2)
	for _, i := range s {
		cnt[i-'0']++
	}
	var ans string
	for i := 0; i < cnt[1]-1; i++ {
		ans += "1"
	}
	for i := 0; i < cnt[0]; i++ {
		ans += "0"
	}
	return ans + "1"
}

func minimumBoxes(apple []int, capacity []int) int {
	var s int
	for _, a := range apple {
		s += a
	}
	var ans int
	sort.Ints(capacity)
	for i := len(capacity) - 1; i >= 0; i-- {
		if s > capacity[i] {
			ans++
			s -= capacity[i]
			continue
		}
		ans++
		break
	}
	return ans
}

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	index := len(happiness) - 1
	var ans int64
	for i := 0; i < k; i++ {
		if happiness[index]-i > 0 {
			ans += int64(happiness[index] - i)
		}
		index--
	}
	return ans
}

func shortestSubstrings(arr []string) []string {
	var ans []string
	for d, s := range arr {
		var t []string
		for i := 0; i < len(s); i++ {
			for j := i + 1; j <= len(s); j++ {
				var f bool
				for n := 0; n < len(arr); n++ {
					if n == d {
						continue
					}
					if strings.Contains(arr[n], s[i:j]) {
						f = true
						continue
					}

				}
				if !f {
					t = append(t, s[i:j])
				}

			}
		}
		if len(t) == 0 {
			ans = append(ans, "")
			continue
		}
		sort.Slice(t, func(i, j int) bool {
			if len(t[i]) > len(t[j]) {
				return false
			}
			if len(t[i]) == len(t[j]) {
				rst := strings.Compare(t[i], t[j])
				return rst < 0
			}
			return true

		})

		ans = append(ans, t[0])
	}
	return ans
}

func minOperations(nums []int, k int) int {
	//sort.Ints(nums)
	tmp := IntHeap(nums)
	heap.Init(&tmp)
	var ans int
	a, b := heap.Pop(&tmp).(int), heap.Pop(&tmp).(int)
	for a < k {
		ans++
		heap.Push(&tmp, min(a, b)*2+max(a, b))
		if tmp.Len() < 2 {
			break
		}
		a, b = heap.Pop(&tmp).(int), heap.Pop(&tmp).(int)
	}
	return ans
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func isPossibleToSplit(nums []int) bool {
	var cnt = make(map[int]int)
	for _, num := range nums {
		if cnt[num] >= 2 {
			return false
		}
		cnt[num] = cnt[num] + 1
	}
	return true
}

func modifiedMatrix(matrix [][]int) [][]int {
	var maxCol []int = make([]int, len(matrix[0]))
	for col := 0; col < len(maxCol); col++ {
		for row := 0; row < len(matrix); row++ {
			if matrix[row][col] > maxCol[col] {
				maxCol[col] = matrix[row][col]
			}
		}
		for row := 0; row < len(matrix); row++ {
			if matrix[row][col] == -1 {
				matrix[row][col] = maxCol[col]
			}
		}
	}
	return matrix
}

func countMatchingSubarrays(nums []int, pattern []int) int {
	var ans int
	for i := 0; i < len(nums)-len(pattern); i++ {
		var f bool
		for j := 0; j < len(pattern); j++ {
			if pattern[j] == 1 && nums[i+j] < nums[i+j+1] {
				continue
			}
			if pattern[j] == -1 && nums[i+j] > nums[i+j+1] {
				continue
			}
			if pattern[j] == 0 && nums[i+j] == nums[i+j+1] {
				continue
			}
			f = true
			break
		}
		if !f {
			ans++
		}
	}
	return ans
}

func sellingWood(m int, n int, prices [][]int) int64 {
	ans := make([][]int, m+1)
	for i, _ := range ans {
		ans[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			for _, price := range prices {
				h, l, w := price[0], price[1], price[2]
				if i-h >= 0 && j-l >= 0 {
					ans[i][j] = max(ans[i-h][j]+ans[i][j-l]-ans[i-h][j-l]+w, ans[i][j])
				}
			}
		}
	}
	var rst int
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if ans[i][j] > rst {
				rst = ans[i][j]
			}
		}
	}
	return int64(rst)
}

func countKeyChanges(s string) int {
	s = strings.ToUpper(s)
	var ans int
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			ans++
		}
	}
	return ans
}
