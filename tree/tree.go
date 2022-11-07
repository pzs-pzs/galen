package tree

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
