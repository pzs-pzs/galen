package tree

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	var f func(root *TreeNode)
	f = func(r *TreeNode) {
		if r == nil {
			return
		}
		ans = append(ans, r.Val)
		f(r.Left)
		f(r.Right)
	}
	f(root)
	return ans
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var ans []string
	var f func(root *TreeNode, prefix string)
	f = func(root *TreeNode, prefix string) {
		if root.Left != nil {
			f(root.Left, fmt.Sprintf("%s->%d", prefix, root.Left.Val))
		}
		if root.Right != nil {
			f(root.Right, fmt.Sprintf("%s->%d", prefix, root.Right.Val))
		}
		if root.Left == nil && root.Right == nil {
			ans = append(ans, prefix)
		}
	}
	f(root, fmt.Sprintf("%d", root.Val))
	return ans
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var fn func(l, r *TreeNode) bool
	fn = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		return l.Val == r.Val && fn(l.Left, r.Right) && fn(l.Right, r.Left)
	}
	return fn(root.Left, root.Right)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	q := append([]*TreeNode(nil), root)

	var tmp []int
	var next []*TreeNode
	for {
		if len(q) == 0 {
			break
		}
		first := q[0]
		tmp = append(tmp, first.Val)
		q = q[1:]
		if first.Left != nil {
			next = append(next, first.Left)
		}
		if first.Right != nil {
			next = append(next, first.Right)
		}
		if len(q) == 0 {
			ans = append(ans, tmp)
			tmp = []int{}
			q = append(q, next...)
			next = []*TreeNode{}
		}
	}
	return ans
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := append([]*TreeNode(nil), root)
	var (
		ans     [][]int
		next    []*TreeNode
		tmp     []int
		reverse bool
	)

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		tmp = append(tmp, first.Val)
		if first.Left != nil {
			next = append(next, first.Left)
		}
		if first.Right != nil {
			next = append(next, first.Right)
		}
		if len(queue) == 0 {
			if reverse {
				var t []int
				for i := len(tmp) - 1; i >= 0; i-- {
					t = append(t, tmp[i])
				}
				ans = append(ans, t)
			} else {
				ans = append(ans, tmp)
			}
			tmp = []int{}
			reverse = !reverse
			queue = append(queue, next...)
			next = []*TreeNode{}
		}
	}
	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	q := append([]*TreeNode(nil), root)
	var tmp []*TreeNode
	for len(q) > 0 {
		first := q[0]
		if first.Left != nil {
			tmp = append(tmp, first.Left)
		}
		if first.Right != nil {
			tmp = append(tmp, first.Right)
		}
		q = q[1:]
		if len(q) == 0 {
			q = append(q, tmp...)
			ans += operationsCnt(tmp)
			tmp = []*TreeNode{}
		}
	}
	return ans
}

func operationsCnt(nums []*TreeNode) int {
	tmp := append([]*TreeNode(nil), nums...)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i].Val < nums[j].Val
	})
	indexMap := map[int]int{}
	for i, v := range nums {
		indexMap[v.Val] = i
	}
	var ans int
	for i := 0; i < len(tmp); i++ {
		for {
			idx := indexMap[tmp[i].Val]
			if idx != i {
				ans++
				tmp[i], tmp[idx] = tmp[idx], tmp[i]
			} else {
				break
			}
		}
	}
	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	if root == nil {
		return nil
	}
	var tmp []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		tmp = append(tmp, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	ans := make([][]int, len(queries))
	for index, q := range queries {
		b := []int{-1, -1}
		i := sort.SearchInts(tmp, q+1) - 1
		if i >= 0 {
			b[0] = tmp[i]
		}
		i = sort.SearchInts(tmp, q)
		if i < len(tmp) {
			b[1] = tmp[i]
		}
		ans[index] = b
	}
	return ans
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	if cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func findKOr(nums []int, k int) (ans int) {
	for i := 0; i < 31; i++ {
		cnt1 := 0
		for _, x := range nums {
			cnt1 += x >> i & 1
		}
		if cnt1 >= k {
			ans |= 1 << i
		}
	}
	return
}
