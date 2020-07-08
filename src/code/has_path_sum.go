package main

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func hasPathSum(root *TreeNode1, sum int) bool {
	if root == nil {
		return false
	}
	if sum == 0 {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
