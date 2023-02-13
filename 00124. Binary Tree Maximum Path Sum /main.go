package _0124__Binary_Tree_Maximum_Path_Sum_

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	mx := math.MinInt
	dfs(root, &mx)
	return mx
}

func dfs(root *TreeNode, mx *int) {
	if root == nil {
		return
	}
	dfs(root.Left, mx)
	dfs(root.Right, mx)

	l, r := 0, 0
	if root.Left != nil && root.Left.Val > 0 {
		l = root.Left.Val
	}
	if root.Right != nil && root.Right.Val > 0 {
		r = root.Right.Val
	}
	*mx = max(*mx, root.Val+l+r)
	root.Val += max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
