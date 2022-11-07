package tree

import "fmt"

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
